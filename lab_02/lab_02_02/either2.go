package main

type Either2[A, B any] interface {
	isEither2()
}

type Left2[T any] struct {
	value T
}

func (Left2[T]) isEither2() {}

type Right2[T any] struct {
	value T
}

func (Right2[T]) isEither2() {}

type FakeType[T any] struct {
	value T
}

func (FakeType[T]) isEither2() {}

func CaseFunction2[A, B, R any](obj Either2[A, B], leftFunc func(o A) R, rightFunc func(b B) R) R {
	switch v := obj.(type) {
	case Left2[A]:
		return leftFunc(v.value)
	case Right2[B]:
		return rightFunc(v.value)
	default:
		panic("unknown type")
	}
}
