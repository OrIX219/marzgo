package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
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
	reqParams, err := params.Params()
	if err != nil {
		return nil, err
	}
	reqHeaders, err := params.Headers()
	if err != nil {
		return nil, err
	}
	return c.MakeRequest(params.Method(), params.Endpoint(), ContentTypeJson, reqParams, reqHeaders)
}

func (c *Client) addAuthHeader(req *http.Request) {
	if c.Token == "" {
		return
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.Token))
}

func (c *Client) MakeRequest(method, endpoint string,
	contentType ContentType, params, headers RequestParams) (json.RawMessage, error) {
	url := fmt.Sprint(c.APIBaseUrl, endpoint)

	var body io.Reader
	if method == "POST" || method == "PUT" {
		switch contentType {
		case ContentTypeJson:
			body = bytes.NewReader(params.JSONEncoded())
		case ContentTypeUrlEncoded:
			body = strings.NewReader(params.URLEncoded())
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
	if method == "GET" {
		req.URL.RawQuery = params.URLEncoded()
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

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s", raw)
	}

	return raw, nil
}

func decodeResponse(responseBody io.Reader) (resp json.RawMessage, err error) {
	dec := json.NewDecoder(responseBody)
	err = dec.Decode(&resp)
	return
}
