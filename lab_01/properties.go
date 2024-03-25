package main

import (
	"cmp"
	"reflect"
	"slices"

	"pgregory.net/rapid"
)

func PropCmpSorts[K cmp.Ordered](sort1 func([]K), sort2 func([]K)) func(*rapid.T) {
	return func(t *rapid.T) {
		generator := rapid.Make[[]K]()

		arr1 := generator.Draw(t, "array")
		var arr2 []K
		copy(arr2, arr1)
		sort1(arr1)
		sort2(arr2)
		if reflect.DeepEqual(arr1, arr2) {
			t.Fatalf("Error")
		}
	}
}

func PropCmpSearches[K cmp.Ordered](search1 func([]K, K) int, search2 func([]K, K) int) func(*rapid.T) {
	return func(t *rapid.T) {
		generatorArr := rapid.Make[[]K]()
		arr := generatorArr.Draw(t, "array")
		slices.Sort(arr)
		generatorInt := rapid.Make[K]()
		elem := generatorInt.Draw(t, "elem")
		if search1(arr, elem) != search2(arr, elem) {
			t.Fatalf("Error")
		}
	}
}
