package utils

import (
	"errors"
	"fmt"

	"github.com/mniak/encoders/core"
)

type rawMapEncoder[K comparable] struct {
	keyEncoder   core.EncoderDecoder[K]
	valueEncoder core.EncoderDecoder[[]byte]
}

func RawMap[K comparable](
	keyEncoder core.EncoderDecoder[K],
	valueEncoder core.EncoderDecoder[[]byte],
) core.EncoderDecoder[map[K][]byte] {
	return rawMapEncoder[K]{
		keyEncoder:   keyEncoder,
		valueEncoder: valueEncoder,
	}
}

func (e rawMapEncoder[K]) Decode(state *map[K][]byte, data []byte) (int, error) {
	if state == nil || *state == nil {
		*state = make(map[K][]byte)
	}
	nonPointerState := *state
	var totalRead int
	for len(data) > 0 {
		var key K
		read, err := e.keyEncoder.Decode(&key, data)
		data = data[read:]
		totalRead += read

		if err != nil {
			return totalRead, err
		} else if read == 0 {
			return totalRead, fmt.Errorf("%d bytes available but the key decoder didn't consume any", len(data))
		}

		var value []byte
		read, err = e.valueEncoder.Decode(&value, data)
		data = data[read:]
		totalRead += read

		if err != nil {
			return totalRead, err
		}
		nonPointerState[key] = value
	}
	*state = nonPointerState
	return totalRead, nil
}

func (e rawMapEncoder[K]) Encode(state map[K][]byte) ([]byte, error) {
	return nil, errors.New("TODO: not implemented (simple map)")
}
