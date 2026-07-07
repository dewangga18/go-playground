package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	mainRing()
}

func mainRing() {
	r := ring.New(5)

	for i := 0; i < r.Len(); i++ {
		r.Value = "Value " + strconv.Itoa(i+1)
		r = r.Next()
	}

	// print
	r.Do(func(i any) {
		fmt.Println(i)
	})

	// move
	r = r.Move(2)
	fmt.Println("\nMove 2")
	r.Do(func(i any) {
		fmt.Println(i)
	})

	// link
	fmt.Println("\nLink")
	r2 := ring.New(2)
	r2.Value = "Value 6"
	r2.Next().Value = "Value 7"
	r.Link(r2)
	r.Do(func(i any) {
		fmt.Println(i)
	})

	// unlink
	fmt.Println("\nUnlink")
	r.Unlink(1)
	r.Do(func(i any) {
		fmt.Println(i)
	})
}