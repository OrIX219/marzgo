package api

import (
	"net/url"
	"strconv"
	"time"
)

// Params provides access to request's body, URL query parameters and headers
type Params interface {
	Body() RequestBody
	Query() QueryParams
	Headers() Headers
}

// Headers represents a set of request headers
type Headers map[string]string

// RequestBody provides access to it's contents in JSON and/or URL encoded format
type RequestBody interface {
	JSONEncoded() []byte
	URLEncoded() string
}

// QueryParams represents a set of URL query parameters
type QueryParams map[string][]string

// Add adds a string value to a key
func (p QueryParams) Add(key, value string) {
	p[key] = append(p[key], value)
}

// AddString adds a string value
func (p QueryParams) AddString(key, value string) {
	p.Add(key, value)
}

// AddStringNonEmpty adds a string value if it's not empty
func (p QueryParams) AddStringNonEmpty(key, value string) {
	if value != "" {
		p.Add(key, value)
	}
}

// AddStringSliceNonEmpty adds a string slice if it's not empty
func (p QueryParams) AddStringSliceNonEmpty(key string, value []string) {
	if len(value) != 0 {
		p[key] = append(p[key], value...)
	}
}

// AddInt adds an int value
func (p QueryParams) AddInt(key string, value int) {
	p.Add(key, strconv.Itoa(value))
}

// AddIntNonZero adds an int value if it's not zero
func (p QueryParams) AddIntNonZero(key string, value int) {
	if value != 0 {
		p.Add(key, strconv.Itoa(value))
	}
}

// AddIntSliceNonEmpty adds an int slice value if it's not empty
func (p QueryParams) AddIntSliceNonEmpty(key string, value []int) {
	if len(value) != 0 {
		for _, v := range value {
			p.Add(key, strconv.Itoa(v))
		}
	}
}

// AddInt64 adds an int64 value
func (p QueryParams) AddInt64(key string, value int64) {
	p.Add(key, strconv.FormatInt(value, 10))
}

// AddInt64NonZero adds an int64 value if it's not zero
func (p QueryParams) AddInt64NonZero(key string, value int64) {
	if value != 0 {
		p.Add(key, strconv.FormatInt(value, 10))
	}
}

// AddInt64SliceNonEmpty adds an int64 slice value if it's not empty
func (p QueryParams) AddInt64SliceNonEmpty(key string, value []int64) {
	if len(value) != 0 {
		for _, v := range value {
			p.Add(key, strconv.FormatInt(v, 10))
		}
	}
}

// AddFloat adds a float value
func (p QueryParams) AddFloat(key string, value float64) {
	p.Add(key, strconv.FormatFloat(value, 'f', 6, 64))
}

// AddFloatNonZero adds a float value if it's not zero
func (p QueryParams) AddFloatNonZero(key string, value float64) {
	if value != 0 {
		p.Add(key, strconv.FormatFloat(value, 'f', 6, 64))
	}
}

// AddFloatSliceNonEmpty adds a float slice value if it's not empty
func (p QueryParams) AddFloatSliceNonEmpty(key string, value []float64) {
	if len(value) != 0 {
		for _, v := range value {
			p.Add(key, strconv.FormatFloat(v, 'f', 6, 64))
		}
	}
}

// AddBool adds a bool value
func (p QueryParams) AddBool(key string, value bool) {
	p.Add(key, strconv.FormatBool(value))
}

// AddBoolSliceNonEmpty adds a bool slice value if it's not empty
func (p QueryParams) AddBoolSliceNonEmpty(key string, value []bool) {
	if len(value) != 0 {
		for _, v := range value {
			p.Add(key, strconv.FormatBool(v))
		}
	}
}

// AddTimeNonZero adds a time.Time value if it's not zero (time.Time{})
func (p QueryParams) AddTimeNonZero(key string, value time.Time) {
	zero := time.Time{}
	if value != zero {
		p.Add(key, value.String())
	}
}

// AddTimeSliceNonEmpty adds a time.Time slice value if it's not empty
func (p QueryParams) AddTimeSliceNonEmpty(key string, value []time.Time) {
	if len(value) != 0 {
		for i := range value {
			p.Add(key, value[i].String())
		}
	}
}

// Encode returns query params in URL-encoded format
func (p QueryParams) Encode() string {
	return url.Values(p).Encode()
}
