package pkg

func Map[T any, A any](slice []T, fn func(T) A) []A {
	list := make([]A, len(slice))
	for i := range slice {
		list[i] = fn(slice[i])
	}
	return list
}
