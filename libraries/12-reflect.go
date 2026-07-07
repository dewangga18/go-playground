package main

import (
	"fmt"
	"reflect"
)

// func main() {
// 	mainReflect()
// }

func mainReflect() {
	fmt.Println("=== TypeOf & ValueOf ===")
	sample := Sample{"Uhuyy", "23"}
	sampleType := reflect.TypeOf(sample)
	sampleValue := reflect.ValueOf(sample)

	fmt.Println("TypeOf:", sampleType.Name())      // Sample
	fmt.Println("ValueOf Name:", sampleValue.FieldByName("Name").String())  // Uhuyy
	fmt.Println("ValueOf Age:", sampleValue.FieldByName("Age").String())    // 23

	fmt.Println("\n=== Kind ===")
	fmt.Println("Kind of Sample:", sampleType.Kind())   // struct

	ptr := &sample
	fmt.Println("Kind of ptr:", reflect.TypeOf(ptr).Kind())  // ptr

	var nums []int
	fmt.Println("Kind of slice:", reflect.TypeOf(nums).Kind()) // slice

	var m map[string]int
	fmt.Println("Kind of map:", reflect.TypeOf(m).Kind())     // map

	fmt.Println("\n=== Struct Fields ===")
	for i := 0; i < sampleType.NumField(); i++ {
		field := sampleType.Field(i)
		fmt.Printf("  Field %d: %s (%s)\n", i, field.Name, field.Type)
	}

	fmt.Println("\n=== Struct Tags ===")
	for i := 0; i < sampleType.NumField(); i++ {
		field := sampleType.Field(i)
		required := field.Tag.Get("required")
		max := field.Tag.Get("max")
		fmt.Printf("  %s → required: %q, max: %q\n", field.Name, required, max)
	}

	fmt.Println("\n=== Elem (dereference pointer) ===")
	num := 42
	ptrToNum := &num
	ptrValue := reflect.ValueOf(ptrToNum)
	elem := ptrValue.Elem() // dereference → gets the int value

	fmt.Println("  ptrValue.Kind():", ptrValue.Kind())   // ptr
	fmt.Println("  elem.Kind():", elem.Kind())            // int
	fmt.Println("  elem.Int():", elem.Int())              // 42

	fmt.Println("\n=== CanSet & Set (modify through pointer) ===")
	fmt.Println("  elem.CanSet():", elem.CanSet())        // true (addressable)
	elem.SetInt(100)
	fmt.Println("  num after SetInt:", num)               // 100

	// Modify struct fields via pointer
	p := Person{"Budi", 25}
	pv := reflect.ValueOf(&p).Elem()                      // must pass pointer!

	fmt.Println("  CanSet Name:", pv.FieldByName("Name").CanSet())  // true
	fmt.Println("  CanSet Age:", pv.FieldByName("Age").CanSet())    // true

	pv.FieldByName("Name").SetString("Agus")
	pv.FieldByName("Age").SetInt(30)
	fmt.Println("  Person after Set:", p)                  // {Agus 30}

	fmt.Println("\n=== Methods (iterate & call) ===")
	calc := Calculator{Value: 10}
	calcType := reflect.TypeOf(calc)
	calcValue := reflect.ValueOf(calc)

	fmt.Printf("  NumMethod: %d\n", calcType.NumMethod())
	for i := 0; i < calcType.NumMethod(); i++ {
		method := calcType.Method(i)
		fmt.Printf("  Method %d: %s (%s)\n", i, method.Name, method.Type)
	}

	// Call Add(5)
	resultAdd := calcValue.MethodByName("Add").Call([]reflect.Value{reflect.ValueOf(5)})
	fmt.Println("  Add(5) =", resultAdd[0].Int())           // 15

	// Call Mul(3)
	resultMul := calcValue.MethodByName("Mul").Call([]reflect.Value{reflect.ValueOf(3)})
	fmt.Println("  Mul(3) =", resultMul[0].Int())           // 30
}

// sample struct with tags
type Sample struct {
	Name string `required:"true" max:"10"`
	Age  string
}

type Person struct {
	Name string
	Age  int
}

type Calculator struct {
	Value int
}

func (c Calculator) Add(n int) int {
	return c.Value + n
}

func (c Calculator) Mul(n int) int {
	return c.Value * n
}
