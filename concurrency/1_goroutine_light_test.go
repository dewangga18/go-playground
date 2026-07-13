package concurrency

import (
	"fmt"
	"testing"
)

func DisplayNumber(number int) {
	fmt.Println("Number: ", number)
}

func TestWithoutGoroutine(t *testing.T) {
	for i := 1; i < 20000; i++ {
		DisplayNumber(i)
	}
}

func TestWithGoroutine(t *testing.T) {
	for i := 1; i < 20000; i++ {
		go DisplayNumber(i)
	}
}
