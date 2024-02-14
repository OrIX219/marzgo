package models

import (
	"encoding/json"
	"strings"
	"time"
)

type Time time.Time

func (t *Time) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	time, err := time.Parse("2006-01-02T15:04:05.999999999", s)
	if err != nil {
		return err
	}
	*t = Time(time)
	return nil
}

func (t Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(t))
}

func (t Time) Format(s string) string {
	return time.Time(t).Format(s)
}

func (t Time) String() string {
	return time.Time(t).String()
}
