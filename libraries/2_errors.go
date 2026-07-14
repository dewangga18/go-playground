package main

import (
	"fmt"
	"errors"
)

// func main() {
// 	mainErrors()
// }

func mainErrors() {
	err := getById("")
	checkErrors(err)
	err = getById("upsss")
	checkErrors(err)
	err = getById("root")
	checkErrors(err)
}



var (
	ValidationError = errors.New("validation error")
	NotFoundError = errors.New("data not found")
)

func getById(id string) error {
	if id == "" {
		return ValidationError
	}
	
	if id != "root" {
		return NotFoundError
	}

	return nil
}

func checkErrors(err error) {
	switch err {
	case ValidationError:
		fmt.Println("Validation Error")
	case NotFoundError:
		fmt.Println("Not Found Error")
	case nil:
		fmt.Println("Success")
	default:
		fmt.Println("Unknown Error")
	}

	// option 2
	if errors.Is(err, ValidationError) {
		fmt.Println("Validation Error")
	} else if errors.Is(err, NotFoundError) {
		fmt.Println("Not Found Error")
	} else if err == nil {
		fmt.Println("Success")
	} else {
		fmt.Println("Unknown Error")
	}
}
