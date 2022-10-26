package primitives

import (
	"strconv"

	"github.com/gdumoulin/goebcdic"
	"github.com/mniak/encoders/core"
	"github.com/mniak/encoders/internal"
)

func EBCDICNumber() core.EncoderDecoder[int] {
	return internal.Inline(
		func(state *int, data []byte) (int, error) {
			if len(data) == 0 {
				return 0, nil
			}
			var err error
			x := goebcdic.EBCDICtoASCIIofBytes(data)
			*state, err = strconv.Atoi(string(x))
			return int(len(data)), err
		},
		func(state int) ([]byte, error) {
			str := strconv.Itoa(state)
			return goebcdic.ASCIItoEBCDICofBytes([]byte(str)), nil
		},
	)
}
