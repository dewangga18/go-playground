package main

import "fmt"

func main() {
	exampleNumericOperation()

	exampleUnaryOperator()

	exampleComparisonOperator()

	exampleLogicalOperator()
}

func exampleNumericOperation() {
	fmt.Println("\n=== numeric operation ===")
	
	const a = 10
	const b = 10
	const c = 5
	const d = 2
	const e = 3

	var i = a / b + c * d - e
	
	fmt.Println("result i (10 / 10 + 5 * 2 - 3): ", i)

	// augmented assignment operations
	i += 5
	fmt.Println("result i += 5: ", i)
	i -= 5
	fmt.Println("result i -= 5: ", i)
	i *= 5
	fmt.Println("result i *= 5: ", i)
	i /= 5
	fmt.Println("result i /= 5: ", i)
	i %= 5
	fmt.Println("result i %= 5: ", i)
}

func exampleUnaryOperator() {
	fmt.Println("\n=== unary operator ===")

	var a = 10
	var b = -10
	
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	fmt.Println("-a = ", -a)
	fmt.Println("-b = ", -b)
	fmt.Println("")

	a++
	b--
	fmt.Println("a++ \n a = ", a)
	fmt.Println("b-- \n b = ", b)
}

func exampleComparisonOperator() {
	fmt.Println("\n=== comparison operator ===")
	
	const a = 1
	const b = 2

	fmt.Println("a = 1, b = 2")
	fmt.Println(a > b, "// a > b")
	fmt.Println(a < b, "// a < b")
	fmt.Println(a >= b, "// a >= b")
	fmt.Println(a <= b, "// a <= b")
	fmt.Println(a == b, "// a == b")
	fmt.Println(a != b, "// a != b")
}

func exampleLogicalOperator() {
	fmt.Println("\n=== logical operator ===")
	
	trueVal := true
	falseVal := false

	fmt.Println("\n== AND (&&) ==")
	fmt.Println(trueVal && falseVal)

	fmt.Println("\n== OR (||) ==")
	fmt.Println(trueVal || falseVal)

	fmt.Println("\n== NOT (!) ==")
	fmt.Println(!trueVal)
}