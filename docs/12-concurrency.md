## Go Concurrency (Goroutine)

Concurrency in Go is built around **goroutines** — lightweight execution threads managed by the Go runtime. No external library needed, just the `go` keyword.

---

### Parallel vs Concurrent vs Thread vs Goroutine

Before diving into goroutines, here's a quick orientation on the key concepts:

| Concept | What it means | In practice |
|---------|---------------|-------------|
| **Sequential** | One task at a time, in order | Normal function call — wait for it to finish, then move on |
| **Concurrent** | Handling multiple tasks, switching between them | Like a chef cooking 3 dishes — starts one, while it simmers works on another, then back to the first |
| **Parallel** | Executing multiple tasks **simultaneously** | Like 3 chefs cooking 3 dishes at the exact same time (needs multiple CPU cores) |
| **Thread** | OS-level execution unit — heavy (MBs of memory per thread) | Managed by the operating system. Creating 10,000 threads would crash most machines |
| **Goroutine** | Lightweight execution unit — tiny (~4KB stack, grows as needed) | Managed by Go runtime. Creating 10,000 goroutines is normal and expected |

> **Key insight:** Goroutines are **NOT** OS threads. They're lightweight "co-routines" that run on top of a small pool of OS threads. Go's scheduler handles the mapping — you just write `go fn()`.

```go
// Thread (OS):     heavy, slow to create, OS-scheduled
// Goroutine (Go):  lightweight, fast to create, Go-runtime-scheduled

go myFunction()  // ← run this function concurrently as a goroutine
```

**Analogy — restaurant kitchen:**

| Concept | Kitchen analogy |
|---------|-----------------|
| **OS Thread** | Hiring a new chef every time an order comes in (expensive!) |
| **Goroutine** | Giving an order to an existing chef who can juggle multiple tasks |
| **Concurrent** | Chef starts dish A, while it cooks starts dish B, checks A again |
| **Parallel** | Multiple chefs cooking different dishes at the same time |

---

### Goroutine

A goroutine is started with the `go` keyword followed by a function call:

```go
go functionName()
```

The `go` statement **does not wait** for the function to finish. It launches the function and immediately continues executing the next line.

```go
go HelloWorld()   // ← starts HelloWorld concurrently
fmt.Println("ups") // ← runs immediately, doesn't wait for HelloWorld
```

---

### Example from the Codebase

**File: `concurrency/0_simple_test.go`**

```go
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
	go HelloWorld()            // ← goroutine: runs concurrently
	fmt.Println("ups")         // ← main goroutine prints immediately

	time.Sleep(1 * time.Second) // ← wait so the program doesn't exit early
}
```

**How it works:**

| Line | What happens |
|------|--------------|
| `go HelloWorld()` | Launches `HelloWorld` as a **goroutine** — it starts running but the main goroutine does NOT wait |
| `fmt.Println("ups")` | Main goroutine prints `"ups"` immediately — this usually runs **before** HelloWorld's print |
| `time.Sleep(1 * time.Second)` | Main goroutine pauses for 1 second — gives the HelloWorld goroutine time to finish |

**Output:**

```bash
$ go test -v -run TestHelloWorld ./concurrency/

ups
Hello world
--- PASS: TestHelloWorld (1.00s)
```

> **Note:** `"ups"` printed **first**, then `"Hello world"`. This is the key behavior of goroutines — the `go` call returns immediately, and the goroutine runs concurrently. The order is **not guaranteed** — if HelloWorld ran before the `fmt.Println("ups")`, it could print in reverse order (but in practice, the main goroutine almost always gets there first).

---

### What Happens Without `time.Sleep`?

If you remove `time.Sleep(1 * time.Second)`:

```go
func TestHelloWorld(t *testing.T) {
	go HelloWorld()
	fmt.Println("ups")
	// no sleep — program exits immediately
}
```

The output would likely be **just `"ups"`** (or sometimes both lines, depending on timing).

Why? When `TestHelloWorld` finishes, the test exits — and **all goroutines are killed** when the program exits. The HelloWorld goroutine may not have had time to print before the program stopped.

> **Key rule:** The program **does not wait** for goroutines to finish. If `main()` (or in this case, the test) exits, all goroutines are terminated.

---

### Goroutines Are Lightweight

Goroutines are **not** OS threads. They start with a tiny stack (~4KB) that grows as needed. This means you can launch **thousands** of them without crashing your machine.

**File: `concurrency/1_goroutine_light_test.go`** — demonstrates the difference clearly:

```go
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
```

**Two tests, two behaviors:**

| Test | What it does | Output order | Why |
|------|-------------|--------------|-----|
| `TestWithoutGoroutine` | Calls `DisplayNumber` 19,999 times **sequentially** | 1, 2, 3, 4, 5... 19,999 (always in order) | Each call waits for the previous to finish |
| `TestWithGoroutine` | Launches 19,999 **goroutines** concurrently | 9, 1, 29, 10, 5... (random, changes each run) | All goroutines run at the same time — no guaranteed order |

**Sample output comparison (19,999 iterations):**

```bash
$ go test -v -run TestWithoutGoroutine ./concurrency/
--- PASS: TestWithoutGoroutine (0.53s)
PASS
ok      goplayground/concurrency        1.278s

$ go test -v -run TestWithGoroutine ./concurrency/
PASS
--- PASS: TestWithGoroutine (0.07s)
ok      goplayground/concurrency        0.906s
```

| Test | Test function time | Total package time |
|------|-------------------|-------------------|
| **Without Goroutine** | 0.53s | 1.278s |
| **With Goroutine** | 0.07s | **0.906s** ✅ |

> **Key insight:** With 19,999 iterations, goroutines finish faster (0.906s vs 1.278s). The test function itself runs in just 0.07s because `go` returns immediately — the goroutines run in the background while the test framework handles them. With smaller iteration counts (like 999), the timing difference may not be visible because `fmt.Println` I/O dominates.

**Why this matters:**

| Scenario | With OS threads | With goroutines |
|----------|----------------|-----------------|
| Handle 10,000 web requests | Would exhaust system memory quickly | Normal and expected — Go handles it |
| Start a background task for each user | Expensive — limited to hundreds | Cheap — can handle millions |
| Memory overhead per unit | ~1MB+ per thread | ~4KB per goroutine |
| Creation/destruction speed | Slow (managed by OS kernel) | Fast (managed by Go runtime) |

---

### Key Takeaways

| Concept | Summary |
|---------|---------|
| **Goroutine** | A lightweight concurrent execution — started with `go fn()` |
| **Non-blocking** | `go` returns immediately — doesn't wait for the function to finish |
| **Lightweight** | ~4KB initial stack vs ~1MB+ for OS threads — thousands of goroutines are fine |
| **No guarantees** | Order of execution between goroutines is **not deterministic** |
| **Need synchronization** | Goroutines exit when the program exits — use `WaitGroup`, channels, or `time.Sleep` to coordinate (more on this later) |
| **No special hardware needed** | Goroutines work on any machine — concurrency != parallelism |

> **Next up:** `sync.WaitGroup` and more advanced goroutine synchronization patterns.

---

## Go Channel

Channels (`chan`) are Go's built-in mechanism for communication **between** goroutines. Think of a channel as a pipe — one goroutine sends data in, another receives data out.

| Channel concept | Syntax | What it does |
|----------------|--------|--------------|
| **Unbuffered channel** | `make(chan Type)` | Blocks until both sender & receiver are ready — synchronous handoff |
| **Buffered channel** | `make(chan Type, N)` | Non-blocking until buffer (N) is full — async send |
| **Send** | `ch <- value` | Send data into the channel |
| **Receive** | `<- ch` | Receive data from the channel |
| **Close** | `close(ch)` | Signal that no more data will be sent |
| **Direction** | `chan<-` / `<-chan` | Restrict channel to send-only or receive-only (for function params) |

```go
ch := make(chan string)        // unbuffered channel
ch := make(chan string, 5)      // buffered channel (capacity 5)

ch <- "hello"                    // send
result := <-ch                   // receive

close(ch)                        // close the channel
```

---

### 1. Unbuffered Channel — Basic Send & Receive

**File: `concurrency/2_channel_test.go` — `TestChannel`**

Unbuffered channels are **synchronous**. The sender blocks until a receiver is ready, and the receiver blocks until a sender is ready.

```go
func TestChannel(t *testing.T) {
	ch := make(chan string)
	defer close(ch)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "Uhuyyyy"
		fmt.Println("completed channel")
	}()

	result := <-ch      // ← blocks here for ~2s until data arrives
	fmt.Println(result)
}
```

**How it works:**

| Step | What happens |
|------|--------------|
| `make(chan string)` | Creates an unbuffered channel that carries `string` values |
| `go func() { ... }()` | Launches goroutine that sleeps 2s, then sends `"Uhuyyyy"` |
| `result := <-ch` | Main goroutine **blocks** — waits for data to arrive |
| `ch <- "Uhuyyyy"` | Goroutine sends — immediately unblocks the main goroutine |
| `fmt.Println(result)` | Main goroutine prints `"Uhuyyyy"` |
| `fmt.Println("completed channel")` | Goroutine prints after successful send |

**Output:**

```bash
$ go test -v -run TestChannel ./concurrency/

Uhuyyyy
completed channel
--- PASS: TestChannel (2.00s)
```

> **Note:** `"Uhuyyyy"` prints first because the main goroutine receives and prints **immediately** when data arrives. The sender goroutine prints `"completed channel"` right after the send completes — but that happens after the main goroutine has already woken up and printed.

---

### 2. Channel as Parameter

**File: `concurrency/2_channel_test.go` — `TestChannelAsParams`**

Channels are **reference types** — passing a channel to a function passes the same channel, not a copy.

```go
func GiveMeResponse(ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- "Sample Response"
}

func TestChannelAsParams(t *testing.T) {
	ch := make(chan string)
	defer close(ch)

	go GiveMeResponse(ch)   // ← pass channel to goroutine
	result := <-ch           // ← receive the response
	fmt.Println(result)
}
```

```bash
$ go test -v -run TestChannelAsParams ./concurrency/

Sample Response
--- PASS: TestChannelAsParams (1.00s)
```

> **Key insight:** This is the most common Go pattern — launch a goroutine, give it a channel, and wait for the result. It's how goroutines "return" values.

---

### 3. Channel Direction — Send-only & Receive-only

**File: `concurrency/2_channel_test.go` — `TestInOutChannel`**

You can restrict a channel parameter to **send-only** (`chan<-`) or **receive-only** (`<-chan`). This makes the intent clear and prevents bugs.

```go
// Send-only: can only write to this channel
func OnlyInChannel(ch chan<- string) {
	time.Sleep(1 * time.Second)
	ch <- "Sample Response"
	// x := <-ch   // ← COMPILE ERROR: can't receive from send-only channel
}

// Receive-only: can only read from this channel
func OnlyOutChannel(ch <-chan string) {
	result := <-ch
	fmt.Println(result)
	// ch <- "data"   // ← COMPILE ERROR: can't send to receive-only channel
}

func TestInOutChannel(t *testing.T) {
	ch := make(chan string)
	defer close(ch)

	go OnlyInChannel(ch)   // ← goroutine sends
	go OnlyOutChannel(ch)  // ← goroutine receives

	time.Sleep(2 * time.Second) // ← wait for both to finish
}
```

```bash
$ go test -v -run TestInOutChannel ./concurrency/

Sample Response
--- PASS: TestInOutChannel (2.00s)
```

| Direction | Syntax | Can send? | Can receive? |
|-----------|--------|-----------|--------------|
| **Bidirectional** | `chan string` | ✅ Yes | ✅ Yes |
| **Send-only** | `chan<- string` | ✅ Yes | ❌ No (compile error) |
| **Receive-only** | `<-chan string` | ❌ No (compile error) | ✅ Yes |

> **Why use direction?** It's a contract — `OnlyInChannel` says "I only write to this channel" and `OnlyOutChannel` says "I only read from this channel". The compiler enforces it. This is a best practice in Go.

---

### 4. Buffered Channel

**File: `concurrency/2_channel_test.go` — `TestBufferChannel`**

Buffered channels have a **capacity**. The sender doesn't block until the buffer is full. The receiver doesn't block until the buffer is empty.

```go
func TestBufferChannel(t *testing.T) {
	ch := make(chan string, 3)   // ← buffer capacity = 3
	defer close(ch)

	time.Sleep(1 * time.Second)
	ch <- "Sample one"           // OK: buffer = [one]
	ch <- "Sample two"           // OK: buffer = [one, two]
	ch <- "Sample three"         // OK: buffer = [one, two, three]
	// ch <- "Sample four"        // ← BLOCK: buffer full!

	fmt.Println("Capacity: ", cap(ch))  // 3
	fmt.Println("Length: ", len(ch))    // 3

	fmt.Println(<-ch)  // "Sample one"   → buffer = [two, three]
	fmt.Println(<-ch)  // "Sample two"   → buffer = [three]
	fmt.Println(<-ch)  // "Sample three" → buffer = []
	// fmt.Println(<-ch) // ← BLOCK: buffer empty!
}
```

**Output:**

```bash
$ go test -v -run TestBufferChannel ./concurrency/

Capacity:  3
Length:  3
Sample one
Sample two
Sample three
--- PASS: TestBufferChannel (1.00s)
```

| State | Can send? | Can receive? |
|-------|-----------|--------------|
| **Buffer empty** | ✅ Yes | ❌ Blocks (no data) |
| **Buffer partially filled** | ✅ Yes | ✅ Yes |
| **Buffer full** | ❌ Blocks (no room) | ✅ Yes |

> **When to use buffered channels:** When you want to decouple sender and receiver — the sender can keep working even if the receiver isn't ready yet. But careful: buffered channels can hide synchronization bugs.

---

### 5. Range over Channel

**File: `concurrency/2_channel_test.go` — `TestRangeChannel`**

You can use `for ... range` to receive values from a channel **until it's closed**.

```go
func TestRangeChannel(t *testing.T) {
	ch := make(chan string)

	go func() {
		for i := 1; i <= 10; i++ {
			ch <- "Data " + strconv.Itoa(i)
		}
		close(ch)   // ← MUST close, or range will deadlock
	}()

	for data := range ch {          // ← loops until channel is closed
		fmt.Println("Data: ", data)
	}
	fmt.Println("Range Done")
}
```

**Output:**

```bash
$ go test -v -run TestRangeChannel ./concurrency/

Data:  Data 1
Data:  Data 2
Data:  Data 3
Data:  Data 4
Data:  Data 5
Data:  Data 6
Data:  Data 7
Data:  Data 8
Data:  Data 9
Data:  Data 10
Range Done
--- PASS: TestRangeChannel (0.00s)
```

> **Critical rule:** Always `close(ch)` when the sender is done. Without `close()`, `for ... range ch` blocks **forever** waiting for more data — that's a **deadlock** (and Go will crash with `fatal error: all goroutines are asleep - deadlock!`).

---

### 6. Select Statement — Waiting on Multiple Channels

**File: `concurrency/2_channel_test.go` — `TestSelectChannel`**

`select` lets a goroutine **wait on multiple channel operations** — it picks whichever one is ready first.

```go
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
```

**Output:**

```bash
$ go test -v -run TestSelectChannel ./concurrency/

Data1:  Sample Response
Data2:  Sample Response
Select Done
--- PASS: TestSelectChannel (1.00s)
```

**How `select` works:**

| Scenario | Behavior |
|----------|----------|
| **One channel ready** | `select` executes that case immediately |
| **Multiple channels ready** | Picks one at **random** (fair) |
| **No channels ready** | `select` blocks until one is ready |
| **`default` case** | If no channel is ready, runs `default` immediately (non-blocking) |

> **Note:** `select` is like `switch` but **for channels**. It's the foundation of advanced concurrency patterns like timeouts, non-blocking sends, and fan-in/fan-out.

---

### Channel Summary

| Concept | Test Function | Description |
|---------|---------------|-------------|
| **Unbuffered channel** | `TestChannel` | Synchronous handoff — sender & receiver block until both are ready |
| **Channel as param** | `TestChannelAsParams` | Pass channel to a goroutine function — how goroutines "return" values |
| **Direction** | `TestInOutChannel` | `chan<-` (send-only) vs `<-chan` (receive-only) — compiler-enforced contracts |
| **Buffered channel** | `TestBufferChannel` | `make(chan T, N)` — async send until buffer is full |
| **Range channel** | `TestRangeChannel` | `for data := range ch` — receive until channel is closed |
| **Select** | `TestSelectChannel` | Wait on multiple channels — pick the first one ready |

---

### Reference

| File | Purpose |
|------|---------|
| `concurrency/0_simple_test.go` | Basic goroutine — launch a function with `go` and see concurrent execution |
| `concurrency/1_goroutine_light_test.go` | Goroutines are lightweight — 19,999 goroutines vs sequential loop comparison |
| `concurrency/2_channel_test.go` | Channels — basic, params, direction, buffer, range, and select |
