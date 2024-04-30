package main

import (
	"pgregory.net/rapid"
)

func PropSum[T1 any, T2 any, EType Either[T1, T2], ResT any](
	f1 func(T1) ResT,
	f2 func(T2) ResT,
	gen1 *rapid.Generator[T1],
	gen2 *rapid.Generator[T2],
	resCmp func(ResT, ResT) bool) func(*rapid.T) {
	return func(t *rapid.T) {
		val1 := gen1.Draw(t, "val1")
		val2 := gen2.Draw(t, "val2")

		if !resCmp(CaseFunction[T1, T2](Left[T1, T2](val1), f1, f2), f1(val1)) ||
			!resCmp(CaseFunction[T1, T2](Right[T1, T2](val2), f1, f2), f2(val2)) {
			t.Fatalf("Error")
		}
	}
}

func PropSum2[T1 any, T2 any, ResT any](
	f1 func(T1) ResT,
	f2 func(T2) ResT,
	gen1 *rapid.Generator[T1],
	gen2 *rapid.Generator[T2],
	resCmp func(ResT, ResT) bool) func(*rapid.T) {
	return func(t *rapid.T) {
		val1 := gen1.Draw(t, "val1")
		val2 := gen2.Draw(t, "val2")

		if !resCmp(CaseFunction2[T1, T2, ResT](Left2[T1]{val1}, f1, f2), f1(val1)) ||
			!resCmp(CaseFunction2[T1, T2, ResT](Right2[T2]{val2}, f1, f2), f2(val2)) {
			t.Fatalf("Error")
		}
	}
}
