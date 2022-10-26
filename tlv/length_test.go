package tlv

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLengthEncoder_EncodeThenDecode(t *testing.T) {
	enc := LengthEncoder()

	for l := Length(0); l < 300; l++ {
		t.Run(fmt.Sprint(l), func(t *testing.T) {
			encoded, err := enc.Encode(l)
			require.NoError(t, err)
			assert.NotEmpty(t, encoded)

			var decoded Length
			read, err := enc.Decode(&decoded, encoded)
			require.NoError(t, err)
			assert.Equal(t, l, decoded)
			assert.Equal(t, len(encoded), read)
		})
	}
}

func TestLengthEncoder_EncodeThenDecode_ExamplesInVisaDocumentation(t *testing.T) {
	enc := LengthEncoder()

	testdata := []struct {
		number uint
		bytes  []byte
	}{
		{
			number: 126,
			bytes:  []byte{0x7E},
		},
		{
			number: 254,
			bytes:  []byte{0x81, 0xFE},
		},
		{
			number: 382,
			bytes:  []byte{0x82, 0x01, 0x7E},
		},
		{
			number: 510,
			bytes:  []byte{0x82, 0x01, 0xFE},
		},
	}
	for _, td := range testdata {
		t.Run(fmt.Sprintf("%d -> 0x%02X", td.number, td.bytes), func(t *testing.T) {
			encoded, err := enc.Encode(Length(td.number))
			require.NoError(t, err)
			assert.NotEmpty(t, encoded)

			var decoded Length
			read, err := enc.Decode(&decoded, encoded)
			require.NoError(t, err)
			assert.Equal(t, td.number, uint(decoded))
			assert.Equal(t, len(encoded), read)
		})
	}
}
