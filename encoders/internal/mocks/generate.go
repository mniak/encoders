package mocks

//go:generate mockgen --package=mocks --destination=encoderdecoder_string_mock.go --source generate.go EncoderDecoderOfString
type EncoderDecoderOfString interface {
	Encode(state string) ([]byte, error)
	Decode(state *string, data []byte) (int, error)
}

//go:generate mockgen --package=mocks --destination=encoderdecoder_int_mock.go --source generate.go EncoderDecoderOfInt
type EncoderDecoderOfInt interface {
	Encode(state int) ([]byte, error)
	Decode(state *int, data []byte) (int, error)
}
