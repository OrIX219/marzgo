package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/OrIX219/marzgo/api/errors"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	APIBaseUrl string

	Username string
	Password string
	Token    string

	HTTPClient HTTPClient
}

func (c *Client) Request(params Params) (json.RawMessage, error) {
	reqBody, err := params.Body()
	if err != nil {
		return nil, err
	}
	queryParams, err := params.Query()
	if err != nil {
		return nil, err
	}
	reqHeaders, err := params.Headers()
	if err != nil {
		return nil, err
	}
	return c.MakeRequest(params.Method(), params.Endpoint(), ContentTypeJson,
		reqBody, queryParams, reqHeaders)
}

func (c *Client) addAuthHeader(req *http.Request) {
	if c.Token == "" {
		return
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.Token))
}

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
