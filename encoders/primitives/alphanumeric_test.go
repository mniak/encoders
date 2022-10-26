package primitives

import (
	"encoding/hex"
	"testing"

	"github.com/mniak/encoders/encoders/core"
	"github.com/stretchr/testify/assert"
)

func TestAlphaNumericAnsSymbols_HappyScenario(t *testing.T) {
	testdata := []struct {
		name              string
		encoder           core.EncoderDecoder[string]
		input             string
		expected          string
		expectedReencoded string
	}{
		{
			name:              "ANS - Empty",
			encoder:           AlphaNumericAndSymbols(),
			input:             "",
			expected:          "",
			expectedReencoded: "",
		},
		{
			name:              "ANS - Sample IP",
			encoder:           AlphaNumericAndSymbols(),
			input:             "F2F5F54BF2F5F54BF2F5F54BF2F5F5",
			expected:          "255.255.255.255",
			expectedReencoded: "F2F5F54BF2F5F54BF2F5F54BF2F5F5",
		},
		{
			name:              "AN - Empty",
			encoder:           AlphaNumeric(),
			input:             "",
			expected:          "",
			expectedReencoded: "",
		},
		{
			name:              "AN - Sample IP",
			encoder:           AlphaNumeric(),
			input:             "F2F5F54BF2F5F54BF2F5F54BF2F5F5",
			expected:          "255.255.255.255",
			expectedReencoded: "F2F5F54BF2F5F54BF2F5F54BF2F5F5",
		},
	}
	for _, td := range testdata {
		t.Run(td.name, func(t *testing.T) {
			inputBytes, err := hex.DecodeString(td.input)

			var result string
			read, err := td.encoder.Decode(&result, inputBytes)
			assert.NoError(t, err)
			assert.Equal(t, len(td.expected), read)
			assert.Equal(t, td.expected, result)

			reencoded, err := td.encoder.Encode(result)
			assert.NoError(t, err)

			expectedReencodedBytes, err := hex.DecodeString(td.expectedReencoded)
			assert.Equal(t, expectedReencodedBytes, reencoded)
		})
	}
}
