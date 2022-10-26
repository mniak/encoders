package encoders

import "github.com/mniak/encoders/core"

type lengthConstraintEncoder[T any] struct {
	inner    core.EncoderDecoder[T]
	maxBytes int
}

func MaxLength[T any](maxBytes int, inner core.EncoderDecoder[T]) core.EncoderDecoder[T] {
	return lengthConstraintEncoder[T]{
		maxBytes: maxBytes,
		inner:    inner,
	}
}

func (e lengthConstraintEncoder[T]) Encode(state T) ([]byte, error) {
	return e.inner.Encode(state)
}

func (e lengthConstraintEncoder[T]) Decode(state *T, data []byte) (int, error) {
	return e.inner.Decode(state, data)
}
