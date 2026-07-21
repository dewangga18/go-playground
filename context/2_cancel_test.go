package context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func LeakCounter() chan int {
	destination := make(chan int)

	go func ()  {
		defer close(destination)

		counter := 1

		for {
			destination <- counter
			counter++
		}
	}()

	return destination
}

func TestLeakCounter(t *testing.T) {
	fmt.Println("Total : ", runtime.NumGoroutine()) // 2 -> from local device
	
	counter := LeakCounter()

	// print first 10 numbers
	for n := range counter {
		fmt.Println("Counter : ", n)
		if n == 10 {
			break
		}
	}

	fmt.Println("Total : ", runtime.NumGoroutine()) // 3 -> go routines still running
}

func CounterWithContextCancel(ctx context.Context) chan int {
	destination := make(chan int)

	go func ()  {
		defer close(destination)

		counter := 1

		for {
			select {
				case <- ctx.Done():
					fmt.Println("Counter goroutine cancelled")
					return
				default:
					destination <- counter
					counter++
			}
		}
	}()

	return destination
}

func TestCancelCounter(t *testing.T) {
	fmt.Println("Total : ", runtime.NumGoroutine()) // 2 -> from local device
	
	// creating context and cancel
	parent := context.Background()
	ctx, cancel := context.WithCancel(parent) // ctx is child context, cancel is function to cancel the context

	counter := CounterWithContextCancel(ctx)

	// print first 10 numbers
	for n := range counter {
		fmt.Println("Counter : ", n)
		if n == 10 {
			break
		}
	}
	cancel() // calling cancel to stop counter goroutine
	time.Sleep(time.Millisecond) // waiting for counter goroutine to stop

	fmt.Println("Total : ", runtime.NumGoroutine()) // 2 -> go routines stop
}
	