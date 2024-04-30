package main

type Either[A, B any] struct {
	isLeft bool
	left   A
	right  B
}

func Left[A, B any](value A) Either[A, B] {
	return Either[A, B]{
		isLeft: true,
		left:   value,
	}
}

func Right[A, B any](value B) Either[A, B] {
	return Either[A, B]{
		isLeft: false,
		right:  value,
	}
}

func IsLeft[A, B any](e Either[A, B]) bool {
	return e.isLeft
}

func IsRight[A, B any](e Either[A, B]) bool {
	return !e.isLeft
}

func CaseFunction[A, B, R any](e Either[A, B], leftFunc func(a A) R, rightFunc func(b B) R) R {
	if IsLeft(e) {
		return leftFunc(e.left)
	} else {
		return rightFunc(e.right)
	}
}
