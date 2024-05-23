package functors

import (
	"reflect"

	"pgregory.net/rapid"
)

func PropMaybeId[T1, T2 any](
	f func(T1) T2,
	cmp func(T2, T2) bool,
	generator *rapid.Generator[T1]) func(*rapid.T) {
	return func(t *rapid.T) {
		switch rapid.IntRange(0, 1).Draw(t, "switch") {
		case 0:
			maybeObj := Nothing[T1]()
			if !MaybeCmp(MapMaybe(maybeObj, f), Nothing[T2](), cmp) {
				t.Fatalf("Error")
			}
		case 1:
			src := generator.Draw(t, "src")
			maybeObj := Just(src)
			if !MaybeCmp(MapMaybe(maybeObj, f), Just(f(src)), cmp) {
				t.Fatalf("Error")
			}
		}
	}
}

func PropMaybeComp[T1, T2, T3 any](
	f1 func(T1) T2,
	f2 func(T2) T3,
	cmp func(T3, T3) bool,
	generator *rapid.Generator[T1]) func(*rapid.T) {
	return func(t *rapid.T) {
		switch rapid.IntRange(0, 1).Draw(t, "switch") {
		case 0:
			maybeObj := Nothing[T1]()
			if !MaybeCmp(MapMaybe(MapMaybe(maybeObj, f1), f2), Nothing[T3](), cmp) {
				t.Fatalf("Error")
			}
		case 1:
			src := generator.Draw(t, "src")
			maybeObj := Just(src)
			if !MaybeCmp(MapMaybe(MapMaybe(maybeObj, f1), f2), Just(f2(f1(src))), cmp) {
				t.Fatalf("Error")
			}
		}
	}
}

func PropNtId[T any](
	generator *rapid.Generator[List[T]]) func(*rapid.T) {
	return func(t *rapid.T) {
		src := generator.Draw(t, "src")
		dst := FromListToMaybe(src)
		if reflect.ValueOf(dst).Type() != reflect.TypeFor[Maybe[T]]() {
			t.Fatalf("Error")
		}
	}
}

func PropNtComp[T1, T2 any](
	f func(T1) T2,
	cmp func(T2, T2) bool,
	generator *rapid.Generator[List[T1]]) func(*rapid.T) {
	return func(t *rapid.T) {
		src := generator.Draw(t, "src")
		if !MaybeCmp(MapMaybe(FromListToMaybe(src), f), FromListToMaybe(MapList(src, f)), cmp) {
			t.Fatalf("Error")
		}
	}
}
