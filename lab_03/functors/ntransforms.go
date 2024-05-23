package functors

func FromListToMaybe[T any](l List[T]) Maybe[T] {
	if len(l) == 0 {
		return Nothing[T]()
	}
	return Just(l[0])
}
