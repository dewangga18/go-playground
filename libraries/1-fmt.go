package main

import "fmt"

// func main() {
// 	mainFmt()
// }

func mainFmt() {
	fmt.Println("Hello, World!")

	var name string = "Aaron"
	var age int = 22
	var isStudent bool = true
	var height float64 = 175.5

	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Age: %d\n", age)
	fmt.Printf("Is Student: %t\n", isStudent)
	fmt.Printf("Height: %.1f\n", height)

	var number float64 = 12345.6789

	fmt.Printf("%f\n", number)
	fmt.Printf("%.2f\n", number)
	fmt.Printf("%e\n", number)
	fmt.Printf("%E\n", number)

}