package primitives

import (
	"encoding/hex"
	"fmt"

	"github.com/mniak/encoders/core"
	"github.com/mniak/encoders/internal"
)

func Hexa() core.EncoderDecoder[string] {
	return internal.Inline(
		func(state *string, data []byte) (int, error) {
			*state = fmt.Sprintf("%X", data)
			return len(data), nil
		},
		func(state string) ([]byte, error) {
			return hex.DecodeString(state)
		},
	)
}
