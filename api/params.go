package api

import (
	"encoding/json"
	"net/url"
	"reflect"
	"strconv"
	"time"
)

type Params interface {
	Method() string
	Endpoint() string
	Params() (RequestParams, error)
	Headers() (RequestParams, error)
}

type RequestParams map[string]string

func (p RequestParams) AddString(key, value string) {
	p[key] = value
}

func (p RequestParams) AddStringNonEmpty(key, value string) {
	if value != "" {
		p[key] = value
	}
}

func (p RequestParams) AddInt(key string, value int) {
	p[key] = strconv.Itoa(value)
}

func (p RequestParams) AddIntNonZero(key string, value int) {
	if value != 0 {
		p[key] = strconv.Itoa(value)
	}
}

func (p RequestParams) AddInt64(key string, value int64) {
	p[key] = strconv.FormatInt(value, 10)
}

func (p RequestParams) AddInt64NonZero(key string, value int64) {
	if value != 0 {
		p[key] = strconv.FormatInt(value, 10)
	}
}

func (p RequestParams) AddFloat(key string, value float64) {
	p[key] = strconv.FormatFloat(value, 'f', 6, 64)
}

func (p RequestParams) AddFloatNonZero(key string, value float64) {
	if value != 0 {
		p[key] = strconv.FormatFloat(value, 'f', 6, 64)
	}
}

func (p RequestParams) AddBool(key string, value bool) {
	p[key] = strconv.FormatBool(value)
}

func (p RequestParams) AddTimeNonZero(key string, value time.Time) {
	zero := time.Time{}
	if value != zero {
		p[key] = value.String()
	}
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
