package hofs

// Map will receive a slice of T values and return a slice of Y values of the same length.
func Map[T any, Y any](slice []T, f func(arg T) Y) []Y {
	result := make([]Y, 0, len(slice))
	for _, item := range slice {
		newItem := f(item)
		result = append(result, newItem)
	}

	return result
}
