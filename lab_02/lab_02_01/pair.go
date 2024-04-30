package main

type Pair[A, B any] struct {
	value1 A
	value2 B
}

func BuildPair[SrcT, A, B any](
	s SrcT, f1 func(SrcT) A, f2 func(SrcT) B) Pair[A, B] {
	return Pair[A, B]{
		value1: f1(s),
		value2: f2(s),
	}
}

func First[A, B any](p Pair[A, B]) A {
	return p.value1
}

func Second[A, B any](p Pair[A, B]) B {
	return p.value2
}
