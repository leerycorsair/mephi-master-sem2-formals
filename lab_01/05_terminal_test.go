package main

import (
	"testing"

	"pgregory.net/rapid"
)

func TestTerminal(t *testing.T) {
	property := func(t *rapid.T) {
		generatorInt := rapid.Int()
		generatorString := rapid.String()
		elemInt := generatorInt.Draw(t, "elemInt")
		elemString := generatorString.Draw(t, "elemStr")
		f := func(x any) Unit { return Unit{} }
		g := func(x any) Unit { return Unit{} }

		if f(elemInt) != g(elemString) {
			t.Fatalf("Error")
		}
	}
	rapid.Check(t, property)
}
