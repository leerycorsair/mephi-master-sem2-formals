package main

import (
	"testing"

	"pgregory.net/rapid"
)

func TestCompositionAssociativity(t *testing.T) {
	property := func(t *rapid.T) {
		generatorInt := rapid.Int()
		elem := generatorInt.Draw(t, "elem")

		cmp1 := Composition(Composition(f, g), h)
		cmp2 := Composition(f, Composition(g, h))

		if cmp1(elem) != cmp2(elem) {
			t.Fatalf("Error")
		}
	}
	rapid.Check(t, property)
}
