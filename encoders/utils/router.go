package utils

import (
	"fmt"

	"github.com/mniak/encoders/encoders/core"
	"github.com/pkg/errors"
	"go.uber.org/multierr"
)

type mapRouter[K comparable, T any] struct {
	mapEncoder       core.EncoderDecoder[map[K][]byte]
	subEncoders      map[K]core.EncoderDecoder[T]
	itemDenomination string
}

func MapRouter[K comparable, T any](itemDenomination string, mapEncoder core.EncoderDecoder[map[K][]byte], subEncoders map[K]core.EncoderDecoder[T]) core.EncoderDecoder[T] {
	return mapRouter[K, T]{
		itemDenomination: itemDenomination,
		mapEncoder:       mapEncoder,
		subEncoders:      subEncoders,
	}
}

func (e mapRouter[K, T]) Decode(state *T, data []byte) (int, error) {
	var dataMap map[K][]byte
	read, err := e.mapEncoder.Decode(&dataMap, data)
	if err != nil {
		return read, err
	}

	var combinedError error
	for key, subdata := range dataMap {
		if enc, found := e.subEncoders[key]; found {
			subread, err := enc.Decode(state, subdata)
			if subread < len(subdata) {
				multierr.AppendInto(&err, fmt.Errorf("%d bytes were available but only %d were read", len(subdata), subread))
			}
			multierr.AppendInto(&combinedError, errors.WithMessagef(err, "%s %v", e.itemDenomination, key))
		}
	}
	return read, combinedError
}

func (e mapRouter[K, T]) Encode(state T) ([]byte, error) {
	dataMap := make(map[K][]byte)
	var combinedError error
	for key, enc := range e.subEncoders {
		subdata, err := enc.Encode(state)
		err = errors.WithMessagef(err, "%s %v", e.itemDenomination, key)
		if err != nil {
			multierr.AppendInto(&combinedError, err)
			continue
		}
		dataMap[key] = subdata
	}

	data, err := e.mapEncoder.Encode(dataMap)
	multierr.AppendInto(&combinedError, err)

	return data, combinedError
}
