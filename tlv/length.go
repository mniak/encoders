package tlv

import (
	"encoding/binary"

	"github.com/mniak/encoders/core"
)

type Length uint

type tlvLengthEncoder struct{}

func LengthEncoder() core.EncoderDecoder[Length] {
	return tlvLengthEncoder{}
}

func (d tlvLengthEncoder) Decode(state *Length, data []byte) (int, error) {
	if len(data) == 0 {
		return 0, newTLVDecodeError(ErrLengthHasNoBytes, "no bytes found")
	}

	// If left-most 1 bit is 0, then it is the Short Form, one single byte
	if data[0]&0b1000_0000 == 0 {
		*state = Length(data[0])
		return 1, nil
	}

	howMuchMore := data[0] & 0b0111_1111
	switch howMuchMore {
	case 1:
		if len(data) < 2 {
			return 0, newTLVDecodeError(ErrLengthFormatError, "expecting at least 2 bytes")
		}
		*state = Length(data[1])
		return 2, nil
	case 2:
		if len(data) < 3 {
			return 0, newTLVDecodeError(ErrLengthFormatError, "expecting at least 3 bytes")
		}
		*state = Length(binary.BigEndian.Uint16(data[1:]))
		return 3, nil
	default:
		return 0, newTLVDecodeErrorf(ErrLengthFormatError, "the length field should have at most 3 bytes but it appears to have %d", 1+howMuchMore)
	}
}

func (d tlvLengthEncoder) Encode(state Length) ([]byte, error) {
	if state < 128 {
		return []byte{byte(state)}, nil
	}

	var result [9]byte
	var count int
	for state > 0 {
		result[8-count] = byte(state % 256)
		state /= 256
		count++
	}
	result[8-count] = byte(0x80 + count)
	return result[8-count:], nil
}
