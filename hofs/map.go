package hofs

func Map[T any, Y any](slice []Y, f func(arg Y) T) []T {
	result := make([]T, 0, len(slice))
	for _, item := range slice {
		newItem := f(item)
		result = append(result, newItem)
	}

	return result
}
