package hofs

import (
	"reflect"
	"testing"
)

type personTest struct {
	Name string
	Age  uint
}

type filterTestingSuit struct {
	receivedSlice []personTest
	expected      []personTest
}

type mapTestingSuit[T any] struct {
	receivedSlice []personTest
	expected      []T
}

type everyTestingSuit struct {
	receivedSlice []personTest
	expected      bool
}

type someTestingSuit struct {
	receivedSlice []personTest
	expected      bool
}

type findTestingSuit struct {
	receivedSlice []personTest
	expected      *personTest
}

func TestHofs(t *testing.T) {
	t.Run("Tests the hofs.Map function", func(t *testing.T) {
		// * Arrange
		testGettingAllNamesFromSlice := mapTestingSuit[string]{
			receivedSlice: []personTest{
				{Name: "Arthur", Age: 24},
				{Name: "Devis", Age: 30},
			},
			expected: []string{"Arthur", "Devis"},
		}
		testGettingAllAgesFromSlice := mapTestingSuit[uint]{
			receivedSlice: []personTest{
				{Name: "Arthur", Age: 24},
				{Name: "Devis", Age: 30},
			},
			expected: []uint{24, 30},
		}

		// * Act
		result1 := Map(testGettingAllNamesFromSlice.receivedSlice, func(person personTest) string { return person.Name })
		result2 := Map(testGettingAllAgesFromSlice.receivedSlice, func(person personTest) uint { return person.Age })

		// * Assert
		if !reflect.DeepEqual(result1, testGettingAllNamesFromSlice.expected) {
			t.Errorf("Map failed: received %+v, expected %+v", result1, testGettingAllNamesFromSlice.expected)
		}
		if !reflect.DeepEqual(result2, testGettingAllAgesFromSlice.expected) {
			t.Errorf("Map failed: received %+v, expected %+v", result2, testGettingAllAgesFromSlice.expected)
		}
	})

	t.Run("Tests the hofs.Filter function", func(t *testing.T) {
		// * Arrange
		testYoungerThan30 := filterTestingSuit{
			receivedSlice: []personTest{
				{Name: "Arthur", Age: 24},
				{Name: "Devis", Age: 30},
			},
			expected: []personTest{
				{Name: "Arthur", Age: 24},
			},
		}
		test30OrOlder := filterTestingSuit{
			receivedSlice: []personTest{
				{Name: "Arthur", Age: 24},
				{Name: "Devis", Age: 30},
			},
			expected: []personTest{
				{Name: "Devis", Age: 30},
			},
		}

		// * Act
		result1 := Filter(testYoungerThan30.receivedSlice, func(person personTest) bool { return person.Age < 30 })
		result2 := Filter(test30OrOlder.receivedSlice, func(person personTest) bool { return person.Age >= 30 })

		// * Assert
		if !reflect.DeepEqual(result1, testYoungerThan30.expected) {
			t.Errorf("Filter failed: received %+v, expected %+v", result1, testYoungerThan30.expected)
		}
		if !reflect.DeepEqual(result2, test30OrOlder.expected) {
			t.Errorf("Filter failed: received %+v, expected %+v", result2, test30OrOlder.expected)
		}
		}
	})

	t.Run("Tests the hofs.Every function", func(t *testing.T) {
		// * Arrange
		testIfEveryoneIsAKid := everyTestingSuit{
			receivedSlice: []personTest{
				{Name: "Elisa", Age: 5},
				{Name: "Sofia", Age: 4},
			},
			expected: true,
		}
		testIfEveryoneIsOver18 := everyTestingSuit{
			receivedSlice: []personTest{
				{Name: "Larissa", Age: 17},
				{Name: "Devis", Age: 30},
			},
			expected: false,
		}

		// * Act
		result1 := Every(testIfEveryoneIsAKid.receivedSlice, func(person personTest) bool { return person.Age < 12 })
		result2 := Every(testIfEveryoneIsOver18.receivedSlice, func(person personTest) bool { return person.Age >= 18 })

		// * Assert
		if result1 != testIfEveryoneIsAKid.expected {
			t.Errorf("Every failed: received %+v, expected %+v", result1, testIfEveryoneIsAKid.expected)
		}
		if result2 != testIfEveryoneIsOver18.expected {
			t.Errorf("Every failed: received %+v, expected %+v", result2, testIfEveryoneIsOver18.expected)
		}
		}
	})

	t.Run("Tests the hofs.Some function", func(t *testing.T) {
		// * Arrange
		testIfSomeoneIsAKid := someTestingSuit{
			receivedSlice: []personTest{
				{Name: "Elisa", Age: 5},
				{Name: "Devis", Age: 30},
			},
			expected: true,
		}
		testIfSomeoneIsOver18 := someTestingSuit{
			receivedSlice: []personTest{
				{Name: "Larissa", Age: 17},
				{Name: "Elisa", Age: 5},
			},
			expected: false,
		}

		// * Act
		result1 := Some(testIfSomeoneIsAKid.receivedSlice, func(person personTest) bool { return person.Age < 12 })
		result2 := Some(testIfSomeoneIsOver18.receivedSlice, func(person personTest) bool { return person.Age >= 18 })

		// * Assert
		if result1 != testIfSomeoneIsAKid.expected {
			t.Errorf("Some failed: received %+v, expected %+v", result1, testIfSomeoneIsAKid.expected)
		}
		if result2 != testIfSomeoneIsOver18.expected {
			t.Errorf("Some failed: received %+v, expected %+v", result2, testIfSomeoneIsOver18.expected)
		}
	})

	t.Run("Tests the hofs.Find function", func(t *testing.T) {
		// * Arrange
		testFindTheFirstKid := findTestingSuit{
			receivedSlice: []personTest{
				{Name: "Elisa", Age: 5},
				{Name: "Devis", Age: 30},
			},
			expected: &personTest{Name: "Elisa", Age: 5},
		}
		testFindFirstPersonOver18 := findTestingSuit{
			receivedSlice: []personTest{
				{Name: "Elisa", Age: 5},
				{Name: "Lucas", Age: 23},
			},
			expected: &personTest{Name: "Lucas", Age: 23},
		}

		// * Act
		result1 := Find(testFindTheFirstKid.receivedSlice, func(person personTest) bool { return person.Age < 12 })
		result2 := Find(testFindFirstPersonOver18.receivedSlice, func(person personTest) bool { return person.Age >= 18 })

		// * Assert
		if !reflect.DeepEqual(result1, testFindTheFirstKid.expected) {
			t.Errorf("Find failed: received %+v, expected %+v", result1, testFindTheFirstKid.expected)
		}
		if !reflect.DeepEqual(result2, testFindFirstPersonOver18.expected) {
			t.Errorf("Find failed: received %+v, expected %+v", result2, testFindFirstPersonOver18.expected)
		}
	})
}
