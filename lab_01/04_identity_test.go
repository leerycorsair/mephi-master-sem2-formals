package main

import (
	"testing"

	"pgregory.net/rapid"
)

func TestIdentityL(t *testing.T) {
	property := func(t *rapid.T) {
		generatorInt := rapid.Int()
		elem := generatorInt.Draw(t, "elem")
		f := func(x int) int {
			return x * x
		}
		if f(idInt(elem)) != f(elem) {
			t.Fatalf("Error")
		}
	}
	rapid.Check(t, property)
}

func TestIdentityR(t *testing.T) {
	property := func(t *rapid.T) {
		generatorInt := rapid.Int()
		elem := generatorInt.Draw(t, "elem")
		f := func(x int) int {
			return x * x
		}
		if idInt(f(elem)) != f(elem) {
			t.Fatalf("Error")
		}
	}
	rapid.Check(t, property)
}
