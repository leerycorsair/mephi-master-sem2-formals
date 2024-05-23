package main

import (
	"testing"

	"pgregory.net/rapid"
)

func TestPropCompL(t *testing.T) {
	gen := rapid.Custom(func(t *rapid.T) List[int] {
		l := rapid.SliceOfN(rapid.IntRange(0, 100), 1, 100).Draw(t, "list")
		return List[int](l)
	})

	property := PropCompL(gen)
	rapid.Check(t, property)
}

func TestPropCompR(t *testing.T) {
	gen := rapid.Custom(func(t *rapid.T) List[int] {
		l := rapid.SliceOfN(rapid.IntRange(0, 100), 1, 100).Draw(t, "list")
		return List[int](l)
	})

	property := PropCompR(gen)
	rapid.Check(t, property)
}

func TestPropComp(t *testing.T) {
	gen := rapid.Custom(func(t *rapid.T) List[List[List[int]]] {
		l1Gen := rapid.SliceOfN(rapid.IntRange(0, 100), 1, 100)
		l2Gen := rapid.SliceOfN(l1Gen, 1, 100)
		l3Gen := rapid.SliceOfN(l2Gen, 1, 100)

		var result List[List[List[int]]]
		input := l3Gen.Draw(t, "list")
		for _, outer := range input {
			var middleList List[List[int]]
			for _, middle := range outer {
				var innerList List[int]
				for _, inner := range middle {
					innerList = append(innerList, inner)
				}
				middleList = append(middleList, innerList)
			}
			result = append(result, middleList)
		}

		return result
	})

	property := PropComp(gen)
	rapid.Check(t, property)
}
