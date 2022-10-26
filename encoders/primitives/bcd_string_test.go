package primitives

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBCDString_Decode_Empty(t *testing.T) {
	enc := BCDString()
	data := []byte{}
	var valueRead string
	bytesRead, err := enc.Decode(&valueRead, data)
	assert.Zero(t, bytesRead)
	assert.NoError(t, err)
	assert.Empty(t, valueRead)
}

func TestBCDString_Decode_Examples(t *testing.T) {
	enc := BCDString()

	testdata := []struct {
		input         []byte
		expectedError error
		expected      string
		expectedValue string
		expectedRead  int
	}{
		{
			input:         []byte{},
			expectedError: nil,
			expectedValue: "",
			expected:      "",
			expectedRead:  0,
		},
		{
			input:         []byte{0x11, 0x22, 0x33, 0x44},
			expectedError: nil,
			expectedValue: "11223344",
			expected:      "11223344",
			expectedRead:  4,
		},
		{
			input:         []byte{0x11, 0x22, 0x33, 0x44, 0xFF},
			expectedError: ErrByteNotBCD,
			expectedValue: "",
			expected:      "",
			expectedRead:  0,
		},
	}
	for _, td := range testdata {
		t.Run(hex.EncodeToString(td.input), func(t *testing.T) {
			var valueRead string
			bytesRead, err := enc.Decode(&valueRead, td.input)
			assert.Equal(t, td.expectedRead, bytesRead)
			assert.ErrorIs(t, err, td.expectedError)
			assert.Equal(t, td.expected, valueRead)
		})
	}
}

func TestBCDString_Decode_Random(t *testing.T) {
	enc := BCDString()

	for _, length := range []int{1, 3, 8, 16, 257} {
		t.Run(fmt.Sprintf("Length %d", length), func(t *testing.T) {
			expectedValue := gofakeit.Numerify(strings.Repeat("#", length*2))

			data, err := hex.DecodeString(expectedValue)
			require.NoError(t, err)

			var valueRead string
			bytesRead, err := enc.Decode(&valueRead, data)
			assert.Equal(t, len(expectedValue)/2, bytesRead)
			assert.NoError(t, err)
			require.NotNil(t, valueRead)
			assert.Equal(t, expectedValue, valueRead)
		})
	}
}

func TestBCDString_Decode_ByteNotBCD(t *testing.T) {
	enc := BCDString()

	t.Run("Data bigger than expected length", func(t *testing.T) {
		for _, nonBCDByte := range []byte{0x1A, 0xB2, 0xC, 0xDE} {
			t.Run(fmt.Sprintf("byte %02x", nonBCDByte), func(t *testing.T) {
				length := 8
				for invalidByteIndex := 0; invalidByteIndex < length; invalidByteIndex++ {
					t.Run(fmt.Sprintf("at idx %d", invalidByteIndex), func(t *testing.T) {
						expectedValue := gofakeit.Numerify(strings.Repeat("#", length*2))

						data, err := hex.DecodeString(expectedValue)
						require.NoError(t, err)

						data[invalidByteIndex] = nonBCDByte

						expectedErrorMessage := fmt.Sprintf("BCD decode: byte at index %d is invalid", invalidByteIndex)
						var valueRead string
						bytesRead, err := enc.Decode(&valueRead, data)
						assert.Zero(t, bytesRead)
						assert.Zero(t, valueRead)
						assert.EqualError(t, err, expectedErrorMessage)
						assert.ErrorIs(t, err, ErrByteNotBCD)
					})
				}
			})
		}
	})
}

func TestBCDString_Encode_Examples(t *testing.T) {
	enc := BCDString()

	testdata := []struct {
		input         string
		expected      []byte
		expectedError error
	}{
		{
			input:    "",
			expected: []byte{},
		},
		{
			input:    "0",
			expected: []byte{0x00},
		},
		{
			input:    "0000",
			expected: []byte{0x00, 0x00},
		},
		{
			input:    "123",
			expected: []byte{0x01, 0x23},
		},
		{
			input:    "1234",
			expected: []byte{0x12, 0x34},
		},
		{
			input:    "90817263540123456789",
			expected: []byte{0x90, 0x81, 0x72, 0x63, 0x54, 0x01, 0x23, 0x45, 0x67, 0x89},
		},
		// Errors
		{
			input:         "A",
			expected:      nil,
			expectedError: ErrNotADigit,
		},
		{
			input:         "0123456789o",
			expected:      nil,
			expectedError: ErrNotADigit,
		},
	}
	for _, td := range testdata {
		t.Run(td.input, func(t *testing.T) {
			result, err := enc.Encode(td.input)

			assert.ErrorIs(t, err, td.expectedError)
			assert.Equal(t, td.expected, result)
		})
	}
}

func TestBCDString_EncodeAndDecode(t *testing.T) {
	enc := BCDString()
	t.Run("Random Numbers", func(t *testing.T) {
		input := int(gofakeit.Uint32())
		inputAsString := strconv.Itoa(input)
		if len(inputAsString)%2 == 1 {
			inputAsString = "0" + inputAsString
		}

		encoded, err := enc.Encode(inputAsString)
		require.NoError(t, err)

		var decodedString string
		read, err := enc.Decode(&decodedString, encoded)
		require.NoError(t, err)
		assert.Equal(t, len(encoded), read)

		assert.Equal(t, inputAsString, decodedString)
		decoded, err := strconv.Atoi(decodedString)
		require.NoError(t, err)
		assert.Equal(t, input, decoded)
	})
	t.Run("Numbers in range 9990-10010", func(t *testing.T) {
		for input := 9990; input <= 10010; input++ {
			t.Run(strconv.Itoa(input), func(t *testing.T) {
				inputAsString := strconv.Itoa(input)
				if len(inputAsString)%2 == 1 {
					inputAsString = "0" + inputAsString
				}

				encoded, err := enc.Encode(inputAsString)
				require.NoError(t, err)

				var decodedString string
				read, err := enc.Decode(&decodedString, encoded)
				require.NoError(t, err)
				assert.Equal(t, len(encoded), read)

				assert.Equal(t, inputAsString, decodedString)
				decoded, err := strconv.Atoi(decodedString)
				require.NoError(t, err)
				assert.Equal(t, input, decoded)
			})
		}
	})
}
