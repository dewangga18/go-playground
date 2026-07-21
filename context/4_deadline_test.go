package context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestDeadlineCounter(t *testing.T) {
	fmt.Println("Total : ", runtime.NumGoroutine()) // 2 -> from local device

	// creating context and cancel
	parent := context.Background()
	deadline := time.Now().Add(3 * time.Second)
	ctx, cancel := context.WithDeadline(parent, deadline)

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
