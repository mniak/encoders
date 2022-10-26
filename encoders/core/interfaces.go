package core

type Encoder[T any] interface {
	Encode(state T) ([]byte, error)
}
type Decoder[T any] interface {
	Decode(state *T, data []byte) (int, error)
}

type EncoderDecoder[T any] interface {
	Encoder[T]
	Decoder[T]
}
