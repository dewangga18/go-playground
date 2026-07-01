package main

import (
	"fmt"
	"strconv"
)

// func main() {
// 	basicFunction()

// 	functionParameter("Aaron", "Evanjulio")

// 	area := returnValue(10, 5)
// 	fmt.Println("Area is", area)

// 	fmt.Println("\n=== Return Value Function 2 ===")
// 	area2, size2 := returnValue2(20, 15)
// 	fmt.Println("Area is", area2, "and size is", size2)
// 	area3, _ := returnValue2(15, 20)
// 	fmt.Println("Area is", area3)

// 	area4, size4 := namedReturnValue(10, 10)
// 	fmt.Println("Area is", area4, "and size is", size4)

// 	fmt.Println("\n=== Variadic Function ===")
// 	average := variadicFunction(10, 20, 30, 40, 50)
// 	fmt.Println("Average is", average)
// 	sliceWithVariadic := []int{10, 20, 30, 40, 50}
// 	average = variadicFunction(sliceWithVariadic...)
// 	fmt.Println("Average is", average)

// 	goodBye := functionAsValue
// 	fmt.Println(goodBye("Aaron"))

// 	fmt.Println("\n=== Function As Params ===")
// 	functionAsParams(11, filterOddNumber)
// 	functionAsParams(12, filterOddNumber)
// 	functionAsParams2(13, filterOddNumber)
// 	functionAsParams2(14, filterOddNumber)

// 	fmt.Println("\n=== Anonymous Function ===")
// 	anonymousFunction()
// 	anonymousFunction2()
// 	anonymousFunction3()

// 	fmt.Println("\n=== Recursive Factorial Function ===")
// 	fmt.Println(recursiveFactorialFunction(10), "Example from 10")

// 	closureFunction()
// }

func basicFunction() {
	fmt.Println("\n=== Basic Function ===")
	fmt.Println("Hi!")
}

func functionParameter(firstName string, lastName string) {
	fmt.Println("\n=== Function Parameter ===")
	fmt.Println("Hello", firstName, lastName)
}

func returnValue(width int, height int) int {
	fmt.Println("\n=== Return Value Function ===")
	return width * height
}

func returnValue2(width int, height int) (int, string) {
	area := width * height
	if area >= 10 {
		return area, "Large"
	}
	return area, "Small"
}

func namedReturnValue(width int, height int) (area int, size string) {
	fmt.Println("\n=== Named Return Value Function ===")
	area = width * height
	if area >= 10 {
		size = "Large"
	}
	return area, size
}

func variadicFunction(numbers ...int) (total int) {
	for _, number := range numbers {
		total += number
	}

	average := total / len(numbers)
	return average
}

func functionAsValue(name string) string {
	fmt.Println("\n=== Function As Value ===")
	return "Good bye, " + name + "!"
}

func filterOddNumber(number int) string {
	if number%2 == 1 {
		return strconv.Itoa(number) + " is Odd"
	}
	return strconv.Itoa(number) + " is Even"
}

func functionAsParams(number int, filter func(int) string) {
	fmt.Println(filter(number))
}

type Filter func(int) string

func functionAsParams2(number int, filter Filter) {
	fmt.Println(filter(number))
}

func anonymousFunction() {
	greet := func(name string) {
		fmt.Println("Hello,", name)
	}

	greet("John")
}

func anonymousFunction2() {
	square := func(number int) int {
		return number * number
	}

	result := square(5)

	fmt.Println(result)
}

func anonymousFunction3() {
	numbers := []int{1, 2, 3}

	for _, number := range numbers {
		func() {
			fmt.Println(number)
		}()
	}
}

func recursiveFactorialFunction(value int) int {
	if value == 1 {
		return value
	}
	return value * recursiveFactorialFunction(value-1)
}

func closureFunction() {
	fmt.Println("\n=== Closure Function ===")

	counter := 0
	
	closure := func() {
		fmt.Println("Increment")
		counter++
	}

	closure()
	closure()
	closure()

	fmt.Println(counter)
}