package api

import (
	"encoding/json"
	"net/url"
	"reflect"
	"strconv"
)

type Params interface {
	Method() string
	Endpoint() string
	Params() (RequestParams, error)
}

type RequestParams map[string]string

func (p RequestParams) AddNonEmpty(key, value string) {
	if value != "" {
		p[key] = value
	}
}

func (p RequestParams) AddNonZero(key string, value int) {
	if value != 0 {
		p[key] = strconv.Itoa(value)
	}
}

func (p RequestParams) AddNonZero64(key string, value int64) {
	if value != 0 {
		p[key] = strconv.FormatInt(value, 10)
	}
}

func (p RequestParams) AddNonZeroFloat(key string, value float64) {
	if value != 0 {
		p[key] = strconv.FormatFloat(value, 'f', 6, 64)
	}
}

func (p RequestParams) AddBool(key string, value bool) {
	p[key] = strconv.FormatBool(value)
}

func (p RequestParams) AddAny(key string, value any) error {
	if value == nil || (reflect.ValueOf(value).Kind() == reflect.Ptr && reflect.ValueOf(value).IsNil()) {
		return nil
	}

	b, err := json.Marshal(value)
	if err != nil {
		return err
	}

	p[key] = string(b)

	return nil
}

func (p RequestParams) URLEncoded() string {
	values := url.Values{}
	for k, v := range p {
		values.Set(k, v)
	}
	return values.Encode()
}

func (p RequestParams) JSONEncoded() []byte {
	buf, _ := json.Marshal(p)
	return buf
}
