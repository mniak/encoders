package tlv

// import (
// 	"fmt"
// 	"testing"

// 	"github.com/brianvoe/gofakeit/v6"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/require"
// )

// func TestDecode_ZeroLength(t *testing.T) {
// 	var tlvmap Map
// 	err := Decode(&tlvmap, []byte{})
// 	require.NoError(t, err)

// 	assert.Empty(t, tlvmap)
// }

// func TestDecode_EmptyPayload(t *testing.T) {
// 	t.Run("Fixed examples", func(t *testing.T) {
// 		testdata := []struct {
// 			suite        string
// 			expectedTag  int
// 			input        []byte
// 			expectedData []byte
// 		}{
// 			{
// 				suite:        "Single-byte tag",
// 				expectedTag:  0xC0,
// 				input:        []byte{0xC0, 0x00},
// 				expectedData: []byte{},
// 			},
// 			{
// 				suite:        "Single-byte tag",
// 				expectedTag:  0xCF,
// 				input:        []byte{0xCF, 0x00},
// 				expectedData: []byte{},
// 			},
// 			{
// 				suite:        "Single-byte tag",
// 				expectedTag:  0xD0,
// 				input:        []byte{0xD0, 0x00},
// 				expectedData: []byte{},
// 			},
// 			{
// 				suite:        "Single-byte tag",
// 				expectedTag:  0xD4,
// 				input:        []byte{0xD4, 0x00},
// 				expectedData: []byte{},
// 			},

// 			{
// 				suite:        "Double-byte tag",
// 				expectedTag:  0x1F01,
// 				input:        []byte{0x1F, 0x01, 0x00},
// 				expectedData: []byte{},
// 			},
// 			{
// 				suite:        "Double-byte tag",
// 				expectedTag:  0x3F12,
// 				input:        []byte{0x3F, 0x12, 0x00},
// 				expectedData: []byte{},
// 			},
// 			{
// 				suite:        "Double-byte tag",
// 				expectedTag:  0x5F24,
// 				input:        []byte{0x5F, 0x24, 0x00},
// 				expectedData: []byte{},
// 			},
// 			{
// 				suite:        "Double-byte tag",
// 				expectedTag:  0x7F38,
// 				input:        []byte{0x7F, 0x38, 0x00},
// 				expectedData: []byte{},
// 			},
// 			{
// 				suite:        "Double-byte tag",
// 				expectedTag:  0x9F4F,
// 				input:        []byte{0x9F, 0x4F, 0x00},
// 				expectedData: []byte{},
// 			},
// 			{
// 				suite:        "Double-byte tag",
// 				expectedTag:  0xBF53,
// 				input:        []byte{0xBF, 0x53, 0x00},
// 				expectedData: []byte{},
// 			},
// 			{
// 				suite:        "Double-byte tag",
// 				expectedTag:  0xDF65,
// 				input:        []byte{0xDF, 0x65, 0x00},
// 				expectedData: []byte{},
// 			},
// 			{
// 				suite:        "Double-byte tag",
// 				expectedTag:  0xFF78,
// 				input:        []byte{0xFF, 0x78, 0x00},
// 				expectedData: []byte{},
// 			},
// 			{
// 				suite:        "Double-byte tag",
// 				expectedTag:  0x1F7D,
// 				input:        []byte{0x1F, 0x7D, 0x00},
// 				expectedData: []byte{},
// 			},
// 		}
// 		for _, td := range testdata {
// 			t.Run(fmt.Sprintf("%s: %X", td.suite, td.input), func(t *testing.T) {
// 				var tlvmap Map
// 				err := Decode(&tlvmap, td.input)
// 				require.NoError(t, err)

// 				expected := Map{
// 					Tag(td.expectedTag): td.expectedData,
// 				}
// 				assert.Equal(t, expected, tlvmap)
// 			})
// 		}

// 		t.Run("All tags together", func(t *testing.T) {
// 			tlvraw := []byte{}
// 			expected := Map{}

// 			for _, td := range testdata {
// 				tlvraw = append(tlvraw, td.input...)
// 				expected[Tag(td.expectedTag)] = td.expectedData
// 			}

// 			var tlvmap Map
// 			err := Decode(&tlvmap, tlvraw)
// 			require.NoError(t, err)

// 			assert.Equal(t, expected, tlvmap)
// 		})
// 	})

// 	t.Run("Single-byte random tags", func(t *testing.T) {
// 		// Force those tags to be single byte (not having the 5 right-most bytes 1)
// 		tag1 := gofakeit.Uint8() & 0b1110_1111
// 		tag2 := gofakeit.Uint8() & 0b1110_1111

// 		var tlvmap Map
// 		err := Decode(&tlvmap, []byte{
// 			tag1, 0x00,
// 			tag2, 0x00,
// 		})
// 		require.NoError(t, err)

// 		expected := Map{
// 			Tag(tag1): []byte{},
// 			Tag(tag2): []byte{},
// 		}
// 		assert.Equal(t, expected, tlvmap)
// 	})
// }

// func TestDecode_HappyScenario(t *testing.T) {
// 	testdata := []struct {
// 		suite  string
// 		length []byte
// 		data   []byte
// 	}{
// 		{
// 			suite:  "Examples in documentation",
// 			length: []byte{0x7E},
// 			data:   []byte(gofakeit.LetterN(126)),
// 		},
// 		{
// 			suite:  "Examples in documentation",
// 			length: []byte{0x81, 0xFE},
// 			data:   []byte(gofakeit.LetterN(254)),
// 		},
// 		{
// 			suite:  "Examples in documentation",
// 			length: []byte{0x82, 0x01, 0x7E},
// 			data:   []byte(gofakeit.LetterN(382)),
// 		},
// 		{
// 			suite:  "Examples in documentation",
// 			length: []byte{0x82, 0x01, 0xFE},
// 			data:   []byte(gofakeit.LetterN(510)),
// 		},

// 		{
// 			suite:  "Short Form",
// 			length: []byte{1},
// 			data:   []byte(gofakeit.LetterN(1)),
// 		},
// 		{
// 			suite:  "Short Form",
// 			length: []byte{127},
// 			data:   []byte(gofakeit.LetterN(127)),
// 		},

// 		{
// 			suite:  "Long Form (2 bytes)",
// 			length: []byte{0x81, 1},
// 			data:   []byte(gofakeit.LetterN(1)),
// 		},
// 		{
// 			suite:  "Long Form (2 bytes)",
// 			length: []byte{0x81, 127},
// 			data:   []byte(gofakeit.LetterN(127)),
// 		},
// 		{
// 			suite:  "Long Form (2 bytes)",
// 			length: []byte{0x81, 128},
// 			data:   []byte(gofakeit.LetterN(128)),
// 		},
// 		{
// 			suite:  "Long Form (2 bytes)",
// 			length: []byte{0x81, 0xFF},
// 			data:   []byte(gofakeit.LetterN(255)),
// 		},

// 		{
// 			suite:  "Long Form (3 bytes)",
// 			length: []byte{0x82, 0x00, 1},
// 			data:   []byte(gofakeit.LetterN(1)),
// 		},
// 		{
// 			suite:  "Long Form (3 bytes)",
// 			length: []byte{0x82, 0x00, 127},
// 			data:   []byte(gofakeit.LetterN(127)),
// 		},
// 		{
// 			suite:  "Long Form (3 bytes)",
// 			length: []byte{0x82, 0x00, 128},
// 			data:   []byte(gofakeit.LetterN(128)),
// 		},
// 		{
// 			suite:  "Long Form (3 bytes)",
// 			length: []byte{0x82, 0x00, 0xFF},
// 			data:   []byte(gofakeit.LetterN(255)),
// 		},
// 		{
// 			suite:  "Long Form (3 bytes)",
// 			length: []byte{0x82, 0x01, 0x00},
// 			data:   []byte(gofakeit.LetterN(256)),
// 		},
// 		{
// 			suite:  "Long Form (3 bytes)",
// 			length: []byte{0x82, 0x01, 0x01},
// 			data:   []byte(gofakeit.LetterN(257)),
// 		},
// 		{
// 			suite:  "Long Form (3 bytes)",
// 			length: []byte{0x82, 0xFF, 0xFF},
// 			data:   []byte(gofakeit.LetterN(65535)),
// 		},
// 	}
// 	for tag, td := range testdata {
// 		t.Run(fmt.Sprintf("%s: %X", td.suite, td.length), func(t *testing.T) {
// 			tlvraw := []byte{}
// 			tlvraw = append(tlvraw, byte(tag))
// 			tlvraw = append(tlvraw, td.length...)
// 			tlvraw = append(tlvraw, td.data...)

// 			var tlvmap Map
// 			err := Decode(&tlvmap, tlvraw)
// 			require.NoError(t, err)

// 			expected := Map{
// 				Tag(tag): td.data,
// 			}
// 			assert.Equal(t, expected, tlvmap)
// 		})
// 	}

// 	t.Run("All tags together", func(t *testing.T) {
// 		tlvraw := []byte{}
// 		expected := Map{}

// 		for tag, td := range testdata {
// 			tlvraw = append(tlvraw, byte(tag))
// 			tlvraw = append(tlvraw, td.length...)
// 			tlvraw = append(tlvraw, td.data...)

// 			expected[Tag(tag)] = td.data
// 		}

// 		var tlvmap Map
// 		err := Decode(&tlvmap, tlvraw)
// 		require.NoError(t, err)

// 		assert.Equal(t, expected, tlvmap)
// 	})
// }

// func TestDecode_AllErrors(t *testing.T) {
// 	testdata := []struct {
// 		name      string
// 		dataFunc  func() []byte
// 		errorFlag error
// 		errorMsg  string
// 	}{
// 		//// Tag
// 		{
// 			name: "Tag expecting second byte",
// 			dataFunc: func() []byte {
// 				return []byte{0xFF}
// 			},
// 			errorFlag: ErrTagShouldHave2Bytes,
// 			errorMsg:  "could not decode TLV tag: expecting at least 2 bytes",
// 		},
// 		{
// 			name: "Tag expecting third byte",
// 			dataFunc: func() []byte {
// 				return []byte{0xFF, 0xFF}
// 			},
// 			errorFlag: ErrTagTooLong,
// 			errorMsg:  "could not decode TLV tag: second byte of tag is not final, but VisaNet supports at most 2 bytes",
// 		},

// 		//// Length
// 		{
// 			name: "Length missing should fail",
// 			dataFunc: func() []byte {
// 				return []byte{0x08}
// 			},
// 			errorFlag: ErrLengthHasNoBytes,
// 			errorMsg:  "could not decode TLV length for tag 08: no bytes found",
// 		},
// 		{
// 			name: "Length not having the expected 2nd byte",
// 			dataFunc: func() []byte {
// 				tlvraw := []byte{}
// 				tlvraw = append(tlvraw, 0x09)
// 				tlvraw = append(tlvraw, 0x81)
// 				return tlvraw
// 			},
// 			errorFlag: ErrLengthFormatError,
// 			errorMsg:  "could not decode TLV length for tag 09: expecting at least 2 bytes",
// 		},
// 		{
// 			name: "Length not having the expected 2nd and 3rd bytes",
// 			dataFunc: func() []byte {
// 				tlvraw := []byte{}
// 				tlvraw = append(tlvraw, 0x10)
// 				tlvraw = append(tlvraw, 0x82, 0x00)
// 				return tlvraw
// 			},
// 			errorFlag: ErrLengthFormatError,
// 			errorMsg:  "could not decode TLV length for tag 10: expecting at least 3 bytes",
// 		},
// 		{
// 			name: "Length having more than 3 bytes",
// 			dataFunc: func() []byte {
// 				tlvraw := []byte{}
// 				tlvraw = append(tlvraw, 0x20)
// 				tlvraw = append(tlvraw, 0x89, 0x00, 0x00)
// 				return tlvraw
// 			},
// 			errorFlag: ErrLengthFormatError,
// 			errorMsg:  "could not decode TLV length for tag 20: the length field should have at most 3 bytes but it appears to have 10",
// 		},

// 		//// Value
// 		{
// 			name: "Value",
// 			dataFunc: func() []byte {
// 				tlvraw := []byte{}
// 				tlvraw = append(tlvraw, 0x01)
// 				tlvraw = append(tlvraw, 0x04)
// 				tlvraw = append(tlvraw, 0x01, 0x02, 0x03)
// 				return tlvraw
// 			},
// 			errorFlag: ErrValueTooShort,
// 			errorMsg:  "could not read TLV tag 01: expecting 4 bytes but only found 3",
// 		},
// 	}
// 	for _, td := range testdata {
// 		t.Run(td.name, func(t *testing.T) {
// 			tlvraw := td.dataFunc()

// 			var tlvmap Map
// 			err := Decode(&tlvmap, tlvraw)

// 			assert.Empty(t, tlvmap)

// 			require.Error(t, err)
// 			assert.ErrorIs(t, err, td.errorFlag)
// 			assert.EqualError(t, err, td.errorMsg)
// 		})
// 	}
// }
