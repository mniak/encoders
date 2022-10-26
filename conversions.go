package encoders

import (
	"github.com/mniak/encoders/core"
	"github.com/mniak/encoders/internal"
)

type numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

func AsInt[A numeric, B numeric](intEncoder core.EncoderDecoder[B]) core.EncoderDecoder[A] {
	return internal.Inline(
		func(state *A, data []byte) (int, error) {
			var tmp B
			read, err := intEncoder.Decode(&tmp, data)
			*state = A(tmp)
			return read, err
		},
		func(state A) ([]byte, error) {
			return intEncoder.Encode(B(state))
		},
	)
}
