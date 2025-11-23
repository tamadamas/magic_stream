package app_errors

import "net/http"

type HTTPError interface {
	error
	Status() int
}

// 404 error
type NotFoundError struct {
	Msg string
	Err error
}

func (e *NotFoundError) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}
	return e.Msg
}

func (e *NotFoundError) Status() int {
	return http.StatusNotFound
}

func NewNotFoundError(err error, msg string) *NotFoundError {
	return &NotFoundError{Err: err, Msg: msg}
}
