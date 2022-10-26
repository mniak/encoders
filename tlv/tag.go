package tlv

import (
	"encoding/binary"
	"fmt"

	"github.com/mniak/encoders/core"
)

type Tag uint16

func (tag Tag) String() string {
	if tag > 0xFF {
		return fmt.Sprintf("%04X", uint16(tag))
	}
	return fmt.Sprintf("%02X", uint8(tag))
}

type tlvTagEncoder struct{}

func TagEncoder() core.EncoderDecoder[Tag] {
	return tlvTagEncoder{}
}

func (d tlvTagEncoder) Decode(state *Tag, data []byte) (int, error) {
	// If right-most 5 bits are 1 then consider that the tag has more than one byte
	if data[0]&0b0001_1111 != 0b0001_1111 {
		*state = Tag(data[0])
		return 1, nil
	}

	if len(data) < 2 {
		return 0, newTLVDecodeError(ErrTagShouldHave2Bytes, "expecting at least 2 bytes")
	}

	// If left-most 3 bytes are 0 then consider it is the last byte
	for data[1]&0b1110_0000 != 0b1110_0000 {
		*state = Tag(binary.BigEndian.Uint16(data))
		return 2, nil
	}

	return 0, newTLVDecodeError(ErrTagTooLong, "second byte of tag is not final, but VisaNet supports at most 2 bytes")
}

func (d tlvTagEncoder) Encode(state Tag) ([]byte, error) {
	var result [3]byte
	var count int
	for state > 0 {
		result[2-count] = byte(state % 256)
		state /= 256
		count++
	}
	return result[3-count:], nil
}
