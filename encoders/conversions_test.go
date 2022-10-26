package encoders

import (
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/mniak/encoders/encoders/internal/mocks"
	"github.com/stretchr/testify/assert"
)

func Test_Int_AsInt_Uint32Int_Decode(t *testing.T) {
	t.Run("Happy Scenario", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		fakeNumber := int(gofakeit.Int16())
		fakeData := []byte(gofakeit.Word())
		fakeRead := int(gofakeit.Int16())

		mock := mocks.NewMockEncoderDecoderOfInt(ctrl)
		mock.EXPECT().
			Decode(gomock.Any(), fakeData).
			Do(func(state *int, data []byte) {
				*state = fakeNumber
			}).
			Return(fakeRead, nil)

		enc := AsInt[uint32, int](mock)

		var result uint32
		read, err := enc.Decode(&result, fakeData)
		assert.Equal(t, fakeRead, read)
		assert.NoError(t, err)
		assert.Equal(t, uint32(fakeNumber), result)
	})

	t.Run("When error should still convert", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		fakeNumber := int(gofakeit.Int16())
		fakeData := []byte(gofakeit.Word())
		fakeRead := int(gofakeit.Int16())
		fakeError := errors.New(gofakeit.Sentence(4))

		mock := mocks.NewMockEncoderDecoderOfInt(ctrl)
		mock.EXPECT().
			Decode(gomock.Any(), fakeData).
			Do(func(state *int, data []byte) {
				*state = fakeNumber
			}).
			Return(fakeRead, fakeError)

		enc := AsInt[uint32, int](mock)

		var result uint32
		read, err := enc.Decode(&result, fakeData)
		assert.Equal(t, fakeRead, read)
		assert.Equal(t, fakeError, err)
		assert.Equal(t, uint32(fakeNumber), result)
	})
}

func Test_Int_AsInt_Uint32Int_Encode(t *testing.T) {
	t.Run("Happy Scenario", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		fakeNumber := uint32(gofakeit.Int16())
		fakeBytes := []byte(gofakeit.Word())

		mock := mocks.NewMockEncoderDecoderOfInt(ctrl)
		mock.EXPECT().
			Encode(int(fakeNumber)).
			Return(fakeBytes, nil)

		enc := AsInt[uint32, int](mock)

		resultBytes, err := enc.Encode(fakeNumber)
		assert.NoError(t, err)
		assert.Equal(t, fakeBytes, resultBytes)
	})
	t.Run("When error should still convert", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		fakeNumber := uint32(gofakeit.Int16())
		fakeBytes := []byte(gofakeit.Word())
		fakeError := errors.New(gofakeit.Sentence(4))

		mock := mocks.NewMockEncoderDecoderOfInt(ctrl)
		mock.EXPECT().
			Encode(int(fakeNumber)).
			Return(fakeBytes, fakeError)

		enc := AsInt[uint32, int](mock)

		resultBytes, err := enc.Encode(fakeNumber)
		assert.Equal(t, fakeError, err)
		assert.Equal(t, fakeBytes, resultBytes)
	})
}
