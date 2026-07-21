package context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func SimulateLongProcess(ctx context.Context) chan int {
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
					time.Sleep(time.Second) // simulate long process
			}
		}
	}()

	return destination
}

func TestTimeoutCounter(t *testing.T) {
	fmt.Println("Total : ", runtime.NumGoroutine()) // 2 -> from local device
	
	// creating context and cancel
	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 3 * time.Second) 
	
	// doesn't affecting counter goroutine to stop
	// cancel() inform system to cancel release internal timer
	defer cancel()

	counter := SimulateLongProcess(ctx)

	// print first 10 numbers
	for n := range counter {
		fmt.Println("Counter : ", n)
		if n == 10 {
			break
		}
	}

	fmt.Println("Total : ", runtime.NumGoroutine()) // 2 -> go routines stop
}