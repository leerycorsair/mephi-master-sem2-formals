package main

import (
	"pgregory.net/rapid"
)

func PropProd[Src, R1, R2 any, P Pair[R1, R2]](
	f1 func(Src) R1,
	f2 func(Src) R2,
	cmp1 func(R1, R1) bool,
	cmp2 func(R2, R2) bool,
	generator *rapid.Generator[Src]) func(*rapid.T) {
	return func(t *rapid.T) {
		src := generator.Draw(t, "Src")
		resultF1 := f1(src)
		resultF2 := f2(src)
		resultF12 := BuildPair(src, f1, f2)
		if !cmp1(resultF1, First(resultF12)) ||
			!cmp2(resultF2, Second(resultF12)) {
			t.Fatalf("Error")
		}
	}
}
