package restserver

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrTypeAssertionError = errors.New("unable to assert type")
)

type ParsingError struct {
	Err error
}

func (e *ParsingError) Unwrap() error {
	return e.Err
}

func (e *ParsingError) Error() string {
	return e.Err.Error()
}

type RequiredError struct {
	Field string
}

func (e *RequiredError) Error() string {
	return fmt.Sprintf("required field '%s' is zero value.", e.Field)
}

type ErrorHandler func(w http.ResponseWriter, r *http.Request, err error, result *ImplResponse)

func DefaultErrorHandler(w http.ResponseWriter, r *http.Request, err error, result *ImplResponse) {
	if _, ok := err.(*ParsingError); ok {
		EncodeJSONResponse(err.Error(), func(i int) *int { return &i }(http.StatusBadRequest), w)
	} else if _, ok := err.(*RequiredError); ok {
		EncodeJSONResponse(err.Error(), func(i int) *int { return &i }(http.StatusUnprocessableEntity), w)
	} else {
		EncodeJSONResponse(err.Error(), &result.Code, w)
	}
}
