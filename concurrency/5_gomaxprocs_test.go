package concurrency

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestMaxProcs(t *testing.T) {
	wg := sync.WaitGroup{}
	fmt.Println("Number of goroutines: ", runtime.NumGoroutine())

	for range 100 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Second)
		}()
	}

	totalCores := runtime.NumCPU()
	fmt.Println("Total cores: ", totalCores)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total threads: ", totalThread)

	goroutines := runtime.NumGoroutine()
	fmt.Println("Number of goroutines: ", goroutines)

	wg.Wait()
}

func TestChangeMaxProcs(t *testing.T) {
	wg := sync.WaitGroup{}
	fmt.Println("Number of goroutines: ", runtime.NumGoroutine())

	for range 100 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Second)
		}()
	}

	totalCores := runtime.NumCPU()
	fmt.Println("Total cores: ", totalCores)

	runtime.GOMAXPROCS(20)
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total threads: ", totalThread)

	goroutines := runtime.NumGoroutine()
	fmt.Println("Number of goroutines: ", goroutines)

	wg.Wait()
}
