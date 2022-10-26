package tlv

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTag_String(t *testing.T) {
	testdata := []struct {
		tag      int
		expected string
	}{
		{
			tag:      0x00,
			expected: "00",
		},
		{
			tag:      0xFF,
			expected: "FF",
		},
		{
			tag:      0x789,
			expected: "0789",
		},
		{
			tag:      0xabc,
			expected: "0ABC",
		},
	}
	for _, td := range testdata {
		t.Run(td.expected, func(t *testing.T) {
			tlvtag := Tag(td.tag)
			str := tlvtag.String()

			assert.Equal(t, td.expected, str)
		})
	}
}

func TestTagEncoder_ExamplesInVisaDocumentation(t *testing.T) {
	enc := TagEncoder()

	testdata := []struct {
		number uint
		bytes  []byte
	}{
		{
			number: 0x01,
			bytes:  []byte{0x01},
		},
		{
			number: 0xC0,
			bytes:  []byte{0xC0},
		},
		{
			number: 0xDF01,
			bytes:  []byte{0xDF, 0x01},
		},
	}
	for _, td := range testdata {
		t.Run(fmt.Sprintf("0x%02X", td.number), func(t *testing.T) {
			t.Run("Encode then Decode", func(t *testing.T) {
				encoded, err := enc.Encode(Tag(td.number))
				require.NoError(t, err)
				assert.NotEmpty(t, encoded)

				var decoded Tag
				read, err := enc.Decode(&decoded, append(encoded, []byte{0x99, 0x88, 0x77}...))
				require.NoError(t, err)
				assert.Equal(t, td.number, uint(decoded))
				assert.Equal(t, len(encoded), read)
			})

			t.Run("Decode then Encode", func(t *testing.T) {
				var decoded Tag
				read, err := enc.Decode(&decoded, append(td.bytes, []byte{0x99, 0x88, 0x77}...))
				require.NoError(t, err)
				assert.Equal(t, td.number, uint(decoded))
				assert.Equal(t, len(td.bytes), read)

				encoded, err := enc.Encode(Tag(td.number))
				require.NoError(t, err)
				assert.NotEmpty(t, encoded)
				assert.Equal(t, td.bytes, encoded)
			})
		})
	}
}
