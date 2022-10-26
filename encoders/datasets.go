package encoders

import (
	"fmt"

	"github.com/mniak/encoders/encoders/core"
	"github.com/mniak/encoders/encoders/primitives"
	"github.com/mniak/encoders/encoders/utils"
)

type DatasetKey uint

func (dk DatasetKey) String() string {
	return fmt.Sprintf("0x%02X", int(dk))
}

func RawDatasetList(datasetLengthSize byte) core.EncoderDecoder[map[DatasetKey][]byte] {
	return utils.RawMap[DatasetKey](
		AsInt[DatasetKey](primitives.FixedSizeNumber(1)),
		PrefixedLengthEncoder(datasetLengthSize, primitives.Bypass()),
	)
}

func DatasetList[T any](datasetLengthSize byte, subEncoders map[DatasetKey]core.EncoderDecoder[T]) core.EncoderDecoder[T] {
	return utils.MapRouter("Dataset", RawDatasetList(datasetLengthSize), subEncoders)
}
