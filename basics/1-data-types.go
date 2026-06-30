package main

import "fmt"

func main() {
	exampleNumericVarAndConst()

	exampleBoolean()

	exampleString()

	exampleVariableDeclaration()

	exampleDataConversation()

	exampleTypeDeclaration()

	exampleNumericOperation()
}

func exampleNumericVarAndConst() {
	// const at package level
	const (
		AppName = "Go Playground"
		Pi      = 3.14159
	)

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

func exampleBoolean() {
	fmt.Println("\n=== boolean ===")
	var trueConstant bool = true
	var falseConstant bool = false
	const numb1 int = 1

	fmt.Println("True constant = ", trueConstant)
	fmt.Println("False constant = ", falseConstant)
	fmt.Println("1 == 1 ", 1 == numb1)
	fmt.Println("1 != 1 ", 1 != numb1)
}

func exampleString() {
	fmt.Println("\n=== string ===")
	var str1 string = "Hello"

	fmt.Println(str1)
	fmt.Println("length of str1 is ", len(str1))
	fmt.Println("index 0 is ", str1[0], "value is byte")
}

func exampleVariableDeclaration() {
	fmt.Println("\n=== variable declaration ===")
	var tempA int = 1
	var tempB int = 2

	// can not swap like this
	// tempA, tempB = tempB, tempA

	// use temporary variable to swap
	tmp := tempA
	tempA = tempB
	tempB = tmp

	fmt.Println("tempA", tempA)
	fmt.Println("tempB", tempB)

	// multiple declaration
	var(
		firstName = "Aaron"
		lastName  = "Evanjulio"
	)

	fmt.Println(firstName, lastName)
}

func exampleDataConversation() {
	firstName := "Aaron"
	fmt.Println("\n=== data conversation ===")
	var byteVal uint8 = firstName[0] 
	var byteToStr = string(byteVal)
	fmt.Println("Value of byteVal from firstName[0]: ", byteVal)
	fmt.Println("Convert byte to string: ", byteToStr)

	var val32 int32 = 32769
	var val64 int64 = int64(val32)
	var val16 int16 = int16(val32)

	fmt.Println("int32 to int64: ", val64)
	fmt.Println("int32 to int16: ", val16, "// number overflow")
}

func exampleTypeDeclaration() {
	fmt.Println("\n=== type declarations ===")
	type WhatsappNumber string

	var w1 WhatsappNumber = "08123456789"
	var w2 string = "08123456789"

	fmt.Println(w1)
	fmt.Println(WhatsappNumber(w2))
	fmt.Println("is equal? ", w1 == WhatsappNumber(w2))
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