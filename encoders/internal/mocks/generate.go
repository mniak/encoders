package mocks

type EncoderDecoderOfString interface {
	Encode(state string) ([]byte, error)
	Decode(state *string, data []byte) (int, error)
}

type EncoderDecoderOfInt interface {
	Encode(state int) ([]byte, error)
	Decode(state *int, data []byte) (int, error)
}

//go:generate mockgen --package=mocks --destination=encoderdecoder_mock.go --source generate.go EncoderDecoderOfInt,EncoderDecoderOfString
