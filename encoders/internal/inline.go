package internal

type inlineEncoder[T any] struct {
	fnDecode func(state *T, data []byte) (int, error)
	fnEncode func(state T) ([]byte, error)
}

func Inline[T any](
	fnDecode func(state *T, data []byte) (int, error),
	fnEncode func(state T) ([]byte, error),
) inlineEncoder[T] {
	return inlineEncoder[T]{
		fnDecode: fnDecode,
		fnEncode: fnEncode,
	}
}

func (e inlineEncoder[T]) Decode(state *T, data []byte) (int, error) {
	return e.fnDecode(state, data)
}

func (e inlineEncoder[T]) Encode(state T) ([]byte, error) {
	return e.fnEncode(state)
}
