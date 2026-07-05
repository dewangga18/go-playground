package main

import "fmt"

// func main() {
// 	fmt.Println("=== Interface Example ===")
// 	var person HasName
// 	person = Person{"Aaron Evanjulio"}
// 	PrintName(person)
	
// 	fmt.Println("\n=== Empty Interface Example ===")
// 	var emptyOld = Upss()
// 	fmt.Println(emptyOld)
// 	var emptyNew = Ups()
// 	fmt.Println(emptyNew)

// 	fmt.Println("\n=== Nil Example ===")
// 	example := newExample("")
// 	fmt.Println(example)
// 	val := example["name"]
// 	if val == "" {
// 		fmt.Println("Value is empty")
// 	} else {
// 		fmt.Println(val)
// 	}
// 	example = newExample("Aaron")
// 	fmt.Println(example["name"])

// 	fmt.Println("\n=== Type Assertion Example ===")
// 	result := random()
// 	resultString := result.(string)
// 	fmt.Println(resultString)

// 	// fmt.Println(result.(int)) // trigger panic
// 	// better use switch
// 	switch value := result.(type) {
// 	case string:
// 		fmt.Println("String", value)
// 	case int:
// 		fmt.Println("Int", value)
// 	default:
// 		fmt.Println("Unknown", value)
// 	}
// }

// interface section
type HasName interface {
	GetName() string
}

func PrintName(h HasName) {
	fmt.Println(h.GetName())
}

// struct section
type Person struct {
	name string
}

func (p Person) GetName() string {
	return p.name
}

// empty interface or any
func Ups() any {
	return "Ups new version"
}

func Upss() interface{} { // this is old version, and get deprecated
	return "Upss old version"
}

// nil section
func newExample(name string) map[string]string {
	if name == "" {
		return nil
	}
	return map[string]string{
		"name": name,
	}
}

// type assertion
func random() any {
	return "OK"
}

