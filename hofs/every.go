package hofs

// Every is a generic function that accepts a slice of type T and a function f
// that accepts a single argument of type T and returns a boolean value.
// The function returns true if all elements in the slice satisfy the
// condition specified by the f function, and false otherwise.
func Every[T any](slice []T, f func(arg T) bool) bool {
	every := true

	for _, item := range slice {
		ok := f(item)

		if !ok {
			every = false
			break
		}
	}

	return every
}
