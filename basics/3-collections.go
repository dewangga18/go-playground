package main

import "fmt"

func main() {
	exampleArray()

	exampleSlice()

	exampleMap()
}

func exampleArray() {
	fmt.Println("=== Array ===")

	var names [3]string
	names[0] = "Aaron"
	names[1] = "Evanjulio"
	names[2] = "Dewangga"
	fmt.Println(names)
	fmt.Println("First name: ", names[0])
	fmt.Println("Second name: ", names[1])
	fmt.Println("Third name: ", names[2])
	fmt.Println()

	var values = [4]int{1, 2, 3}
	fmt.Println(values)
	fmt.Println("Length of values: ", len(values))
	fmt.Println("Capacity of values: ", cap(values))
	fmt.Println()

	// compiler calculates capacity
	var computedCapacity = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(computedCapacity)
	fmt.Println("Capacity of computedCapacity: ", cap(computedCapacity))
}

func exampleSlice() {
	fmt.Println("\n=== Slice ===")

	// sample array
	var days = [...]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	fmt.Println(days)

	// create slice from array
	sliceList := days[1:4]
	fmt.Println("Slice from 'days[1:4]' → ", sliceList)
	fmt.Println("Pointer (low value of slice) → 1")
	fmt.Println("Length (high value of slice) → ", len(sliceList))
	fmt.Println("Capacity (Capacity of array - low value of slice) → ", cap(sliceList))
	fmt.Println()

	tmp1 := days[4:]
	tmp2 := days[:4]
	tmp3 := days[:]

	fmt.Println("tmp1 from 'days[4:]' → ", tmp1)
	fmt.Println("tmp2 from 'days[:4]' → ", tmp2)
	fmt.Println("tmp3 from 'days[:]' → ", tmp3)
	fmt.Println()

	// example slice from array
	fmt.Println("Example of slice from array")
	colors := [...]string{"Orange", "Yellow", "Brown", "Green", "Purple", "Blue"}
	tmpColors1 := colors[3:]
	fmt.Println("Initial state:", colors)
	fmt.Println("tmpColors1 from 'colors[3:]' → ", tmpColors1)

	fmt.Println("\nChange value slice effect on array")
	tmpColors1[0] = "Lime"
	fmt.Println("tmpColors1 after update: ", tmpColors1)
	fmt.Println("colors after update: ", colors)

	fmt.Println("\nAppend data on slice will create new array")
	tmpColors2 := append(tmpColors1, "White", "Black")
	fmt.Println("tmpColors2 after append: ", tmpColors2)
	fmt.Println("colors after append: ", colors)

	tmpColors2[2] = "Cyan"
	fmt.Println("\nUpdate on slice will create new array if capacity is full")
	fmt.Println("tmpColors2 after update: ", tmpColors2)
	fmt.Println("colors after update: ", colors)

	fmt.Println("\nMake Function slice make([]string, 2, 10) → [type] [length] [capacity]")
	newColors := make([]string, 2, 10)
	newColors[0] = "Maroon"
	newColors[1] = "Magenta"

	fmt.Println(newColors)
	fmt.Println("Length of newColors: ", len(newColors))
	fmt.Println("Capacity of newColors: ", cap(newColors))

	fmt.Println("\nAppend data on slice with capacity")
	newColors2 := append(newColors, "Navy")
	fmt.Println(newColors2)
	fmt.Println("Length of newColors2: ", len(newColors2))
	fmt.Println("Capacity of newColors2: ", cap(newColors2))
	
	fmt.Println("\nCompare with array")
	thisArray := [...]int {1, 2, 3}
	thisSlice := []int{1,2,3}

	fmt.Println("thisArray:", thisArray, "\tLength:", len(thisArray), "\tCapacity:", cap(thisArray))
	fmt.Println("thisSlice:", thisSlice, "\tLength:", len(thisSlice), "\tCapacity:", cap(thisSlice))
	
	fmt.Println("\nArray vs Slice with append function")
	// thisArray = append(thisArray, 4) // compiler error
	thisSlice = append(thisSlice, 4)
	fmt.Println("thisArray:", thisArray, "\tLength:", len(thisArray), "\tCapacity:", cap(thisArray))
	fmt.Println("thisSlice:", thisSlice, "\tLength:", len(thisSlice), "\tCapacity:", cap(thisSlice)) // return 2x capacity if capacity is full

	thisSlice = append(thisSlice, 5, 6, 7, 8)
	fmt.Println("\nAppend multiple data")
	fmt.Println("thisSlice:", thisSlice, "\tLength:", len(thisSlice), "\tCapacity:", cap(thisSlice))
}

func exampleMap() {
	fmt.Println("\n=== Map ===")

	person := map[string]string {
		"name": "Aaron",
		"age": "22",
		"city": "Malang",
	}

	fmt.Println(person, "Length:", len(person))
	fmt.Println("Name: ", person["name"])
	fmt.Println("Age: ", person["age"])
	fmt.Println("City: ", person["city"])

	delete(person, "age")
	fmt.Println("After delete age:", person)

	wrongKey := person["jobs"]
	fmt.Println("Call wrong key", wrongKey) // return empty

	fmt.Println("\nCreate new map with make function")
	device := make(map[string]any)
	device["name"] = "iQOO"
	device["os"] = "android"
	device["ram"] = 8
	device["rom"] = 256
	
	fmt.Println(device)
}
