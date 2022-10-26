package encoders

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/mniak/encoders/encoders/internal/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_NewChildPointerEncoder(t *testing.T) {
	type A struct {
		pointerValue *string
	}
	var state A

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fakeRead := int(gofakeit.Int32())
	fakeData := []byte(gofakeit.SentenceSimple())
	fakeStateChanged := gofakeit.BuzzWord()

	mockEncoder := mocks.NewMockEncoderDecoderOfString(ctrl)
	mockEncoder.EXPECT().
		Decode(gomock.Any(), fakeData).
		Do(func(state *string, data []byte) {
			*state = fakeStateChanged
		}).
		Return(fakeRead, nil)

	encAtoStringPointer := ForChildPointer[A, string](func(x *A) **string { return &x.pointerValue }, mockEncoder)

	read, err := encAtoStringPointer.Decode(&state, fakeData)

	require.NoError(t, err)
	assert.Equal(t, fakeRead, read)
	assert.NotNil(t, state.pointerValue)
	assert.Equal(t, fakeStateChanged, *state.pointerValue)
}

func Test_NewChildEncoder(t *testing.T) {
	type A struct {
		value string
	}
	var state A

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fakeRead := int(gofakeit.Int32())
	fakeData := []byte(gofakeit.SentenceSimple())
	fakeStateChanged := gofakeit.BuzzWord()

	mockEncoder := mocks.NewMockEncoderDecoderOfString(ctrl)
	mockEncoder.EXPECT().
		Decode(gomock.Any(), fakeData).
		Do(func(state *string, data []byte) {
			*state = fakeStateChanged
		}).
		Return(fakeRead, nil)

	encAtoStringPointer := ForChild[A, string](func(x *A) *string { return &x.value }, mockEncoder)

	read, err := encAtoStringPointer.Decode(&state, fakeData)

	require.NoError(t, err)
	assert.Equal(t, fakeRead, read)
	assert.NotNil(t, state.value)
	assert.Equal(t, fakeStateChanged, state.value)
}
