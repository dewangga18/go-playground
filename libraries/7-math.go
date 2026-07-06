package main

import (
	"fmt"
	"math"
	"math/rand/v2"
)

func main() {
	mainMath()
}

func mainMath() {
	basicOperation()
	randomFunction()
}

func basicOperation() {
	fmt.Println(math.Abs(-10.5))  // 10.5
	fmt.Println(math.Max(10, 20)) // 20
	fmt.Println(math.Min(10, 20)) // 10
	fmt.Println(math.Round(3.6))  // 4
	fmt.Println(math.Ceil(3.2))   // 4
	fmt.Println(math.Floor(3.8))  // 3
	fmt.Println(math.Pow(2, 3))   // 8
	fmt.Println(math.Sqrt(9))     // 3

	// like % but float only can use math
	fmt.Println(math.Mod(5, 2))   // 1
	fmt.Println(math.Mod(5.5, 2)) // 1.5

	fmt.Println(math.Pi)
}

func randomFunction() {
	fmt.Println("Random Int:", rand.Int())
	fmt.Println("Random IntN (0-9):", rand.IntN(10))
	fmt.Println("Random Float:", rand.Float64())

	min := 10.0
	max := 20.0
	fmt.Println("Random FloatN:", rand.Float64()*(max-min))
}
