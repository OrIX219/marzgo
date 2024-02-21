package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/OrIX219/marzgo/api/errors"
)

// HTTPClient performs http requests
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client represents http client coupled with admin credentials and token
type Client struct {
	APIBaseUrl string

	Username string
	Password string

	token      string
	tokenMutex sync.RWMutex

	refreshTokenTicker *time.Ticker
	refreshTokenErrs   chan error

	HTTPClient HTTPClient
}

// NewClient creates a new instance of Client
func NewClient(apiBase, username, password string, client HTTPClient) Client {
	return Client{
		APIBaseUrl: apiBase,
		Username:   username,
		Password:   password,
		HTTPClient: client,
	}
}

// RefreshTokenEvery starts a goroutine which will request a new admin token
// every duration and set it to current Client instance.
// If the goroutine was already started then it will just update the refresh period.
func (c *Client) RefreshTokenEvery(duration time.Duration) <-chan error {
	if c.refreshTokenTicker == nil {
		c.refreshTokenTicker = time.NewTicker(duration)
		c.refreshTokenErrs = make(chan error)
		go c.refreshToken(c.refreshTokenErrs)
	} else {
		c.refreshTokenTicker.Reset(duration)
	}
	return c.refreshTokenErrs
}

// StopRefreshingToken stops token refreshing.
func (c *Client) StopRefreshingToken() {
	c.refreshTokenTicker.Stop()
}

func (c *Client) refreshToken(errChan chan<- error) {
	for {
		select {
		case <-c.refreshTokenTicker.C:
			token, err := c.AdminToken()
			if err != nil {
				errChan <- err
			} else {
				c.SetToken(token.AccessToken)
			}
		}
	}
}

// SetToken sets an admin token to the current client instance in a concurrency safe way
func (c *Client) SetToken(token string) {
	c.tokenMutex.Lock()
	c.token = token
	c.tokenMutex.Unlock()
}

// Token returns an admin token stored in the current client instance
func (c *Client) Token() string {
	c.tokenMutex.RLock()
	tmp := c.token
	c.tokenMutex.RUnlock()
	return tmp
}

// A wrapper for MakeRequest.
// Request makes a request with given params to specified endpoint using specified method.
//
// Returns response in raw json.
func (c *Client) Request(method, endpoint string, params Params) (json.RawMessage, error) {
	return c.MakeRequest(method, endpoint, ContentTypeJson,
		params.Body(), params.Query(), params.Headers())
}

func (c *Client) addAuthHeader(req *http.Request) {
	c.tokenMutex.RLock()
	defer c.tokenMutex.RUnlock()
	if c.token == "" {
		return
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token))
}

// MakeRequest makes a request to specified endpoint using specified http method
// Request is populated with given reqBody, queryParams and headers
// reqBody is expected to return valid contents of contentType
//
// Returns response in raw json
func (c *Client) MakeRequest(method, endpoint string, contentType ContentType,
	reqBody RequestBody, queryParams QueryParams, headers Headers) (json.RawMessage, error) {
	url := fmt.Sprint(c.APIBaseUrl, endpoint)

	var body io.Reader
	if reqBody != nil {
		switch contentType {
		case ContentTypeJson:
			body = bytes.NewReader(reqBody.JSONEncoded())
		case ContentTypeUrlEncoded:
			body = strings.NewReader(reqBody.URLEncoded())
		}
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	c.addAuthHeader(req)
	req.Header.Set("Content-Type", contentType.String())
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	if queryParams != nil {
		req.URL.RawQuery = queryParams.Encode()
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	raw, err := decodeResponse(resp.Body)
	if err != nil {
		return nil, err
	}

	var respErr error
	if resp.StatusCode != http.StatusOK {
		switch resp.StatusCode {
		case http.StatusUnprocessableEntity:
			var ve errors.ValidationError
			err = json.Unmarshal(raw, &ve)
			respErr = ve
		case http.StatusForbidden:
			respErr = errors.NewForbidden(endpoint)
		case http.StatusNotFound:
			respErr = errors.NewNotFound(endpoint)
		case http.StatusConflict:
			respErr = errors.NewAlreadyExists()
		default:
			respErr = fmt.Errorf("%s", raw)
		}
		if err != nil {
			return nil, err
		}
	}
	if respErr != nil {
		return nil, respErr
	}

	return raw, nil
}

func decodeResponse(responseBody io.Reader) (resp json.RawMessage, err error) {
	dec := json.NewDecoder(responseBody)
	err = dec.Decode(&resp)
	return
}
