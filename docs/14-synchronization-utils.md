## Go Synchronization Utilities

Beyond basic synchronization primitives (Mutex, RWMutex), Go provides utility types in the `sync` and `sync/atomic` packages for specific patterns: one-time initialization, object reuse, condition variables, and lock-free operations.

**File:** `concurrency/4_synchronization_utils_test.go`

---

### sync.Once — Run Exactly Once

`sync.Once` ensures a function is executed **only one time**, no matter how many goroutines call it. Perfect for **lazy initialization**, **config loading**, and **singleton setup**.

```go
func TestOnce(t *testing.T) {
    var once sync.Once
    var group sync.WaitGroup

    for range 100 {
        group.Add(1)
        go func() {
            defer group.Done()
            once.Do(func() {
                fmt.Println("Loading config...")
                time.Sleep(2 * time.Second)
            })
        }()
    }

    group.Wait()
}
```

```bash
$ go test -v -run TestOnce ./concurrency/

Loading config...
--- PASS: TestOnce (2.00s)
```

Despite **100 goroutines** calling `once.Do(...)`, `"Loading config..."` prints **exactly once**.

| Without `sync.Once` | With `sync.Once` |
|---------------------|------------------|
| Load config **100 times** — waste of resources | Load config **once** — efficient |
| Race condition on "is config loaded?" flag | Thread-safe by design |
| Need mutex + boolean flag | Single function call |

> **When to use:** Database connection pool initialization, loading configuration files, setting up loggers, or any "do this one time" setup that multiple goroutines might trigger simultaneously.

---

### sync.Pool — Object Reuse Pool

`sync.Pool` is a **temporary object store** — grab an object from the pool instead of allocating a new one, and put it back when you're done. This reduces **memory allocations** and **GC pressure**.

#### Pool without `New` — May Return `nil`

```go
var pool sync.Pool

pool.Put("Aaron")
pool.Put("Evan")
pool.Put("Juli")

for range 10 {
    go func() {
        data := pool.Get()
        fmt.Println(data)      // ← may print <nil> if pool is empty
        time.Sleep(1 * time.Second)
        pool.Put(data)
    }()
}
```

```bash
$ go test -v -run TestPool ./concurrency/

Aaron
Juli
Evan
<nil>
<nil>
<nil>
<nil>
<nil>
<nil>
<nil>
Done
--- PASS: TestPool (2.00s)
```

Only the first 3 `Get()` calls get actual values — the rest get `<nil>` because there's no `New` function to create objects on demand.

#### Pool with `New` — Auto-create When Empty

```go
pool := sync.Pool{
    New: func() any {
        return "New"            // ← called whenever pool is empty
    },
}

pool.Put("Aaron")
pool.Put("Evan")
pool.Put("Juli")

for range 10 {
    go func() {
        data := pool.Get()
        fmt.Println(data)       // ← never <nil> — returns "New" when empty
        time.Sleep(1 * time.Second)
        pool.Put(data)
    }()
}
```

```bash
$ go test -v -run TestPoolNew ./concurrency/

Aaron
Evan
Juli
New
New
New
New
New
New
New
Done
--- PASS: TestPoolNew (2.00s)
```

No `<nil>` values — when the pool is empty, `Get()` calls the `New` function to create a fresh object.

#### Real-world Example — JSON Parser Pool

```go
type JSONParser struct {
    Data []byte
}

var parserPool = sync.Pool{
    New: func() any {
        fmt.Println("Create parser")
        return &JSONParser{}
    },
}

func Parse(data []byte) {
    parser := parserPool.Get().(*JSONParser)
    parser.Data = data
    fmt.Println(string(parser.Data))
    parser.Data = nil
    parserPool.Put(parser)       // ← return to pool for reuse
}
```

```bash
$ go test -v -run TestAnotherPool ./concurrency/

Create parser
Create parser
User 0
User 9
User 1
User 2
User 3
User 4
User 5
User 6
User 7
User 8
--- PASS: TestAnotherPool (0.00s)
```

Notice `"Create parser"` only printed **twice** for **10 goroutines** — the pool created 2 `JSONParser` objects and reused them across all 10 calls. That's the power of pooling.

| Pool behavior | Without `New` | With `New` |
|---------------|---------------|------------|
| **Pool is empty** | `Get()` returns `nil` — crash risk if you don't check | `Get()` calls `New()` and returns a fresh object |
| **Pool has items** | `Get()` returns an existing object ✅ | `Get()` returns an existing object ✅ |
| **Best for** | Pre-warmed pool (put before get) | Dynamic allocation — objects created on demand |

> **When to use `sync.Pool`:** Heavy allocations — JSON parsing, buffer reuse, template execution. Objects in a pool can be **garbage collected** by Go at any time, so don't use it for long-lived objects or connection pools. Use `sync.Pool` for **temporary, short-lived** objects that are expensive to allocate.

---

### sync.Cond — Condition Variable

`sync.Cond` is a **condition variable** — it lets goroutines **wait** for a signal and **resume** when the condition is met. Think of it as a "wake me up when something happens" mechanism.

There are two ways to wake up waiting goroutines:

| Wake method | Behavior |
|-------------|----------|
| **`Signal()`** | Wakes **one** waiting goroutine (like tapping one person's shoulder) |
| **`Broadcast()`** | Wakes **all** waiting goroutines (like a bell ringing for everyone) |

---

#### sync.Cond with Signal — Wake One at a Time

```go
func TestSyncCond(t *testing.T) {
    cond := sync.NewCond(&sync.Mutex{})
    group := &sync.WaitGroup{}

    for i := 1; i <= 10; i++ {
        group.Add(1)
        go func() {
            cond.L.Lock()
            cond.Wait()                // ← blocks here until signaled
            fmt.Println("Done", i)
            cond.L.Unlock()
            group.Done()
        }()
    }

    go func() {
        for range 10 {
            time.Sleep(10 * time.Millisecond)
            cond.Signal()              // ← wakes ONE goroutine at a time
        }
    }()

    group.Wait()
}
```

```bash
$ go test -v -run TestSyncCond ./concurrency/

Done 1
Done 2
Done 3
Done 4
Done 5
Done 6
Done 7
Done 8
Done 9
Done 10
--- PASS: TestSyncCond (0.10s)
```

10 goroutines are waiting. A signal goroutine sends 10 `Signal()` calls with 10ms intervals — each one wakes **exactly one** goroutine. The goroutines complete in order because `Signal()` picks one at a time.

---

#### sync.Cond with Broadcast — Wake All at Once

```go
func TestSyncCondBroadcast(t *testing.T) {
    cond := sync.NewCond(&sync.Mutex{})
    group := &sync.WaitGroup{}

    for i := 1; i <= 10; i++ {
        group.Add(1)
        go func() {
            cond.L.Lock()
            cond.Wait()                // ← all goroutines wait here
            fmt.Println("Done", i)
            cond.L.Unlock()
            group.Done()
        }()
    }

    go func() {
        time.Sleep(10 * time.Millisecond)
        cond.Broadcast()               // ← wakes ALL goroutines at once!
    }()

    group.Wait()
}
```

```bash
$ go test -v -run TestSyncCondBroadcast ./concurrency/

Done 1
Done 6
Done 2
Done 7
Done 4
Done 5
Done 8
Done 9
Done 3
Done 10
--- PASS: TestSyncCondBroadcast (0.01s)
```

All 10 goroutines wake up **at the same time** — notice the order is random (not sequential like `Signal()`). `Broadcast()` wakes everyone in one shot.

| `Signal()` | `Broadcast()` |
|------------|---------------|
| Wakes **one** goroutine | Wakes **all** goroutines |
| Good for work queues (one job = one worker) | Good for events ("config updated" — all waiters should know) |
| Goroutines complete sequentially | Goroutines wake simultaneously (order depends on scheduler) |

> **How to use `sync.Cond`:**
> 1. Create with `sync.NewCond(&sync.Mutex{})` — Cond wraps a mutex
> 2. Waiters: `cond.L.Lock()` → `cond.Wait()` (blocks here) → do work → `cond.L.Unlock()`
> 3. Signallers: `cond.Signal()` for one, `cond.Broadcast()` for all
> 4. **Important:** `cond.Wait()` **automatically unlocks** the mutex while waiting and **re-locks** before returning — this is how other goroutines can acquire the lock to signal.

> **Without `Signal()` or `Broadcast()`, all goroutines block forever** — it's a silent deadlock. Always make sure a signal will be sent.

---

### sync/atomic — Atomic Operations

`sync/atomic` provides **lock-free** atomic operations for basic types. Unlike `sync.Mutex` which uses OS-level locking, atomic operations use **CPU-level instructions** — they're lightweight and fast.

**Package:**

```go
import "sync/atomic"
```

| Function | Description |
|----------|-------------|
| `atomic.AddInt64(ptr, delta)` | Atomically adds `delta` to `*ptr` — returns the new value |
| `atomic.LoadInt64(ptr)` | Atomically reads the value of `*ptr` — safe even while other goroutines write |
| `atomic.StoreInt64(ptr, val)` | Atomically writes `val` to `*ptr` — safe even while other goroutines read |
| `atomic.CompareAndSwapInt64(ptr, old, new)` | If `*ptr == old`, set to `new` and return `true`. Otherwise return `false` (atomic check-and-set) |
| `atomic.Int32` / `atomic.Int64` | **Typed wrappers** (Go 1.19+) — `.Add()`, `.Load()`, `.Store()`, `.CompareAndSwap()` as methods |

---

#### atomic.AddInt64 & atomic.LoadInt64 — Lock-Free Counter

```go
func TestAtomicInt64(t *testing.T) {
    var wg sync.WaitGroup
    var counter int64 = 0

    for range 100 {
        wg.Add(1)
        go func() {
            for j := 1; j <= 33; j++ {
                atomic.AddInt64(&counter, 1)    // ← lock-free atomic increment
            }
            wg.Done()
        }()
    }

    wg.Wait()

    fmt.Println("Counter Load", atomic.LoadInt64(&counter))  // safe read
    fmt.Println("Counter", counter)                           // also works (but not atomic)
}
```

```bash
$ go test -v -run TestAtomicInt64 ./concurrency/

Counter Load 3300
Counter 3300
--- PASS: TestAtomicInt64 (0.00s)
```

100 goroutines × 33 increments = 3300. **Always correct** — no race condition, no mutex needed.

| Approach | Code | Speed | Safety |
|----------|------|-------|--------|
| `x = x + 1` | Simple | ✅ Fast | ❌ Race condition |
| `mutex.Lock(); x++; mutex.Unlock()` | Medium | ❌ Slower (OS lock) | ✅ Safe |
| `atomic.AddInt64(&x, 1)` | Shortest | ✅ Fast (CPU-level) | ✅ Safe |

---

#### atomic.CompareAndSwap — Check-and-Set (CAS)

`CompareAndSwap` is the foundation of **lock-free algorithms** — it atomically checks a value and swaps it if it matches the expected value.

```go
func TestAtomicCompareAndSwap(t *testing.T) {
    var running atomic.Int32    // ← typed wrapper (Go 1.19+)
    var wg sync.WaitGroup

    for range 10 {
        wg.Add(1)
        go func() {
            defer wg.Done()

            if running.CompareAndSwap(0, 1) {    // ← only ONE goroutine succeeds
                fmt.Println("Server started")
                time.Sleep(time.Second)
                running.Store(0)                 // ← reset for next time
                fmt.Println("Server stopped")
            } else {
                fmt.Println("Already running")   // ← the other 9 goroutines see this
            }
        }()
    }

    wg.Wait()
    fmt.Println("Running:", running.Load())
}
```

```bash
$ go test -v -run TestAtomicCompareAndSwap ./concurrency/

Server started
Already running
Already running
Already running
Already running
Already running
Already running
Already running
Already running
Already running
Server stopped
Running: 0
--- PASS: TestAtomicCompareAndSwap (1.00s)
```

**How it works:**

```
CAS(0 → 1): checks if running == 0, if so set to 1

Goroutine A: CAS(0, 1) → succeeds (running was 0, now 1) → "Server started"
Goroutine B: CAS(0, 1) → fails (running is 1, not 0) → "Already running"
Goroutine C: CAS(0, 1) → fails → "Already running"
...
```

| `sync.Mutex` | `atomic.CompareAndSwap` |
|-------------|------------------------|
| Heavyweight (OS lock) | Lightweight (CPU instruction) |
| Blocks goroutine (context switch) | Spin-retry pattern (no context switch) |
| Good for long critical sections | Good for quick check-and-set |

> **When to use `sync/atomic`:** Simple counters, flags, and status values. CAS is perfect for leader election (start a server once), throttling, and one-shot initialization. For complex data structures or long operations, stick with `sync.Mutex`.

---

### Summary

| Concept | Test Function | Description | When to use |
|---------|---------------|-------------|-------------|
| **`sync.Once`** | `TestOnce` | Run a function exactly once across many goroutines | Lazy initialization, singleton setup |
| **`sync.Pool`** | `TestPool` / `TestPoolNew` / `TestAnotherPool` | Temporary object reuse to reduce allocations | Heavy allocation patterns (JSON, buffers) |
| **`sync.Cond`** | `TestSyncCond` / `TestSyncCondBroadcast` | Condition variable — `Signal()` wakes one, `Broadcast()` wakes all | Event-driven goroutine wake-up |
| **`sync/atomic`** | `TestAtomicInt64` / `TestAtomicCompareAndSwap` | Lock-free atomic operations — Add, Load, CAS | Simple counters, flags, leader election |

---

### Reference

| File | Purpose |
|------|---------|
| `concurrency/4_synchronization_utils_test.go` | Synchronization utilities — Once, Pool, Cond, Atomic |
