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
	Body() (RequestBody, error)
	Query() (QueryParams, error)
	Headers() (Headers, error)
}

type Headers map[string]string

type RequestBody map[string]string

func (p RequestBody) AddString(key, value string) {
	p[key] = value
}

func (p RequestBody) AddStringNonEmpty(key, value string) {
	if value != "" {
		p[key] = value
	}
}

func (p RequestBody) AddStringSliceNonEmpty(key string, value []string) {
	if len(value) != 0 {
		// Ingoring error
		// What can go wrong here? Clueless
		b, _ := json.Marshal(value)
		p[key] = string(b)
	}
}

func (p RequestBody) AddInt(key string, value int) {
	p[key] = strconv.Itoa(value)
}

func (p RequestBody) AddIntNonZero(key string, value int) {
	if value != 0 {
		p[key] = strconv.Itoa(value)
	}
}

func (p RequestBody) AddIntSliceNonEmpty(key string, value []int) {
	if len(value) != 0 {
		// Ingoring error
		// What can go wrong here? Clueless
		b, _ := json.Marshal(value)
		p[key] = string(b)
	}
}

func (p RequestBody) AddInt64(key string, value int64) {
	p[key] = strconv.FormatInt(value, 10)
}

func (p RequestBody) AddInt64NonZero(key string, value int64) {
	if value != 0 {
		p[key] = strconv.FormatInt(value, 10)
	}
}

func (p RequestBody) AddInt64SliceNonEmpty(key string, value []int64) {
	if len(value) != 0 {
		// Ingoring error
		// What can go wrong here? Clueless
		b, _ := json.Marshal(value)
		p[key] = string(b)
	}
}

func (p RequestBody) AddFloat(key string, value float64) {
	p[key] = strconv.FormatFloat(value, 'f', 6, 64)
}

func (p RequestBody) AddFloatNonZero(key string, value float64) {
	if value != 0 {
		p[key] = strconv.FormatFloat(value, 'f', 6, 64)
	}
}

func (p RequestBody) AddFloatSliceNonEmpty(key string, value []float64) {
	if len(value) != 0 {
		// Ingoring error
		// What can go wrong here? Clueless
		b, _ := json.Marshal(value)
		p[key] = string(b)
	}
}

func (p RequestBody) AddBool(key string, value bool) {
	p[key] = strconv.FormatBool(value)
}

func (p RequestBody) AddBoolSliceNonEmpty(key string, value []bool) {
	if len(value) != 0 {
		// Ingoring error
		// What can go wrong here? Clueless
		b, _ := json.Marshal(value)
		p[key] = string(b)
	}
}

func (p RequestBody) AddTimeNonZero(key string, value time.Time) {
	zero := time.Time{}
	if value != zero {
		p[key] = value.String()
	}
}

func (p RequestBody) AddTimeSliceNonEmpty(key string, value []time.Time) {
	if len(value) != 0 {
		// Ingoring error
		// What can go wrong here? Clueless
		b, _ := json.Marshal(value)
		p[key] = string(b)
	}
}

func (p RequestBody) AddAny(key string, value any) error {
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

func (p RequestBody) JSONEncoded() []byte {
	buf, _ := json.Marshal(p)
	return buf
}

func (p RequestBody) URLEncoded() string {
	values := url.Values{}
	for k, v := range p {
		values.Set(k, v)
	}
	return values.Encode()
}

type QueryParams map[string][]string

func (p QueryParams) Add(key, value string) {
	p[key] = append(p[key], value)
}

func (p QueryParams) AddString(key, value string) {
	p.Add(key, value)
}

func (p QueryParams) AddStringNonEmpty(key, value string) {
	if value != "" {
		p.Add(key, value)
	}
}

func (p QueryParams) AddStringSliceNonEmpty(key string, value []string) {
	if len(value) != 0 {
		p[key] = append(p[key], value...)
	}
}

func (p QueryParams) AddInt(key string, value int) {
	p.Add(key, strconv.Itoa(value))
}

func (p QueryParams) AddIntNonZero(key string, value int) {
	if value != 0 {
		p.Add(key, strconv.Itoa(value))
	}
}

func (p QueryParams) AddIntSliceNonEmpty(key string, value []int) {
	if len(value) != 0 {
		for _, v := range value {
			p.Add(key, strconv.Itoa(v))
		}
	}
}

func (p QueryParams) AddInt64(key string, value int64) {
	p.Add(key, strconv.FormatInt(value, 10))
}

func (p QueryParams) AddInt64NonZero(key string, value int64) {
	if value != 0 {
		p.Add(key, strconv.FormatInt(value, 10))
	}
}

func (p QueryParams) AddInt64SliceNonEmpty(key string, value []int64) {
	if len(value) != 0 {
		for _, v := range value {
			p.Add(key, strconv.FormatInt(v, 10))
		}
	}
}

func (p QueryParams) AddFloat(key string, value float64) {
	p.Add(key, strconv.FormatFloat(value, 'f', 6, 64))
}

func (p QueryParams) AddFloatNonZero(key string, value float64) {
	if value != 0 {
		p.Add(key, strconv.FormatFloat(value, 'f', 6, 64))
	}
}

func (p QueryParams) AddFloatSliceNonEmpty(key string, value []float64) {
	if len(value) != 0 {
		for _, v := range value {
			p.Add(key, strconv.FormatFloat(v, 'f', 6, 64))
		}
	}
}

func (p QueryParams) AddBool(key string, value bool) {
	p.Add(key, strconv.FormatBool(value))
}

func (p QueryParams) AddBoolSliceNonEmpty(key string, value []bool) {
	if len(value) != 0 {
		for _, v := range value {
			p.Add(key, strconv.FormatBool(v))
		}
	}
}

func (p QueryParams) AddTimeNonZero(key string, value time.Time) {
	zero := time.Time{}
	if value != zero {
		p.Add(key, value.String())
	}
}

func (p QueryParams) AddTimeSliceNonEmpty(key string, value []time.Time) {
	if len(value) != 0 {
		for i := range value {
			p.Add(key, value[i].String())
		}
	}
}

func (p QueryParams) Encode() string {
	return url.Values(p).Encode()
}
