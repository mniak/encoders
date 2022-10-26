package encoders

import (
	"fmt"

	"github.com/mniak/encoders/encoders/core"
	"github.com/mniak/encoders/encoders/primitives"
	"github.com/pkg/errors"
)

type prefixedLength[T any] struct {
	lengthEncoder core.EncoderDecoder[uint]
	inner         core.EncoderDecoder[T]
}

func PrefixedLengthEncoder[T any](lengthSize byte, inner core.EncoderDecoder[T]) core.EncoderDecoder[T] {
	return prefixedLength[T]{
		lengthEncoder: primitives.FixedSizeNumber(lengthSize),
		inner:         inner,
	}
}

func (e prefixedLength[T]) Decode(state *T, data []byte) (int, error) {
	var length uint
	consumed, err := e.lengthEncoder.Decode(&length, data)
	if err != nil {
		return 0, errors.WithMessage(err, "failed to decode length")
	}
	data = data[consumed:]
	if len(data) < int(length) {
		return 0, fmt.Errorf("expecting %d bytes but got only %d", length, len(data))
	}
	data = data[:int(length)]

	_, err = e.inner.Decode(state, data)
	return consumed + int(length), err
}

func (e prefixedLength[T]) Encode(state T) ([]byte, error) {
	data, err := e.inner.Encode(state)
	if err != nil {
		return nil, err
	}

	lengthData, err := e.lengthEncoder.Encode(uint(len(data)))
	if err != nil {
		errors.WithMessage(err, "failed to encode length")
	}

	return append(lengthData, data...), nil
}
