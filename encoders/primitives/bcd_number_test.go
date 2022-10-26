package primitives

import (
	"encoding/hex"
	"fmt"
	"math"
	"strconv"
	"strings"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBCDNumber_Decode_Empty(t *testing.T) {
	enc := BCDNumber()
	data := []byte{}
	var valueRead int
	bytesRead, err := enc.Decode(&valueRead, data)
	assert.Zero(t, bytesRead)
	assert.NoError(t, err)
	assert.Zero(t, valueRead)
}

func TestBCDNumber_Decode_Examples(t *testing.T) {
	enc := BCDNumber()

	testdata := []struct {
		input         []byte
		expectedError error
		expected      int
		expectedRead  int
	}{
		{
			input:        []byte{},
			expected:     0,
			expectedRead: 0,
		},
		{
			input:        []byte{0x01, 0x02, 0x03},
			expected:     10203,
			expectedRead: 3,
		},
		{
			input:        []byte{0x11, 0x22, 0x33, 0x44},
			expected:     11223344,
			expectedRead: 4,
		},
		{
			input:        []byte{0x09, 0x22, 0x33, 0x72, 0x03, 0x68, 0x54, 0x77, 0x58, 0x07},
			expected:     math.MaxInt,
			expectedRead: 10,
		},
		{
			input:         []byte{0x11, 0x22, 0x33, 0x44, 0xFF},
			expectedError: ErrByteNotBCD,
			expected:      0,
			expectedRead:  0,
		},
	}
	for _, td := range testdata {
		t.Run(hex.EncodeToString(td.input), func(t *testing.T) {
			var valueRead int
			bytesRead, err := enc.Decode(&valueRead, td.input)
			assert.Equal(t, td.expectedRead, bytesRead)
			assert.ErrorIs(t, err, td.expectedError)
			assert.Equal(t, td.expected, valueRead)
		})
	}
}

func TestBCDNumber_Decode_Random(t *testing.T) {
	enc := BCDNumber()

	for _, length := range []int{1, 3, 8} {
		t.Run(fmt.Sprintf("%d bytes", length), func(t *testing.T) {
			randomBCDString := gofakeit.Numerify(strings.Repeat("#", length*2))
			expectedValue, err := strconv.Atoi(randomBCDString)
			require.NoError(t, err)

			data, err := hex.DecodeString(randomBCDString)
			require.NoError(t, err)

			var valueRead int
			bytesRead, err := enc.Decode(&valueRead, data)
			assert.Equal(t, len(randomBCDString)/2, bytesRead)
			assert.NoError(t, err)
			require.NotNil(t, valueRead)
			assert.Equal(t, expectedValue, valueRead)
		})
	}
}

func TestBCDNumber_Decode_ByteNotBCD(t *testing.T) {
	enc := BCDNumber()

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
						var valueRead int
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

func TestBCDNumber_Encode_Examples(t *testing.T) {
	enc := BCDNumber()

	testdata := []struct {
		input         int
		expected      []byte
		expectedError error
	}{
		{
			input:    0,
			expected: []byte{0x00},
		},
		{
			input:    123,
			expected: []byte{0x01, 0x23},
		},
		{
			input:    1234,
			expected: []byte{0x12, 0x34},
		},
		{
			input:    math.MaxInt,
			expected: []byte{0x09, 0x22, 0x33, 0x72, 0x03, 0x68, 0x54, 0x77, 0x58, 0x07},
		},
	}
	for _, td := range testdata {
		t.Run(fmt.Sprint(td.input), func(t *testing.T) {
			result, err := enc.Encode(td.input)

			assert.ErrorIs(t, err, td.expectedError)
			assert.Equal(t, td.expected, result)
		})
	}
}

func TestBCDNumber_EncodeAndDecode(t *testing.T) {
	enc := BCDNumber()
	t.Run("Random Numbers", func(t *testing.T) {
		input := int(gofakeit.Uint32())

		encoded, err := enc.Encode(input)
		require.NoError(t, err)

		var decoded int
		read, err := enc.Decode(&decoded, encoded)
		require.NoError(t, err)
		assert.Equal(t, len(encoded), read)
		assert.Equal(t, input, decoded)
	})
	t.Run("Numbers in range 9990-10010", func(t *testing.T) {
		for input := 9990; input <= 10010; input++ {
			t.Run(strconv.Itoa(input), func(t *testing.T) {
				encoded, err := enc.Encode(input)
				require.NoError(t, err)

				var decoded int
				read, err := enc.Decode(&decoded, encoded)
				require.NoError(t, err)
				assert.Equal(t, len(encoded), read)
				assert.Equal(t, input, decoded)
			})
		}
	})
}
