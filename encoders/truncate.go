package encoders

import (
	"errors"

	"github.com/mniak/encoders/encoders/core"
	"github.com/mniak/encoders/encoders/internal"
)

func TruncateStringPadLeft(size byte, padchar rune, inner core.EncoderDecoder[string]) core.EncoderDecoder[string] {
	return internal.Inline[string](
		func(state *string, data []byte) (int, error) {
			return 0, nil
		},
		func(state string) ([]byte, error) {
			return nil, errors.New("TODO: not implemented (truncate)")
		},
	)
}
