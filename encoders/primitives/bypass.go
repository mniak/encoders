package primitives

import "github.com/mniak/encoders/encoders/core"

type bypassEncoder struct{}

func Bypass() core.EncoderDecoder[[]byte] {
	return bypassEncoder{}
}

func (e bypassEncoder) Decode(state *[]byte, data []byte) (int, error) {
	*state = data
	return len(data), nil
}

func (e bypassEncoder) Encode(state []byte) ([]byte, error) {
	return state, nil
}
