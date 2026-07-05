package main

import (
	"errors"
	"fmt"
)

// func main() {
// 	mainError()
// }

func mainError() {
	fmt.Println("\n=== Error Example ===")
	res, err := divide(10, 0)
	if err != nil {
		fmt.Println("ERROR:", err)
	} else {
		fmt.Println("RESULT:", res)
	}
	res2, err2 := divide(10, 2)
	if err2 != nil {
		fmt.Println("ERROR:", err2)
	} else {
		fmt.Println("RESULT:", res2)
	}

	fmt.Println("\n=== Custom Error Example ===")
	errSave1 := saveData("", "data1")
	if errSave1 != nil {
		fmt.Println("ERROR:", errSave1)
	}

	errSave2 := saveData("root", "data2")
	if errSave2 != nil {
		fmt.Println("ERROR:", errSave2)
	}

	errSave3 := saveData("Aaron", "data3")
	if errSave3 != nil {
		fmt.Println("ERROR:", errSave3)
	}

	switch err := errSave1.(type) {
	case *validationError:
		fmt.Println("Validation Error:", err.Message)
	case *notFoundError:
		fmt.Println("Not Found Error:", err.Message)
	default:
		fmt.Println("Unknown Error:", err)
	}
}

func divide(num1, div int) (int, error) {
	if div == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return num1 / div, nil
}

type validationError struct {
	Message string
}

func (ve *validationError) Error() string {
	return ve.Message
}

type notFoundError struct {
	Message string
}

func (ve *notFoundError) Error() string {
	return ve.Message
}

func saveData(name string, data any) error {
	if name == "" {
		return &validationError{
			Message: "name is required",
		}
	}

	if name != "root" {
		return &notFoundError{
			Message: "username not found",
		}
	}

	fmt.Println("Saving data:", name)
	return nil
}
