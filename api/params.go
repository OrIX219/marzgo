package api

import (
	"net/url"
	"strconv"
	"time"
)

type Params interface {
	Body() (RequestBody, error)
	Query() (QueryParams, error)
	Headers() (Headers, error)
}

type Headers map[string]string

type RequestBody interface {
	JSONEncoded() []byte
	URLEncoded() string
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
