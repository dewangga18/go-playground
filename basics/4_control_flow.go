package main

import "fmt"

// func main() {
// 	mainControlFlow()
// }

func mainControlFlow() {
	exampleIfExpression()

	exampleSwitch()

	exampleForLoop()
}

func exampleIfExpression() {
	fmt.Println("\n=== If Expression ===")
	score := 60
	grade := ""

	if score >= 80 {
		grade = "A"
	} else if score >= 70 {
		grade = "B"
	} else if score >= 60 {
		grade = "C"
	} else if score >= 50 {
		grade = "D"
	} else {
		grade = "E"
	}
	fmt.Println("Your score: ", score, "\nYour grade is ", grade)

	fmt.Println("\n== Short statement ==")
	examplePassword := "123456"

	if length := len(examplePassword); length >= 8 {
		fmt.Println("Password is valid")
	} else {
		fmt.Println("Password is too short")
	}
}

func exampleSwitch() {
	fmt.Println("\n=== Switch Expression ===")
	score := 90
	grade := ""

	switch {
	case score >= 90:
		grade = "A"
	case score >= 80:
		grade = "B"
	case score >= 70:
		grade = "C"
	case score >= 60:
		grade = "D"
	case score >= 50:
		grade = "D"
	default:
		grade = "E"
	}
	fmt.Println("Your score: ", score, "\nYour grade is ", grade)

	switch {
	case grade == "A":
		fmt.Println("Excellent!")
	case grade == "B":
		fmt.Println("Good!")
	case grade == "C":
		fmt.Println("Average!")
	case grade == "D":
		fmt.Println("Poor!")
	default:
		fmt.Println("Fail!")
	}

	fmt.Println("\n== Short statement ==")
	examplePassword := "123456"

	switch length := len(examplePassword); length >= 8 {
	case true:
		fmt.Println("Password is valid")
	default:
		fmt.Println("Password is too short")
	}
}

func exampleForLoop() {
	fmt.Println("\n=== For Loop ===")
	
	counter := 1

	for counter <= 10 {
		fmt.Println("Iteration: ", counter)
		counter++
	}

	fmt.Println("\n== For with Initialization ==")

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Iteration: ", counter)
	}

	fmt.Println("\n== For with Range over slice ==")
	colors := []string{"Red", "Yellow", "Green", "Blue"}

	for i, color := range colors {
		fmt.Println("Index: ", i, "Color: ", color)
	}
	
	fmt.Println("\n== For with break")
	for _, color := range colors {
		if color == "Green" {
			break
		}
		fmt.Println("Color: ", color)
	}
	
	fmt.Println("\n== For with continue")
	for _, color := range colors {
		if color == "Green" {
			continue
		}
		fmt.Println("Color: ", color)
	}
}