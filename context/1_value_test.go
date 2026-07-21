package context

import (
	"context"
	"fmt"
	"testing"
)

func TestWithValue(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextC, "e", "E")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)

	// get value with key
	fmt.Println(contextE.Value("e")) // E
	fmt.Println(contextE.Value("d")) // nil -> different parent context
	fmt.Println(contextE.Value("c")) // C -> parent of contextE
	fmt.Println(contextC.Value("e")) // nil -> parent can't access value of child
}