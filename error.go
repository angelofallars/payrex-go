package payrex

import (
	"errors"
	"fmt"
)

var ErrNilParams = errors.New("expected params argument to not be nil")

type Error struct {
	Errors []ErrorMessage `json:"errors"`
}

func (e Error) Error() string {
	s := ""
	for i, err := range e.Errors {
		s += fmt.Sprintf("err %d:\n", i+1)
		s += err.Error()
	}
	return s
}

func (e *Error) Unwrap() []error {
	errors := make([]error, len(e.Errors))
	for _, err := range e.Errors {
		errors = append(errors, err)
	}
	return errors
}

type ErrorMessage struct {
	Code      string `json:"code"`
	Detail    string `json:"detail"`
	Parameter string `json:"parameter"`
}

func (em ErrorMessage) Error() string {
	s := ""
	s += "code: '" + em.Code + "'\n"
	s += "detail: '" + em.Detail + "'\n"
	s += "parameter: '" + em.Parameter + "'\n"
	return s
}
