package main

import (
	"reflect"

	"pgregory.net/rapid"
)

func PropComp[T any](
	generator *rapid.Generator[List[List[List[T]]]]) func(*rapid.T) {
	return func(t *rapid.T) {
		src := generator.Draw(t, "src")
		left := Join(MapList(src, Join[T]))
		right := Join(Join(src))
		if !reflect.DeepEqual(left, right) {
			t.Fatalf("Error")
		}
	}
}

func PropCompR[T any](
	generator *rapid.Generator[List[T]]) func(*rapid.T) {
	return func(t *rapid.T) {
		src := generator.Draw(t, "src")
		left := Join(MapList(src, Pure[T]))
		right := src
		if !reflect.DeepEqual(left, right) {
			t.Fatalf("Error")
		}
	}
}

func PropCompL[T any](
	generator *rapid.Generator[List[T]]) func(*rapid.T) {
	return func(t *rapid.T) {
		src := generator.Draw(t, "src")
		left := Join(Pure(src))
		right := src
		if !reflect.DeepEqual(left, right) {
			t.Fatalf("Error, left %v, right %v", left, right)
		}
	}
}
