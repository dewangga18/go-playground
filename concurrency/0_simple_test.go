package concurrency

import (
	"fmt"
	"testing"
	"time"
)

func HelloWorld() {
	fmt.Println("Hello world")
}

func TestHelloWorld(t *testing.T) {
	go HelloWorld()
	fmt.Println("ups")

	time.Sleep(time.Second)
}
