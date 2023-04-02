package hofs

// Some is a generic function that accepts a slice of type T and a function f
// that accepts a single argument of type T and returns a boolean value.
// Some returns true if at least one element in the slice satisfies
// the condition specified by the f function, and false otherwise.
func Some[T any](slice []T, f func(arg T) bool) bool {
	some := false

	for _, item := range slice {
		ok := f(item)
		if ok {
			some = true
			break
		}
	}

	return some
}
