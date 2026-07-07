package main

import (
	"fmt"
	"slices"
)

func main() {
	mainSlices()
}

func mainSlices() {
	fmt.Println("=== Contains — check if element exists ===")
	numbers := []int{10, 20, 30, 40, 50}
	fmt.Println(slices.Contains(numbers, 30))  // true
	fmt.Println(slices.Contains(numbers, 99))  // false

	names := []string{"Budi", "Sari", "Agus"}
	fmt.Println(slices.Contains(names, "Sari")) // true

	fmt.Println("\n=== Index — find first index of element (-1 if not found) ===")
	fmt.Println(slices.Index(numbers, 30))      // 2
	fmt.Println(slices.Index(numbers, 99))      // -1

	fmt.Println("\n=== Equal — compare two slices ===")
	a := []int{1, 2, 3}
	b := []int{1, 2, 3}
	c := []int{1, 2, 4}
	fmt.Println(slices.Equal(a, b))  // true
	fmt.Println(slices.Equal(a, c))  // false

	// Equal with strings
	fmt.Println(slices.Equal([]string{"a", "b"}, []string{"a", "b"})) // true

	fmt.Println("\n=== Clone — create independent copy ===")
	original := []int{1, 2, 3}
	clone := slices.Clone(original)

	clone[0] = 99
	fmt.Println("original:", original)  // [1 2 3]
	fmt.Println("clone:", clone)        // [99 2 3]

	fmt.Println("\n=== Sort — sort in ascending order (in-place) ===")
	unsorted := []int{5, 2, 8, 1, 9, 3}
	slices.Sort(unsorted)
	fmt.Println(unsorted)  // [1 2 3 5 8 9]

	// Sort strings
	words := []string{"banana", "apple", "cherry"}
	slices.Sort(words)
	fmt.Println(words)     // [apple banana cherry]

	fmt.Println("\n=== SortFunc — sort with custom comparator ===")
	type Person struct {
		Name string
		Age  int
	}
	people := []Person{
		{"Budi", 30},
		{"Sari", 25},
		{"Agus", 35},
		{"Dewi", 28},
	}

	slices.SortFunc(people, func(a, b Person) int {
		if a.Age < b.Age {
			return -1
		}
		if a.Age > b.Age {
			return 1
		}
		return 0
	})
	fmt.Println(people)  // [{Sari 25} {Dewi 28} {Budi 30} {Agus 35}]

	// Sort by name descending
	slices.SortFunc(people, func(a, b Person) int {
		if a.Name > b.Name {
			return -1
		}
		if a.Name < b.Name {
			return 1
		}
		return 0
	})
	fmt.Println("Desc name:", people)  // [{Sari 25} {Dewi 28} {Budi 30} {Agus 35}]

	fmt.Println("\n=== Reverse — reverse slice in-place ===")
	nums := []int{1, 2, 3, 4, 5}
	slices.Reverse(nums)
	fmt.Println(nums)  // [5 4 3 2 1]

	words2 := []string{"a", "b", "c"}
	slices.Reverse(words2)
	fmt.Println(words2)  // [c b a]

	fmt.Println("\n=== Insert — insert elements at index ===")
	inserted := []int{1, 2, 5, 6}
	inserted = slices.Insert(inserted, 2, 3, 4)
	fmt.Println(inserted)  // [1 2 3 4 5 6]

	// Insert at beginning
	inserted = slices.Insert(inserted, 0, 0)
	fmt.Println(inserted)  // [0 1 2 3 4 5 6]

	// Insert at end (like append)
	inserted = slices.Insert(inserted, len(inserted), 7)
	fmt.Println(inserted)  // [0 1 2 3 4 5 6 7]

	fmt.Println("\n=== Delete — remove elements [i:j) ===")
	deleted := []int{0, 1, 2, 3, 4, 5, 6, 7}
	deleted = slices.Delete(deleted, 0, 2)
	fmt.Println(deleted)  // [2 3 4 5 6 7]

	// Delete from middle
	deleted = slices.Delete(deleted, 1, 3)
	fmt.Println(deleted)  // [2 5 6 7]

	// Delete last element
	deleted = slices.Delete(deleted, len(deleted)-1, len(deleted))
	fmt.Println(deleted)  // [2 5 6]

	fmt.Println("\n=== Replace — replace elements [i:j) with new ones ===")
	replaced := []int{10, 20, 30, 40, 50}
	replaced = slices.Replace(replaced, 1, 3, 25, 35)
	fmt.Println(replaced)  // [10 25 35 40 50]

	// Replace with more elements than removed
	replaced = slices.Replace(replaced, 2, 3, 100, 200, 300)
	fmt.Println(replaced)  // [10 25 100 200 300 40 50]

	fmt.Println("\n=== Compact — remove adjacent duplicates (in-place) ===")
	duplicates := []int{1, 1, 2, 2, 2, 3, 4, 4, 5}
	duplicates = slices.Compact(duplicates)
	fmt.Println(duplicates)  // [1 2 3 4 5]

	// Only removes adjacent duplicates!
	notAdjacent := []int{1, 2, 1, 2, 3}
	notAdjacent = slices.Compact(notAdjacent)
	fmt.Println(notAdjacent)  // [1 2 1 2 3] — no change because not adjacent

	// Compact strings
	strs := []string{"a", "a", "b", "b", "b", "c"}
	strs = slices.Compact(strs)
	fmt.Println(strs)  // [a b c]
}
