package goa_errors

import "net/http"

type Error struct {
	Message string `json:"error"`
	status  int
}

func (e *Error) StatusCode() int {
	return e.status
}

func (e *Error) Error() string {
	return "application error"
}

func NewBadRequestError(err error) *Error {
	return &Error{
		Message: err.Error(),
		status:  http.StatusBadRequest,
	}
}

func NewInternalServerError(err error) *Error {
	return &Error{
		Message: err.Error(),
		status:  http.StatusInternalServerError,
	}
}
