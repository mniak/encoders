package primitives

import (
	"github.com/gdumoulin/goebcdic"
	"github.com/mniak/encoders/core"
	"github.com/mniak/encoders/internal"
)

func AlphaNumeric() core.EncoderDecoder[string] {
	return AlphaNumericAndSymbols()
}

func AlphaNumericAndSymbols() core.EncoderDecoder[string] {
	return internal.Inline(
		func(state *string, data []byte) (int, error) {
			if len(data) == 0 {
				return 0, nil
			}

			x := goebcdic.EBCDICtoASCIIofBytes(data)
			*state = string(x)
			return int(len(data)), nil
		},
		func(state string) ([]byte, error) {
			return goebcdic.ASCIItoEBCDICofBytes([]byte(state)), nil
		},
	)
}
