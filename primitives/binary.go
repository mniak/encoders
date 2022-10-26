package primitives

import (
	"encoding/binary"
	"errors"
	"math"

	"github.com/mniak/encoders/core"
	"github.com/mniak/encoders/internal"
)

func Binary() core.EncoderDecoder[int] {
	return internal.Inline[int](
		func(state *int, data []byte) (int, error) {
			totalLength := len(data)
			var result int
			for len(data) > 0 {
				head := data[0]
				data = data[1:]
				result *= 256
				result += int(head)
				if result < 0 {
					*state = result
					return totalLength - len(data), errors.New("int overflow")
				}
			}
			*state = result
			return totalLength, nil
		},
		func(state int) ([]byte, error) {
			if state <= math.MaxUint8 {
				return []byte{byte(state)}, nil
			} else if state <= math.MaxUint16 {
				var result [2]byte
				binary.BigEndian.PutUint16(result[:], uint16(state))
				return result[:], nil
			} else if state <= math.MaxUint32 {
				var result [4]byte
				binary.BigEndian.PutUint32(result[:], uint32(state))
				return result[:], nil
			} else {
				var result [8]byte
				binary.BigEndian.PutUint64(result[:], uint64(state))
				return result[:], nil
			}
		},
	)
}
