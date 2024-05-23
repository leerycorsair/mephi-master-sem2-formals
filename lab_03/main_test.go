package main

import (
	"local/functors"
	"math"
	"testing"

	"pgregory.net/rapid"
)

var PRECISION float64 = 1e-9

func TestPropMaybeId(t *testing.T) {
	gen := rapid.IntRange(0, 100)

	intCmp := func(a, b int) bool {
		return a == b
	}
	property := functors.PropMaybeId(sqr, intCmp, gen)
	rapid.Check(t, property)
}

func TestPropMaybeComp(t *testing.T) {
	gen := rapid.IntRange(0, 100)

	cmp := func(a, b float64) bool {
		return math.Abs(a-b) < PRECISION
	}
	property := functors.PropMaybeComp(sqr, double, cmp, gen)
	rapid.Check(t, property)
}

func TestPropNtId(t *testing.T) {
	gen := rapid.Custom(func(t *rapid.T) functors.List[int] {
		l := rapid.SliceOf(rapid.IntRange(0, 100)).Draw(t, "list")
		return functors.List[int](l)
	})
	property := functors.PropNtId(gen)
	rapid.Check(t, property)
}

func TestPropNtComp(t *testing.T) {
	gen := rapid.Custom(func(t *rapid.T) functors.List[int] {
		l := rapid.SliceOf(rapid.IntRange(0, 100)).Draw(t, "list")
		return functors.List[int](l)
	})

	cmp := func(a, b float64) bool {
		return math.Abs(a-b) < PRECISION
	}
	property := functors.PropNtComp(double, cmp, gen)
	rapid.Check(t, property)
}
