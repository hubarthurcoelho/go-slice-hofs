package hofs

// Find is a generic function that accepts a slice of type T
// and a function that takes an argument of type T and returns
// a boolean value. The function applies the provided function
// to each element in the slice and returns a pointer to the
// first element for which the function returns true, or nil if
// no such element is found.
//
// Example usage:
//
//	type person struct {
//		name string
//		age int
//	}
//
//	persons := []person{
//		{name: "Alice", age: 25},
//		{name: "Bob", age: 30},
//		{name: "Joe", age: 40},
//	}
//
//	joe := Find(persons, func(p person) bool { return p.name == "Joe" })
//	// joe points to the person struct with name "Joe"
func Find[T any](slice []T, f func(arg T) bool) *T {
	for _, item := range slice {
		ok := f(item)
		if ok {
			return &item
		}
	}

	return nil
}
