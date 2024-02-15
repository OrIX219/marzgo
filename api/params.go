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
	Body() (BodyParams, error)
	Url() (UrlParams, error)
	Headers() (Headers, error)
}

type Headers map[string]string

type BodyParams map[string]string

func (p BodyParams) AddString(key, value string) {
	p[key] = value
}

func (p BodyParams) AddStringNonEmpty(key, value string) {
	if value != "" {
		p[key] = value
	}
}

func (p BodyParams) AddStringSliceNonEmpty(key string, value []string) {
	if len(value) != 0 {
		// Ingoring error
		// What can go wrong here? Clueless
		b, _ := json.Marshal(value)
		p[key] = string(b)
	}
}

func (p BodyParams) AddInt(key string, value int) {
	p[key] = strconv.Itoa(value)
}

func (p BodyParams) AddIntNonZero(key string, value int) {
	if value != 0 {
		p[key] = strconv.Itoa(value)
	}
}

func (p BodyParams) AddIntSliceNonEmpty(key string, value []int) {
	if len(value) != 0 {
		// Ingoring error
		// What can go wrong here? Clueless
		b, _ := json.Marshal(value)
		p[key] = string(b)
	}
}

func (p BodyParams) AddInt64(key string, value int64) {
	p[key] = strconv.FormatInt(value, 10)
}

func (p BodyParams) AddInt64NonZero(key string, value int64) {
	if value != 0 {
		p[key] = strconv.FormatInt(value, 10)
	}
}

func (p BodyParams) AddInt64SliceNonEmpty(key string, value []int64) {
	if len(value) != 0 {
		// Ingoring error
		// What can go wrong here? Clueless
		b, _ := json.Marshal(value)
		p[key] = string(b)
	}
}

func (p BodyParams) AddFloat(key string, value float64) {
	p[key] = strconv.FormatFloat(value, 'f', 6, 64)
}

func (p BodyParams) AddFloatNonZero(key string, value float64) {
	if value != 0 {
		p[key] = strconv.FormatFloat(value, 'f', 6, 64)
	}
}

func (p BodyParams) AddFloatSliceNonEmpty(key string, value []float64) {
	if len(value) != 0 {
		// Ingoring error
		// What can go wrong here? Clueless
		b, _ := json.Marshal(value)
		p[key] = string(b)
	}
}

func (p BodyParams) AddBool(key string, value bool) {
	p[key] = strconv.FormatBool(value)
}

func (p BodyParams) AddBoolSliceNonEmpty(key string, value []bool) {
	if len(value) != 0 {
		// Ingoring error
		// What can go wrong here? Clueless
		b, _ := json.Marshal(value)
		p[key] = string(b)
	}
}

func (p BodyParams) AddTimeNonZero(key string, value time.Time) {
	zero := time.Time{}
	if value != zero {
		p[key] = value.String()
	}
}

func (p BodyParams) AddTimeSliceNonEmpty(key string, value []time.Time) {
	if len(value) != 0 {
		// Ingoring error
		// What can go wrong here? Clueless
		b, _ := json.Marshal(value)
		p[key] = string(b)
	}
}

func (p BodyParams) AddAny(key string, value any) error {
	if value == nil {
		return nil
	}
	val := reflect.ValueOf(value)
	if (val.Kind() == reflect.Ptr && val.IsNil()) ||
		(val.Kind() == reflect.Slice && val.IsZero()) {
		return nil
	}

	b, err := json.Marshal(value)
	if err != nil {
		return err
	}

	p[key] = string(b)

	return nil
}

func (p BodyParams) JSONEncoded() []byte {
	buf, _ := json.Marshal(p)
	return buf
}

func (p BodyParams) URLEncoded() string {
	values := url.Values{}
	for k, v := range p {
		values.Set(k, v)
	}
	return values.Encode()
}

type UrlParams map[string][]string

func (p UrlParams) Add(key, value string) {
	p[key] = append(p[key], value)
}

func (p UrlParams) AddString(key, value string) {
	p.Add(key, value)
}

func (p UrlParams) AddStringNonEmpty(key, value string) {
	if value != "" {
		p.Add(key, value)
	}
}

func (p UrlParams) AddStringSliceNonEmpty(key string, value []string) {
	if len(value) != 0 {
		p[key] = append(p[key], value...)
	}
}

func (p UrlParams) AddInt(key string, value int) {
	p.Add(key, strconv.Itoa(value))
}

func (p UrlParams) AddIntNonZero(key string, value int) {
	if value != 0 {
		p.Add(key, strconv.Itoa(value))
	}
}

func (p UrlParams) AddIntSliceNonEmpty(key string, value []int) {
	if len(value) != 0 {
		for _, v := range value {
			p.Add(key, strconv.Itoa(v))
		}
	}
}

func (p UrlParams) AddInt64(key string, value int64) {
	p.Add(key, strconv.FormatInt(value, 10))
}

func (p UrlParams) AddInt64NonZero(key string, value int64) {
	if value != 0 {
		p.Add(key, strconv.FormatInt(value, 10))
	}
}

func (p UrlParams) AddInt64SliceNonEmpty(key string, value []int64) {
	if len(value) != 0 {
		for _, v := range value {
			p.Add(key, strconv.FormatInt(v, 10))
		}
	}
}

func (p UrlParams) AddFloat(key string, value float64) {
	p.Add(key, strconv.FormatFloat(value, 'f', 6, 64))
}

func (p UrlParams) AddFloatNonZero(key string, value float64) {
	if value != 0 {
		p.Add(key, strconv.FormatFloat(value, 'f', 6, 64))
	}
}

func (p UrlParams) AddFloatSliceNonEmpty(key string, value []float64) {
	if len(value) != 0 {
		for _, v := range value {
			p.Add(key, strconv.FormatFloat(v, 'f', 6, 64))
		}
	}
}

func (p UrlParams) AddBool(key string, value bool) {
	p.Add(key, strconv.FormatBool(value))
}

func (p UrlParams) AddBoolSliceNonEmpty(key string, value []bool) {
	if len(value) != 0 {
		for _, v := range value {
			p.Add(key, strconv.FormatBool(v))
		}
	}
}

func (p UrlParams) AddTimeNonZero(key string, value time.Time) {
	zero := time.Time{}
	if value != zero {
		p.Add(key, value.String())
	}
}

func (p UrlParams) AddTimeSliceNonEmpty(key string, value []time.Time) {
	if len(value) != 0 {
		for i := range value {
			p.Add(key, value[i].String())
		}
	}
}

func (p UrlParams) Encode() string {
	return url.Values(p).Encode()
}
