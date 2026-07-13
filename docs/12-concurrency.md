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

**File: `concurrency/0-simple.go`**

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

> **What we'll learn next:** Channels (`chan`), `sync.WaitGroup`, `select` statement, and proper goroutine synchronization patterns.

---

### Reference

| File | Purpose |
|------|---------|
| `concurrency/0-simple.go` | Basic goroutine — launch a function with `go` and see concurrent execution |
| `concurrency/1_goroutine_light_test.go` | Goroutines are lightweight — 19,999 goroutines vs sequential loop comparison |
