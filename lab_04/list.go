package main

type List[T any] []T

func Pure[T any](value T) List[T] {
	return List[T]{value}
}

func Join[T any](list List[List[T]]) List[T] {
	var result List[T]
	for _, innerList := range list {
		result = append(result, innerList...)
	}
	return result
}

func MapList[T1, T2 any](l List[T1], f func(T1) T2) List[T2] {
	newList := make(List[T2], len(l))
	for i, elem := range l {
		newList[i] = f(elem)
	}
	return newList
}
