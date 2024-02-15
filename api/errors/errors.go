package errors

import (
	"fmt"
)

type ValidationError struct {
	Detail map[string]string `json:"detail"`
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("%s", e.Detail)
}

type ForbiddenError struct {
	endpoint string
}

func NewForbidden(endpoint string) ForbiddenError {
	return ForbiddenError{
		endpoint: endpoint,
	}
}

func (e ForbiddenError) Error() string {
	return fmt.Sprintf("Forbidden: %s", e.endpoint)
}

type NotFoundError struct {
	endpoint string
}

func NewNotFound(endpoint string) NotFoundError {
	return NotFoundError{
		endpoint: endpoint,
	}
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("Not Found: %s", e.endpoint)
}

type AlreadyExistsError struct{}

func NewAlreadyExists() AlreadyExistsError {
	return AlreadyExistsError{}
}

func (e AlreadyExistsError) Error() string {
	return fmt.Sprintf("Alredy Exists")
}
