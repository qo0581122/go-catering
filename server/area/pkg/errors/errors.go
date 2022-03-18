package errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var (
	databaseRowsNotFound = errors.New("database rows not found")
)

// Package errors provides a way to return detailed information
// for an RPC request error. The error is normally JSON encoded.

// Error implements the error interface.
type Error struct {
	Code   int32  `json:"code"`
	Detail string `json:"detail"`
	Status string `json:"status"`
}

func (e *Error) Error() string {
	b, _ := json.Marshal(e)
	return string(b)
}

func (e *Error) GetCode() int32 {
	return e.Code
}

func (e *Error) GetDetail() string {
	return e.Detail
}

// New generates a custom error.
func New(detail string, code int32) error {
	return &Error{
		Code:   code,
		Detail: detail,
		Status: http.StatusText(int(code)),
	}
}

// Parse tries to parse a JSON string into an error. If that
// fails, it will set the given string as the error detail.
func Parse(err string) *Error {
	e := new(Error)
	errr := json.Unmarshal([]byte(err), e)
	if errr != nil {
		e.Detail = err
	}
	return e
}

// BadRequest generates a 400 error.
func BadRequest(format string, a ...interface{}) error {
	return &Error{
		Code:   http.StatusBadRequest,
		Detail: fmt.Sprintf(format, a...),
		Status: http.StatusText(http.StatusBadRequest),
	}
}

// Unauthorized generates a 401 error.
func Unauthorized(format string, a ...interface{}) error {
	return &Error{
		Code:   http.StatusUnauthorized,
		Detail: fmt.Sprintf(format, a...),
		Status: http.StatusText(http.StatusUnauthorized),
	}
}

// Forbidden generates a 403 error.
func Forbidden(format string, a ...interface{}) error {
	return &Error{
		Code:   http.StatusForbidden,
		Detail: fmt.Sprintf(format, a...),
		Status: http.StatusText(http.StatusForbidden),
	}
}

// NotFound generates a 404 error.
func NotFound(format string, a ...interface{}) error {
	return &Error{
		Code:   http.StatusNotFound,
		Detail: fmt.Sprintf(format, a...),
		Status: http.StatusText(http.StatusNotFound),
	}
}

// MethodNotAllowed generates a 405 error.
func MethodNotAllowed(format string, a ...interface{}) error {
	return &Error{
		Code:   http.StatusMethodNotAllowed,
		Detail: fmt.Sprintf(format, a...),
		Status: http.StatusText(http.StatusMethodNotAllowed),
	}
}

// Timeout generates a 408 error.
func Timeout(format string, a ...interface{}) error {
	return &Error{
		Code:   http.StatusRequestTimeout,
		Detail: fmt.Sprintf(format, a...),
		Status: http.StatusText(http.StatusRequestTimeout),
	}
}

// Conflict generates a 409 error.
func Conflict(format string, a ...interface{}) error {
	return &Error{
		Code:   http.StatusConflict,
		Detail: fmt.Sprintf(format, a...),
		Status: http.StatusText(http.StatusConflict),
	}
}

// InternalServerError generates a 500 error.
func InternalServerError(format string, a ...interface{}) error {
	return &Error{
		Code:   http.StatusInternalServerError,
		Detail: fmt.Sprintf(format, a...),
		Status: http.StatusText(http.StatusInternalServerError),
	}
}
