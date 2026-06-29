package main

import "fmt"

// const at package level
const (
	AppName = "Go Playground"
	Pi      = 3.14159
)

func main() {
	// === var ===
	var a int8 = 127
	var b uint8 = 255
	var c float32 = 3.14
	var d complex64 = 1 + 2i

	fmt.Println("=== var ===")
	fmt.Println("int8   =", a)
	fmt.Println("uint8  =", b)
	fmt.Println("float32 =", c)
	fmt.Println("complex64 =", d)

	// var is mutable — can reassign
	a = 100
	fmt.Println("int8 after reassign =", a)

	// === const ===
	const e = 42       // untyped constant
	const f float64 = 2.718

	fmt.Println("\n=== const ===")
	fmt.Println("const e =", e)
	fmt.Println("const f =", f)

	// const is immutable — uncommenting below will error
	// e = 99  // ❌ cannot assign to const

	// package-level const
	fmt.Println("AppName =", AppName)
	fmt.Println("Pi      =", Pi)
}
