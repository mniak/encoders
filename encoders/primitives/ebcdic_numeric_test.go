package primitives

import (
	"encoding/hex"
	"testing"

	"github.com/mniak/encoders/encoders/core"
	"github.com/stretchr/testify/assert"
)

func TestEBCDICNumeric_HappyScenario(t *testing.T) {
	testdata := []struct {
		name              string
		encoder           core.EncoderDecoder[int]
		input             string
		expected          int
		expectedRead      int
		expectedReencoded string
	}{
		{
			name:              "Empty",
			encoder:           EBCDICNumber(),
			input:             "",
			expected:          0,
			expectedRead:      0,
			expectedReencoded: "F0",
		},
		{
			name:              "1234",
			encoder:           EBCDICNumber(),
			input:             "F1F2F3F4",
			expected:          1234,
			expectedRead:      4,
			expectedReencoded: "F1F2F3F4",
		},
		{
			name:              "1234 (maxBytes 3)",
			encoder:           EBCDICNumber(),
			input:             "F1F2F3F4",
			expected:          123,
			expectedRead:      3,
			expectedReencoded: "F1F2F3",
		},
	}
	for _, td := range testdata {
		t.Run(td.name, func(t *testing.T) {
			inputBytes, err := hex.DecodeString(td.input)

			var result int
			read, err := td.encoder.Decode(&result, inputBytes)
			assert.NoError(t, err)
			assert.Equal(t, td.expectedRead, read)
			assert.Equal(t, td.expected, result)

			reencoded, err := td.encoder.Encode(result)
			assert.NoError(t, err)

			expectedReencodedBytes, err := hex.DecodeString(td.expectedReencoded)
			assert.Equal(t, expectedReencodedBytes, reencoded)
		})
	}
}
