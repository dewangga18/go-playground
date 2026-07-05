package main

import "fmt"

// func main() {
// 	mainHandling()
// }

func mainHandling() {
	exampleDefer()

	examplePanicAndDefer(false)
	// examplePanicAndDefer(true)

	examplePanicAndRecover()
}

func exampleDefer() {
	fmt.Println("\n=== Defer Handling ===")
	defer fmt.Println("end session defer")

	fmt.Println("start session")
}

func examplePanicAndDefer(err bool) {
	fmt.Println("\n=== Panic and Defer Handling ===")

	defer fmt.Println("end session defer")

	if err {
		panic("something went wrong")
	}

	fmt.Println("function end")
}

func deferAndRecover() {
	fmt.Println("end session defer")
	msg := recover()
	fmt.Println("recover :", msg)
}

func examplePanicAndRecover() {
	fmt.Println("\n=== Panic and Recover Handling ===")

	defer deferAndRecover()
	panic("something went wrong")
}