package primitives

import (
	"encoding/hex"

	"github.com/mniak/encoders/encoders/core"
	"github.com/mniak/encoders/encoders/internal"
)

func BCDString() core.EncoderDecoder[string] {
	return internal.Inline(
		func(state *string, data []byte) (int, error) {
			for idx, b := range data {
				high, low := int(b>>4), int(b&0x0f)
				if high > 9 || low > 9 {
					return 0, newErrorf(ErrByteNotBCD, "BCD decode: byte at index %d is invalid", idx)
				}
			}
			*state = hex.EncodeToString(data)
			return len(data), nil
		},
		func(state string) ([]byte, error) {
			result := []byte{}
			var high byte
			odd := len(state)%2 == 1

			for idx, digitRune := range state {
				digit := byte(digitRune - '0')
				if digit > 9 || digit < 0 {
					return nil, newErrorf(ErrNotADigit, "BCD encode: character at index %d is not a digit", idx)
				}

				odd = !odd
				if odd {
					high = digit
				} else {
					result = append(result, high<<4+digit)
				}
			}
			return result, nil
		},
	)
}
