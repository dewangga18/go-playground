package concurrency

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer := time.NewTimer(3 * time.Second)

	fmt.Println("Timer started at ", time.Now())
	time := <-timer.C
	fmt.Println("Timer triggered at ", time)
}

func TestAfter(t *testing.T) {
	ch := time.After(2 * time.Second)

	fmt.Println("After started at ", time.Now())
	after := <-ch

	fmt.Println("After triggered at ", after)
}

func TestAfterFunc(t *testing.T) {
	fmt.Println("Now", time.Now())

	time.AfterFunc(time.Second, func() {
		fmt.Println("AfterFunc triggered at ", time.Now())
	})
}

func TestAfterFuncWithWG(t *testing.T) {
	fmt.Println("Now", time.Now())
	wg := sync.WaitGroup{}
	wg.Add(1)

	time.AfterFunc(time.Second, func() {
		fmt.Println("AfterFunc triggered at ", time.Now())
		wg.Done()
	})
	wg.Wait()
}

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(time.Second)

	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("Stopping ticker")
		ticker.Stop()
	}()

	// return deadlock
	for t := range ticker.C {
		fmt.Println(t)
	}
}

func TestTickerNoDeadlock(t *testing.T) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	done := make(chan struct{})

	go func() {
		time.Sleep(3 * time.Second)
		close(done)
	}()

	for {
		select {
		case tm := <-ticker.C:
			fmt.Println(tm)

		case <-done:
			fmt.Println("Ticker stopped")
			return
		}
	}
}