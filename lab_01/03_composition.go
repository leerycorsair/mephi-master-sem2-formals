package main

func f(b int) int {
	return b * 2
}

func g(a int) int {
	return a * 3
}

func h(a int) int {
	return a + 1
}

func Composition(f1 func(int) int, f2 func(int) int) func(int) int {
	return func(a int) int {
		return f(g(a))
	}
}
