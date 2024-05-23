package functors

type List[T any] []T

func MapList[T1, T2 any](l List[T1], f func(T1) T2) List[T2] {
	newList := make(List[T2], len(l))
	for i, elem := range l {
		newList[i] = f(elem)
	}
	return newList
}
