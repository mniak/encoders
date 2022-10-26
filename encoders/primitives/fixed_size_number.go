package primitives

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math"

	"github.com/mniak/encoders/encoders/core"
)

type uintEncoder struct {
	size byte
}

func FixedSizeNumber(size byte) core.EncoderDecoder[uint] {
	return uintEncoder{
		size: size,
	}
}

func (e uintEncoder) Decode(state *uint, data []byte) (int, error) {
	if len(data) < int(e.size) {
		return 0, fmt.Errorf("expecting %d bytes but got %d", e.size, len(data))
	}
	switch e.size {
	case 1:
		*state = uint(data[0])
		return 1, nil
	case 2:
		*state = uint(binary.BigEndian.Uint16(data))
		return 2, nil
	}

	return 0, fmt.Errorf("invalid integer size: %d", e.size)
}

func (e uintEncoder) Encode(state uint) ([]byte, error) {
	if state > math.MaxUint16 {
		fmt.Errorf("cannot represent the value in %d bytes", e.size)
	}
	switch e.size {
	case 0:
		return nil, errors.New("impossible to represent anything in zero bytes")
	case 1:
		return []byte{byte(state)}, nil
	default:
		result := make([]byte, e.size)
		binary.BigEndian.PutUint16(result, uint16(state))
		return result, nil
	}

	return nil, fmt.Errorf("invalid integer size: %d", e.size)
}
