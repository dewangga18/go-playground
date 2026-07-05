package main

import "fmt"

// func main() {
// 	mainOperators()
// }

func mainOperators() {
	exampleNumericOperation()

	exampleUnaryOperator()

	exampleComparisonOperator()

	exampleLogicalOperator()
}

func exampleNumericOperation() {
	fmt.Println("\n=== Numeric Operation ===")
	
	const a = 10
	const b = 10
	const c = 5
	const d = 2
	const e = 3

	var i = a / b + c * d - e
	
	fmt.Println("result i (10 / 10 + 5 * 2 - 3): ", i)

	// augmented assignment operations
	fmt.Println("\n=== Augmented Assignment Operations ===")
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
	fmt.Println("\n=== Unary Operator ===")

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
	fmt.Println("\n=== Comparison Operator ===")
	
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
	fmt.Println("\n=== Logical Operator ===")
	
	trueVal := true
	falseVal := false

	fmt.Println("\n== AND (&&) ==")
	fmt.Println(trueVal && falseVal)

	fmt.Println("\n== OR (||) ==")
	fmt.Println(trueVal || falseVal)

	fmt.Println("\n== NOT (!) ==")
	fmt.Println(!trueVal)
}