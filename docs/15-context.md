## Go context

The `context` package provides a standard way to carry **deadlines**, **cancellation signals**, and **request-scoped values** across API boundaries and between goroutines. It's widely used in Go servers, HTTP handlers, database calls, and any long-running operation that might need to be cancelled.

**File:** `context/0_context_test.go`

---

### What is context.Context?

`context.Context` is an **interface** with four methods:

```go
type Context interface {
    Deadline() (deadline time.Time, ok bool)
    Done() <-chan struct{}
    Err() error
    Value(key any) any
}
```

| Method | Returns | Purpose |
|--------|---------|---------|
| `Deadline()` | `(time.Time, bool)` | Returns the time when this context will be cancelled (`ok=false` if no deadline) |
| `Done()` | `<-chan struct{}` | A channel that's **closed** when the context is cancelled or times out — used with `select` |
| `Err()` | `error` | Returns `nil` if the context is still active. Returns `Canceled` or `DeadlineExceeded` once cancelled |
| `Value(key)` | `any` | Returns the value associated with the key, or `nil` if not found |

Contexts form a **tree** — you create a root context (no parent), then derive child contexts from it:

```
context.Background()          ← root
├── WithValue(...)            ← add request-scoped data
├── WithCancel(...)           ← can be cancelled manually
├── WithTimeout(...)          ← auto-cancels after a duration
└── WithDeadline(...)         ← auto-cancels at an absolute time
```

Each derived context inherits the parent's deadline, cancellation, and values — and you can add your own.

---

### Two Root Contexts

**File: `context/0_context_test.go` — `TestContext`**

Every context tree starts from one of two root contexts:

```go
func TestContext(t *testing.T) {
    // Background: root context, no deadline, no cancellation
    background := context.Background()
    fmt.Println(background)

    // TODO: placeholder context, use when unsure which context to use
    todo := context.TODO()
    fmt.Println(todo)
}
```

```bash
$ go test -v -run TestContext ./context/

context.Background
context.TODO
--- PASS: TestContext (0.00s)
```

| Root context | When to use |
|-------------|-------------|
| **`context.Background()`** | The **default root** — used in `main()`, initialization, and top-level handlers. Never cancelled, no deadline, no values. |
| **`context.TODO()`** | A **placeholder** — use when you're not sure which context to pass yet (e.g., during refactoring). Same behavior as `Background()`, but signals that a proper context should be provided later. |

> **Rule of thumb:** Always use `context.Background()` as the root. `context.TODO()` is a temporary marker — like a `TODO` comment in code. Both are empty contexts; the difference is **semantic intent**.

---

### WithValue — Passing Data Through Context

**File: `context/1_value_test.go` — `TestWithValue`**

`context.WithValue(parent, key, value)` creates a child context that carries a key-value pair. Values flow **downward** — a child can access its own values and its ancestors' values, but a **parent cannot access a child's values**.

```go
func TestWithValue(t *testing.T) {
    contextA := context.Background()

    contextB := context.WithValue(contextA, "b", "B")
    contextC := context.WithValue(contextA, "c", "C")

    contextD := context.WithValue(contextB, "d", "D")
    contextE := context.WithValue(contextC, "e", "E")

    fmt.Println(contextA)
    fmt.Println(contextB)
    fmt.Println(contextC)
    fmt.Println(contextD)
    fmt.Println(contextE)

    // get value with key
    fmt.Println(contextE.Value("e")) // E
    fmt.Println(contextE.Value("d")) // nil -> different parent context
    fmt.Println(contextE.Value("c")) // C -> parent of contextE
    fmt.Println(contextC.Value("e")) // nil -> parent can't access value of child
}
```

```bash
$ go test -v -run TestWithValue ./context/

context.Background
context.Background.WithValue(b, B)
context.Background.WithValue(c, C)
context.Background.WithValue(b, B).WithValue(d, D)
context.Background.WithValue(c, C).WithValue(e, E)
E
<nil>
C
<nil>
--- PASS: TestWithValue (0.00s)
```

**The context tree looks like this:**

```
contextA (Background)
├── contextB (key: "b" = "B")
│   └── contextD (key: "d" = "D")
└── contextC (key: "c" = "C")
    └── contextE (key: "e" = "E")
```

| Expression | Result | Why |
|------------|--------|-----|
| `contextE.Value("e")` | `"E"` | contextE has its own key `"e"` |
| `contextE.Value("d")` | `<nil>` | contextE is under contextC — not under contextB. Different branch |
| `contextE.Value("c")` | `"C"` | contextE's **parent** (contextC) has key `"c"` — values flow downward |
| `contextC.Value("e")` | `<nil>` | contextC's **child** (contextE) has key `"e"` — parents can't see child values |

> **Key insight:** `WithValue` walks **up** the tree — it checks the current context, then its parent, then its grandparent, until it finds the key or reaches the root. This is why `contextE` can find `"c"` from `contextC`, but `contextC` can't find `"e"` from `contextE`.

---

### WithCancel — Graceful Cancellation

**File: `context/2_cancel_test.go`**

`context.WithCancel(parent)` returns a derived context and a `cancel()` function. When `cancel()` is called, the context's `Done()` channel is **closed**, signalling all goroutines watching it to stop.

---

#### Without Cancel — Goroutine Leak

```go
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
    fmt.Println("Total : ", runtime.NumGoroutine()) // 2

    counter := LeakCounter()

    for n := range counter {
        fmt.Println("Counter : ", n)
        if n == 10 {
            break
        }
    }

    fmt.Println("Total : ", runtime.NumGoroutine()) // 3 -> goroutine still running!
}
```

```bash
$ go test -v -run TestLeakCounter ./context/

Total :  2
Counter :  1
Counter :  2
Counter :  3
Counter :  4
Counter :  5
Counter :  6
Counter :  7
Counter :  8
Counter :  9
Counter :  10
Total :  3
--- PASS: TestLeakCounter (0.00s)
```

| Time | Goroutines | What happened |
|------|-----------|---------------|
| Start | **2** | Test goroutine + GC goroutine |
| After loop | **3** | The counter goroutine is still running — **leaked!** |

When the main goroutine breaks out of the `for range` loop at `n == 10`, the counter goroutine keeps running forever — it has no way to know the consumer has stopped listening.

---

#### With Cancel — Clean Shutdown

```go
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
    fmt.Println("Total : ", runtime.NumGoroutine()) // 2

    parent := context.Background()
    ctx, cancel := context.WithCancel(parent)

    counter := CounterWithContextCancel(ctx)

    for n := range counter {
        fmt.Println("Counter : ", n)
        if n == 10 {
            break
        }
    }
    cancel() // signal: stop the goroutine
    time.Sleep(time.Millisecond) // wait for goroutine to respond

    fmt.Println("Total : ", runtime.NumGoroutine()) // 2 -> back to normal!
}
```

```bash
$ go test -v -run TestCancelCounter ./context/

Total :  2
Counter :  1
Counter :  2
Counter :  3
Counter :  4
Counter :  5
Counter :  6
Counter :  7
Counter :  8
Counter :  9
Counter :  10
Counter goroutine cancelled
Total :  2
--- PASS: TestCancelCounter (0.00s)
```

| Step | What happens |
|------|--------------|
| `ctx, cancel := context.WithCancel(parent)` | Creates a cancellable context — `Done()` channel is open |
| `for n := range counter` | Reads numbers until break at 10 |
| `cancel()` | Closes `ctx.Done()` — the goroutine sees `<-ctx.Done()` in its `select` |
| `"Counter goroutine cancelled"` | Goroutine prints this, then returns and closes `destination` |
| `Total: 2` | Goroutine count back to normal — **no leak** |

| Pattern | Goroutine count after loop | Result |
|---------|---------------------------|--------|
| **No cancel** (TestLeakCounter) | 3 → **leaked** ❌ | Goroutine runs forever |
| **With cancel** (TestCancelCounter) | 2 → **clean** ✅ | Goroutine exits via `ctx.Done()` |

> **The pattern:** The goroutine uses `select` to listen on **both** `ctx.Done()` and the default (send data). When `cancel()` is called, `ctx.Done()` fires first, the goroutine returns, and the channel is closed via `defer close(destination)`. This is the standard Go pattern for graceful goroutine shutdown.

---

### Cancelled Context is Immutable

Once a context is cancelled, **it cannot be revived or reused**. This is a fundamental design principle of the `context` package.

```go
ctx, cancel := context.WithCancel(parent)

cancel()      // ← context is now cancelled
cancel()      // ← safe: second call is a no-op (does nothing)

ctx.Err()     // ← returns context.Canceled (will never change)
```

**What "immutable" means:**

| Aspect | Behavior |
|--------|----------|
| **`ctx.Done()`** | Returns a **closed channel** — `<-ctx.Done()` returns immediately (non-blocking) |
| **`ctx.Err()`** | Returns `context.Canceled` (or `DeadlineExceeded`) — **never returns to `nil`** |
| **Call `cancel()` again** | **No-op** — safe to call multiple times, but has no effect |
| **Reuse the context** | **Not possible** — once cancelled, the context is "dead". Create a **new** derived context from the original parent |

```go
// ❌ WRONG: Trying to reuse a cancelled context
ctx, cancel := context.WithCancel(parent)
cancel()
ctx2, _ := context.WithCancel(ctx)  // ctx2 is ALSO cancelled immediately!
fmt.Println(ctx2.Err())             // context.Canceled

// ✅ CORRECT: Always derive from the original parent
ctx, cancel := context.WithCancel(parent)
cancel()
ctx2, _ := context.WithCancel(parent)  // fresh context from parent
fmt.Println(ctx2.Err())                // <nil> — alive and usable!
```

| Operation | First call | Second call |
|-----------|-----------|-------------|
| `cancel()` | Closes `Done()` channel, sets `Err()` | **No-op** — safe, but does nothing |
| `ctx.Err()` | `context.Canceled` | `context.Canceled` (same value — immutable) |
| `<-ctx.Done()` | Blocks until cancel | Returns immediately (channel is closed) |

> **Why immutable?** Multiple goroutines may be watching the same context. If a cancelled context could be revived, different goroutines would see different states — leading to race conditions. Immutability guarantees that once a goroutine sees `ctx.Done()` closed, it **knows** the context is cancelled and will stay cancelled.

---

### Timeout vs Deadline — What's the Difference?

**File: `context/3_timeout_test.go` — `TestTimeoutCounter`**
**File: `context/4_deadline_test.go` — `TestDeadlineCounter`**

Both `WithTimeout` and `WithDeadline` create contexts that auto-cancel after a certain time. The difference is **how you specify the time**:

| Function | Parameter | Meaning | Example |
|----------|-----------|---------|---------|
| `context.WithTimeout(parent, duration)` | `time.Duration` | "Cancel after **3 seconds from now**" | `WithTimeout(parent, 3*time.Second)` |
| `context.WithDeadline(parent, time.Time)` | `time.Time` | "Cancel **at 14:30:05 exactly**" | `WithDeadline(parent, time.Now().Add(3*time.Second))` |

> **Internal implementation:** `WithTimeout` calls `WithDeadline` internally: `WithDeadline(parent, time.Now().Add(d))`. They are the same mechanism — just different APIs for convenience.

---

#### WithTimeout — Relative Duration

```go
func TestTimeoutCounter(t *testing.T) {
    fmt.Println("Total : ", runtime.NumGoroutine()) // 2

    parent := context.Background()
    ctx, cancel := context.WithTimeout(parent, 3 * time.Second)
    defer cancel() // cancel() releases internal timer resources

    counter := SimulateLongProcess(ctx)

    for n := range counter {
        fmt.Println("Counter : ", n)
        if n == 10 {
            break
        }
    }

    fmt.Println("Total : ", runtime.NumGoroutine()) // 2 -> goroutines stop
}
```

```bash
$ go test -v -run TestTimeoutCounter ./context/

Total :  2
Counter :  1
Counter :  2
Counter :  3
Counter goroutine cancelled
Total :  2
--- PASS: TestTimeoutCounter (3.00s)
```

The counter goroutine processes 3 items (1 per second), then the 3-second timeout fires — `ctx.Done()` closes, and the goroutine exits cleanly.

---

#### WithDeadline — Absolute Time

```go
func TestDeadlineCounter(t *testing.T) {
    fmt.Println("Total : ", runtime.NumGoroutine()) // 2

    parent := context.Background()
    deadline := time.Now().Add(3 * time.Second)
    ctx, cancel := context.WithDeadline(parent, deadline)
    defer cancel()

    counter := SimulateLongProcess(ctx)

    for n := range counter {
        fmt.Println("Counter : ", n)
        if n == 10 {
            break
        }
    }

    fmt.Println("Total : ", runtime.NumGoroutine()) // 2 -> goroutines stop
}
```

```bash
$ go test -v -run TestDeadlineCounter ./context/

Total :  2
Counter :  1
Counter :  2
Counter :  3
Counter goroutine cancelled
Total :  2
--- PASS: TestDeadlineCounter (3.00s)
```

Same behavior — 3 items processed, then deadline reached and context cancelled.

---

#### Timeout vs Deadline Comparison

| Aspect | `WithTimeout` | `WithDeadline` |
|--------|--------------|----------------|
| **Input** | `time.Duration` (e.g. `3*time.Second`) | `time.Time` (e.g. `time.Now().Add(3s)`) |
| **When does it cancel?** | `time.Now() + duration` | At the specified absolute time |
| **Can the deadline be in the past?** | N/A | ✅ Yes — context cancels **immediately** |
| **Internal relationship** | Calls `WithDeadline(parent, time.Now().Add(d))` | The core implementation |
| **When to use** | Simple timeout: "give it 5 seconds" | Specific deadline: "must finish before 3pm" |
| **Can `cancel()` be called early?** | ✅ Yes — `defer cancel()` releases timer | ✅ Yes — same pattern |

> **Key insight:** You can always call `cancel()` **before** the timeout/deadline to cancel early. The `defer cancel()` pattern is important even with WithTimeout/WithDeadline — it releases internal timer resources immediately instead of waiting for the timeout to expire naturally.

---

### The Shared `SimulateLongProcess` Function

Both `TestTimeoutCounter` and `TestDeadlineCounter` use the same helper function, defined in `context/3_timeout_test.go`:

```go
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
```

This is the same pattern as `CounterWithContextCancel` — but with `time.Sleep(time.Second)` added to simulate slow work. This makes the timeout/deadline visible: the goroutine only completes 3 items before the 3-second timeout fires.

---

### Context Summary

| Concept | Function | Description |
|---------|----------|-------------|
| **Root context** | `context.Background()` | Default root — never cancelled, no deadline, no values |
| **Placeholder** | `context.TODO()` | Temporary placeholder — same as Background, signals intent to replace |
| **With value** | `context.WithValue(parent, key, val)` | Attach request-scoped data — child can access parent's values, not vice versa |
| **With cancel** | `context.WithCancel(parent)` | Returns a `cancel()` function — call to signal all watchers to stop |
| **With timeout** | `context.WithTimeout(parent, dur)` | Auto-cancels after a duration — internally calls `WithDeadline` |
| **With deadline** | `context.WithDeadline(parent, time)` | Auto-cancels at an absolute time — can be set in the past for immediate cancel |
| **Check done** | `ctx.Done()` | Returns a closed channel if cancelled — use with `select` |
| **Check error** | `ctx.Err()` | Returns `Canceled` or `DeadlineExceeded` after cancel — `nil` if still active |
| **Cancel is immutable** | — | Once cancelled, context cannot be revived. `cancel()` is no-op on second call |
| **Release resources** | `defer cancel()` | Important even with timeout/deadline — releases internal timer immediately |

---

### Reference

| File | Purpose |
|------|---------|
| `context/0_context_test.go` | Root contexts — `Background()` and `TODO()` |
| `context/1_value_test.go` | WithValue — passing data through context tree, value inheritance |
| `context/2_cancel_test.go` | WithCancel — goroutine leak vs graceful cancellation with cancel() |
| `context/3_timeout_test.go` | WithTimeout — auto-cancel after a relative duration |
| `context/4_deadline_test.go` | WithDeadline — auto-cancel at an absolute time |
