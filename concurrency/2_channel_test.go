package concurrency

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	ch := make(chan string)
	defer close(ch)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "Uhuyyyy"
		fmt.Println("completed channel")
	}()

	result := <-ch
	fmt.Println(result)
}

func GiveMeResponse(ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- "Sample Response"
}

func TestChannelAsParams(t *testing.T) {
	ch := make(chan string)
	defer close(ch)

	go GiveMeResponse(ch)
	result := <-ch
	fmt.Println(result)
}

func OnlyInChannel(ch chan<- string) {
	time.Sleep(1 * time.Second)
	ch <- "Sample Response"
}

func OnlyOutChannel(ch <-chan string) {
	result := <-ch
	fmt.Println(result)
}

func TestInOutChannel(t *testing.T) {
	ch := make(chan string)
	defer close(ch)

	go OnlyInChannel(ch)
	go OnlyOutChannel(ch)

	time.Sleep(2 * time.Second)
}

func TestBufferChannel(t *testing.T) {
	ch := make(chan string, 3)
	defer close(ch)

	time.Sleep(1 * time.Second)
	ch <- "Sample one"
	ch <- "Sample two"
	ch <- "Sample three"
	// ch <- "Sample four" // trigger fail

	fmt.Println("Capacity: ", cap(ch))
	fmt.Println("Length: ", len(ch))

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	// fmt.Println(<-ch) // trigger fail
}

func TestRangeChannel(t *testing.T) {
	ch := make(chan string)
	
	go func() {
		for i := 1; i <= 10; i++ {
			ch <- "Data " + strconv.Itoa(i)
		}
		close(ch)
	}()

	for data := range ch {
		fmt.Println("Data: ", data)
	}
	fmt.Println("Range Done")
}

func TestSelectChannel(t *testing.T) {
	ch1 := make(chan string)
	ch2 := make(chan string)
	channels := [...]chan string{ch1, ch2}
	defer close(ch1)
	defer close(ch2)

	go GiveMeResponse(ch1)
	go GiveMeResponse(ch2)

	for range len(channels) {
		select {
		case data := <-ch1:
			fmt.Println("Data1: ", data)
		case data := <-ch2:
			fmt.Println("Data2: ", data)
		}
	}

	fmt.Println("Select Done")
}
