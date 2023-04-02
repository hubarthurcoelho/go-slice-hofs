package hofs

// Find receives a slice of type T and returns the first T element that returns true
// in the function passed as the 2nd argument.
func Find[T any](slice []T, f func(arg T) bool) *T {
	for _, item := range slice {
		ok := f(item)
		if ok {
			return &item
		}
	}

	return nil
}
