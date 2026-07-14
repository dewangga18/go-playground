package main

import (
	"fmt"
	"strconv"
)

// func main() {
// 	mainStrconv()
// }

func mainStrconv() {
	result, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Result: ", result)
	}

	parseBool := strconv.FormatBool(result)
	fmt.Println("ParseBool: ", parseBool)

	result2, err2 := strconv.ParseBool("FALSE")
	if err2 != nil {
		fmt.Println("Error: ", err2)
	} else {
		fmt.Println("Result: ", result2)
	}

	result3, err3 := strconv.Atoi("123")
	if err3 != nil {
		fmt.Println("Error: ", err3)
	} else {
		fmt.Println("Result: ", result3)
	}

	result4 := strconv.Itoa(123)
	fmt.Println("Result: ", result4)

	result5, err5 := strconv.ParseFloat("123.45", 64)
	if err5 != nil {
		fmt.Println("Error: ", err5)
	} else {
		fmt.Println("Result: ", result5)
	}

	// FormatFloat(value, format, precision, bitSize)
	parseFloat := strconv.FormatFloat(result5, 'f', 2, 64)
	fmt.Println("ParseFloat: ", parseFloat)
}
