package tlv

import (
	"errors"
	"fmt"
)

type constantError string

const (
	ErrTagHasNoBytes       constantError = "tag has no bytes"
	ErrTagShouldHave2Bytes constantError = "tag should have 2 bytes"
	ErrTagTooLong          constantError = "tag too long"

	ErrLengthHasNoBytes     constantError = "length has no bytes"
	ErrLengthBiggerThanData constantError = "length bigger than data"
	ErrLengthFormatError    constantError = "length format error"

	ErrValueTooShort constantError = "value too short"
)

func (e constantError) Error() string {
	return string(e)
}

type TLVDecodeError struct {
	err        error
	constError constantError
}

func (e TLVDecodeError) Error() string {
	return e.err.Error()
}

func (e TLVDecodeError) Is(err error) bool {
	if err == e.constError {
		return true
	}
	return errors.Is(e.err, err)
}

func newTLVDecodeError(constError constantError, message string) TLVDecodeError {
	return TLVDecodeError{
		err:        errors.New(message),
		constError: constError,
	}
}

func newTLVDecodeErrorf(constError constantError, message string, args ...interface{}) TLVDecodeError {
	return TLVDecodeError{
		err:        fmt.Errorf(message, args...),
		constError: constError,
	}
}
