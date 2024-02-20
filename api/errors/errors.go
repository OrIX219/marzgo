package errors

import (
	"fmt"
)

// ValidationError represents Marzban's response body for error 422
type ValidationError struct {
	Detail map[string]string `json:"detail"`
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("%s", e.Detail)
}

// ForbiddenError represents http error 403 coupled with endpoint
type ForbiddenError struct {
	endpoint string
}

// Creates new ForbiddenError
func NewForbidden(endpoint string) ForbiddenError {
	return ForbiddenError{
		endpoint: endpoint,
	}
}

func (e ForbiddenError) Error() string {
	return fmt.Sprintf("Forbidden: %s", e.endpoint)
}

// NotFoundError represents http error 404 coupled with endpoint
type NotFoundError struct {
	endpoint string
}

// Creates new NotFoundError
func NewNotFound(endpoint string) NotFoundError {
	return NotFoundError{
		endpoint: endpoint,
	}
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("Not Found: %s", e.endpoint)
}

// AlreadyExistsError represents http error 409
type AlreadyExistsError struct{}

// Creates new AlreadyExistsError
func NewAlreadyExists() AlreadyExistsError {
	return AlreadyExistsError{}
}

func (e AlreadyExistsError) Error() string {
	return fmt.Sprintf("Already Exists")
}
