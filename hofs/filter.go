package hofs

// Filter is a generic function that accepts a slice of type T
// and a function that takes an argument of type T and returns
// a boolean value. The function applies the provided function
// to each element in the slice and returns a new slice containing
// only those elements for which the function returns true.
//
// Example usage:
//
//	ints := []int{1, 2, 3, 4, 5}
//	evens := Filter(ints, func(n int) bool { return n%2 == 0 })
//	// evens will be an int slice with the ints 2 and 4
func Filter[T any](slice []T, f func(arg T) bool) []T {
	var filteredSlice []T

	for _, item := range slice {
		if f(item) {
			filteredSlice = append(filteredSlice, item)
		}
	}

	return filteredSlice
}
