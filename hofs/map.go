package hofs

// Map is a generic function that applies a function f to each element of a slice of type T,
// producing a new slice of type Y with the same length. The function f takes an argument of type T
// and returns a value of type Y. The resulting slice is returned.
//
// Example usage:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	doubleNumbers := Map(numbers, func(n int) int { return n * 2 })
//	// doubleNumbers is now []int{2, 4, 6, 8, 10}
func Map[T any, Y any](slice []T, f func(arg T) Y) []Y {
	result := make([]Y, 0, len(slice))
	for _, item := range slice {
		newItem := f(item)
		result = append(result, newItem)
	}

	return result
}
