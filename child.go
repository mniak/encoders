package encoders

import "github.com/mniak/encoders/core"

type childEncoder[Parent any, Child any] struct {
	inner    core.EncoderDecoder[Child]
	getchild func(x *Parent) *Child
}

func ForChild[Parent any, Child any](getchild func(x *Parent) *Child, inner core.EncoderDecoder[Child]) core.EncoderDecoder[Parent] {
	return childEncoder[Parent, Child]{
		inner:    inner,
		getchild: getchild,
	}
}

func ForChildPointer[Parent any, Child any](getchild func(x *Parent) **Child, inner core.EncoderDecoder[Child]) core.EncoderDecoder[Parent] {
	return childEncoder[Parent, Child]{
		inner: inner,
		getchild: func(x *Parent) *Child {
			pointer := getchild(x)
			if (*pointer) == nil {
				var empty Child
				(*pointer) = &empty
			}
			return *pointer
		},
	}
}

func (e childEncoder[Parent, Child]) Decode(state *Parent, data []byte) (int, error) {
	if state == nil {
		var zerovalue Parent
		state = &zerovalue
	}
	child := e.getchild(state)
	return e.inner.Decode(child, data)
}

func (e childEncoder[Parent, Child]) Encode(state Parent) ([]byte, error) {
	child := e.getchild(&state)
	return e.inner.Encode(*child)
}
