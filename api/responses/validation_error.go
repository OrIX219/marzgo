package responses

import "fmt"

type ValidationErrors struct {
	Details []ValidationError `json:"detail"`
}

func (e ValidationErrors) Error() string {
	return fmt.Sprint(e.Details)
}

type ValidationError struct {
	Loc  []string `json:"loc"`
	Msg  string   `json:"msg"`
	Type string   `json:"type"`
}

func (e ValidationError) Error() string {
	return fmt.Sprint(e.Loc, e.Msg, e.Type)
}
