package hofs

// Filter will receive a slice of T and return a slice of the T elements
// that return true in the function passed as 2nd argument.
func Filter[T any](slice []T, f func(arg T) bool) []T {
	var filteredSlice []T

	for _, item := range slice {
		if f(item) {
			filteredSlice = append(filteredSlice, item)
		}
	}

	return filteredSlice
}
