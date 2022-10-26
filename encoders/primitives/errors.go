package primitives

import (
	"errors"
	"fmt"
)

type constantError string

const (
	ErrByteNotBCD constantError = "byte not in BCD format"
	ErrNotADigit  constantError = "character is not a digit"
)

func (e constantError) Error() string {
	return string(e)
}

type DataReadError struct {
	err        error
	constError constantError
}

func (e DataReadError) Error() string {
	return e.err.Error()
}

func (e DataReadError) Is(err error) bool {
	if err == e.constError {
		return true
	}
	return errors.Is(e.err, err)
}

func newError(constError constantError, message string) DataReadError {
	return DataReadError{
		err:        errors.New(message),
		constError: constError,
	}
}

func newErrorf(constError constantError, message string, args ...interface{}) DataReadError {
	return DataReadError{
		err:        fmt.Errorf(message, args...),
		constError: constError,
	}
}
