package functors

import "fmt"

type Maybe[T any] struct {
	value *T
	isNil bool
}

func Just[T any](value T) Maybe[T] {
	return Maybe[T]{value: &value, isNil: false}
}

func Nothing[T any]() Maybe[T] {
	return Maybe[T]{value: nil, isNil: true}
}

func MapMaybe[T1, T2 any](m Maybe[T1], f func(T1) T2) Maybe[T2] {
	if m.isNil {
		return Nothing[T2]()
	}
	return Just(f(*m.value))
}

func MaybeCmp[T any](m1 Maybe[T], m2 Maybe[T], cmp func(a, b T) bool) bool {
	if m1.isNil && m2.isNil {
		return true
	} else if m1.isNil && !m2.isNil || !m1.isNil && m2.isNil {
		return false
	} else {
		return cmp(*m1.value, *m2.value)
	}
}

func (m Maybe[T]) String() string {
	if m.isNil {
		return "Maybe Obj = nil"
	}
	return fmt.Sprintf("Maybe Obj = %v", *m.value)
}
