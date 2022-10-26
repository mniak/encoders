package primitives

import (
	"github.com/mniak/encoders/encoders/core"
	"github.com/mniak/encoders/encoders/internal"
)

func BCDNumber() core.EncoderDecoder[int] {
	return internal.Inline(
		func(state *int, data []byte) (int, error) {
			result := 0
			for idx, b := range data {
				high, low := int(b>>4), int(b&0x0f)
				if high > 9 || low > 9 {
					return 0, newErrorf(ErrByteNotBCD, "BCD decode: byte at index %d is invalid", idx)
				}
				result *= 100
				result += 10*high + low
			}
			*state = result
			return len(data), nil
		},
		func(state int) ([]byte, error) {
			if state <= 0 {
				return []byte{0}, nil
			}

			const arraySize = 10
			var result [arraySize]byte
			idx := arraySize - 1
			for state > 0 {
				hl := state % 100
				state /= 100

				high, low := hl/10, hl%10
				b := high<<4 + low
				result[idx] = byte(b)

				idx--
			}
			return result[idx+1:], nil
		},
	)
}
