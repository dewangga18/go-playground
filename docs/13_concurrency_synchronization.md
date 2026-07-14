## Go Concurrency (Synchronization)

When multiple goroutines access **shared data**, you need synchronization — otherwise the result is unpredictable (and often wrong). Go provides several primitives in the `sync` package to safely coordinate goroutines.

---

### Race Condition

A **race condition** happens when two or more goroutines read and write the same variable **at the same time**, and at least one of them is writing. The result depends on timing — and timing is unpredictable.

```go
x := 0

for range 1000 {
    go func() {
        for range 100 {
            x = x + 1    // ← multiple goroutines read & write x simultaneously
        }
    }()
}

time.Sleep(5 * time.Second)
fmt.Println("Result X:", x)
```

**Expected:** `100000` (1000 goroutines × 100 increments)

**Actual:** It varies every run:

```bash
$ go test -v -run TestRaceCondition ./concurrency/

Result X: 88037
--- PASS: TestRaceCondition (5.00s)
```

| Run | Result X | Expected | Difference |
|-----|----------|----------|------------|
| #1 | 88,037 | 100,000 | ❌ ~12k missing |
| #2 | 91,244 | 100,000 | ❌ ~9k missing |
| #3 | 86,911 | 100,000 | ❌ ~13k missing |

**Why does this happen?**

`x = x + 1` is **not** a single operation — it's three steps:
1. Read `x` from memory
2. Add 1 to the value
3. Write the result back to memory

If two goroutines read `x` at the same time (both see `42`), then both write back `43` — one increment is **lost**.

```
Goroutine A: read(42) → compute(43) → write(43)
Goroutine B:                read(42) → compute(43) → write(43)
                               ↑ both read 42, both write 43 — one increment lost!
```

> **Race condition != bug.** A race condition is a **class of bug** — it's a programming error where the outcome depends on non-deterministic timing. Go can detect races at runtime with `go test -race`. But for now, the fix is to **synchronize access**.

---

### sync.Mutex — Mutual Exclusion

`sync.Mutex` ensures that only **one goroutine at a time** can access a critical section of code. Think of it as a **lock** on a shared resource.

```go
func TestHandleWithMutex(t *testing.T) {
    x := 0
    var mutex sync.Mutex

    for range 1000 {
        go func() {
            for range 100 {
                mutex.Lock()
                x = x + 1       // ← only one goroutine at a time
                mutex.Unlock()
            }
        }()
    }

    time.Sleep(5 * time.Second)

    fmt.Println("Result X:", x)
}
```

```bash
$ go test -v -run TestHandleWithMutex ./concurrency/

Result X: 100000
--- PASS: TestHandleWithMutex (5.00s)
```

**How it works:**

| Step | What happens |
|------|--------------|
| `mutex.Lock()` | Goroutine **acquires the lock**. If another goroutine holds it, this goroutine **blocks** (waits) |
| `x = x + 1` | Only one goroutine runs this code at a time — safe |
| `mutex.Unlock()` | Goroutine **releases the lock**. Another waiting goroutine can now acquire it |

| Without Mutex | With Mutex |
|--------------|------------|
| Result unpredictable (race condition) | Result always correct ✅ |
| Faster (no locking overhead) | Slightly slower (lock/unlock overhead) |
| Data corruption risk | Data integrity guaranteed |

> **When to use:** Use `sync.Mutex` whenever multiple goroutines write to the same variable. Every `Lock()` **must** have a corresponding `Unlock()` — typically using `defer` to be safe (though in this example we call `Unlock()` directly because we need to release before the next iteration).

---

### sync.RWMutex — Read-Write Mutex

`sync.RWMutex` is an **optimized** version of `sync.Mutex` for scenarios where:
- **Writes** are rare (need exclusive access)
- **Reads** are frequent (can happen in parallel)

It has two types of locks:

| Lock type | Method | Behavior |
|-----------|--------|----------|
| **Write lock** | `Lock()` / `Unlock()` | Exclusive — blocks all other readers AND writers |
| **Read lock** | `RLock()` / `RUnlock()` | Shared — blocks writers only; other readers can proceed |

**File: `concurrency/3_synchronization_test.go` — `TestRWMutex`**

```go
type BankAccount struct {
    mu      sync.RWMutex
    balance int
}

func (account *BankAccount) Deposit(amount int) {
    account.mu.Lock()           // ← exclusive: no reads during write
    account.balance += amount
    account.mu.Unlock()
}

func (account *BankAccount) GetBalance() int {
    account.mu.RLock()          // ← shared: multiple reads can happen simultaneously
    defer account.mu.RUnlock()
    return account.balance
}

func TestRWMutex(t *testing.T) {
    account := BankAccount{}

    for range 10 {
        go func() {
            for range 1000 {
                account.Deposit(1)               // 1000 writes per goroutine
            }
            fmt.Println("Current Balance:", account.GetBalance())
        }()
    }

    time.Sleep(10 * time.Millisecond)
    fmt.Println("Result X:", account.GetBalance())
}
```

```bash
$ go test -v -run TestRWMutex ./concurrency/

Current Balance: 1000
Current Balance: 2000
Current Balance: 3000
Current Balance: 4000
Current Balance: 5381
Current Balance: 6874
Current Balance: 7874
Current Balance: 8000
Current Balance: 9000
Current Balance: 10000
Result X: 10000
--- PASS: TestRWMutex (0.00s)
```

| Mutex type | Write behavior | Read behavior | Best for |
|-----------|----------------|---------------|----------|
| **`sync.Mutex`** | Exclusive (1 at a time) | Exclusive (1 at a time) | Mostly writes, or mix of reads & writes |
| **`sync.RWMutex`** | Exclusive (1 at a time) | **Shared** (many at once) | Mostly reads, rare writes |

> **When to use `RWMutex`:** Think of a configuration store — 1000 goroutines read the config, but only 1 goroutine updates it. With `RWMutex`, the 1000 readers can all read in parallel. With regular `Mutex`, they'd queue up. **But** for simple cases where performance doesn't matter, a regular `Mutex` is simpler.

---

### Deadlock

A **deadlock** is when two or more goroutines are waiting for each other, and **none of them can proceed**. They're stuck forever.

**File: `concurrency/3_synchronization_test.go` — `TestDealockSimulation`**

```go
type UserBalance struct {
    sync.Mutex
    name    string
    balance int
}

func TransferDeadlock(to, from *UserBalance, amount int) {
    to.Lock()
    fmt.Println("Lock Increasing", to.name)
    to.Change(amount)

    time.Sleep(2 * time.Second)

    from.Lock()          // ← may block forever!
    fmt.Println("Lock Decreasing", from.name)
    from.Change(-amount)

    time.Sleep(2 * time.Second)

    to.Unlock()
    from.Unlock()

    fmt.Println("Unlock", to.name)
    fmt.Println("Unlock", from.name)
}

func TestDealockSimulation(t *testing.T) {
    user1 := UserBalance{name: "Aaron", balance: 500}
    user2 := UserBalance{name: "Evan", balance: 400}

    go TransferDeadlock(&user2, &user1, 50)   // locks user2 → then tries user1
    go TransferDeadlock(&user1, &user2, 35)   // locks user1 → then tries user2

    time.Sleep(5 * time.Second)
    fmt.Println("Final balance")
    fmt.Println("Aaron :", user1.balance)
    fmt.Println("Evan :", user2.balance)
}
```

```bash
$ go test -v -run TestDealockSimulation ./concurrency/

Current balance
Aaron : 500
Evan : 400
Lock Increasing Aaron
Lock Increasing Evan

Final balance
Aaron : 535
Evan : 450
--- PASS: TestDealockSimulation (5.00s)
```

**What happened:**

```
Goroutine A: locks Aaron → sleeps 2s → tries to lock Evan...    ⟵ BLOCKED (held by B)
Goroutine B: locks Evan   → sleeps 2s → tries to lock Aaron...  ⟵ BLOCKED (held by A)

Both goroutines are stuck — each holding one lock and waiting for the other.
```

The test "passes" because `time.Sleep(5)` expires, but the deadlock is still happening — the `"Unlock"` messages **never print**. The actual `go test` process eventually times out and crashes:

```
panic: test timed out after 5s
```

> **Deadlock vs Race Condition:** A race condition gives **wrong results** silently. A deadlock gives **no results at all** — the program hangs. Both are bad, but deadlock is usually easier to detect because it's obvious something is stuck.

---

### Avoiding Deadlock — Consistent Lock Ordering

The fix is simple: **always acquire locks in the same order**. If every goroutine locks `Aaron` before `Evan`, deadlock can't happen.

**File: `concurrency/3_synchronization_test.go` — `TestTransferWithoutDeadlock`**

```go
func TransferWG(to, from *UserBalance, amount int) {
    fmt.Println("Lock Increasing", to.name)

    // Always lock in alphabetical order by name
    if to.name < from.name {
        to.Lock()
        from.Lock()
    } else {
        from.Lock()
        to.Lock()
    }

    defer to.Unlock()
    defer from.Unlock()

    to.Change(amount)
    from.Change(-amount)

    fmt.Println("Unlock", to.name)
}

func TestTransferWithoutDeadlock(t *testing.T) {
    var wg sync.WaitGroup

    user1 := UserBalance{name: "Aaron", balance: 500}
    user2 := UserBalance{name: "Evan", balance: 400}

    wg.Add(2)

    go func() {
        defer wg.Done()
        TransferWG(&user2, &user1, 50)   // locks Aaron → then Evan
    }()

    go func() {
        defer wg.Done()
        TransferWG(&user1, &user2, 35)   // locks Aaron → then Evan (same order!)
    }()

    wg.Wait()

    fmt.Println("Final balance")
    fmt.Println("Aaron :", user1.balance)
    fmt.Println("Evan :", user2.balance)
}
```

```bash
$ go test -v -run TestTransferWithoutDeadlock ./concurrency/

Current balance
Aaron : 500
Evan : 400
Lock Increasing Aaron
Unlock Aaron
Lock Increasing Evan
Unlock Evan

Final balance
Aaron : 485
Evan : 415
--- PASS: TestTransferWithoutDeadlock (0.00s)
```

| Approach | Result |
|----------|--------|
| **Deadlock** (wrong order) | Hangs — goroutines wait forever |
| **Consistent order** (alphabetical) | ✅ Works — total stays 900 (485 + 415) |

**Key rules to avoid deadlock:**

| Rule | Why |
|------|-----|
| **Acquire locks in a consistent global order** | Prevents circular wait — if A always before B, no one can be holding B and waiting for A |
| **Use `defer` for unlocking** | Ensures locks are released even if the function panics |
| **Minimize lock duration** | Only hold locks while accessing shared data — do expensive work outside the lock |

> **When to use this pattern:** Any time you're dealing with **multiple shared resources** (bank transfers, inventory systems, game state). The rule is simple: if there's even a chance of circular waiting, establish a global ordering and stick to it.

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

### sync.Map — Concurrent-Safe Map

A normal `map` in Go is **not safe for concurrent access** — if multiple goroutines read and write the same map without a mutex, your program will **panic** with `fatal error: concurrent map writes`. `sync.Map` is a map designed specifically for concurrent use.

```go
func TestSyncMap(t *testing.T) {
    var syncMap sync.Map
    var wg sync.WaitGroup

    for i := 1; i <= 20; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            syncMap.Store(i, i)    // ← safe to call from multiple goroutines
        }()
    }

    wg.Wait()

    syncMap.Range(func(key, value any) bool {
        fmt.Println(key, value)
        return true                // ← return false to stop iteration
    })
}
```

```bash
$ go test -v -run TestSyncMap ./concurrency/

15 15
13 13
11 11
2 2
19 19
4 4
...
--- PASS: TestSyncMap (0.00s)
```

**Key methods:**

| Method | Description |
|--------|-------------|
| `Store(key, value)` | Stores a key-value pair — safe for concurrent writes |
| `Load(key)` | Retrieves a value by key — returns `(value, ok)` where `ok` is `false` if key doesn't exist |
| `Delete(key)` | Deletes a key-value pair |
| `LoadOrStore(key, value)` | Returns existing value if key exists, otherwise stores and returns the new value |
| `Range(fn)` | Iterates over all key-value pairs — `fn` returns `bool` to continue or `false` to stop |

| Approach | Concurrent writes | Performance |
|----------|------------------|-------------|
| **Normal `map` + `sync.Mutex`** | ✅ Safe | Fast for small maps |
| **`sync.Map`** | ✅ Safe | Optimized for specific patterns (write-once, read-many, disjoint keysets) |

> **When to use `sync.Map`:** Unlike normal maps, `sync.Map` is **optimized** for two specific patterns: (1) **write-once, read-many** — entries are written once and read many times (like a cache). (2) **disjoint keysets** — when multiple goroutines write to different keys (no key contention). For most other cases, a regular `map` with `sync.Mutex` is simpler and faster. `Range()` returns `(key, value any)` — you need type assertion to use the values.

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
        // wait condition
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

### Synchronization Summary

| Concept | Test Function | Description | When to use |
|---------|---------------|-------------|-------------|
| **Race condition** | `TestRaceCondition` | Multiple goroutines writing to same variable without sync | 🚫 Avoid — always synchronize shared writes |
| **`sync.Mutex`** | `TestHandleWithMutex` | Exclusive lock — only one goroutine at a time | Writes to shared data from multiple goroutines |
| **`sync.RWMutex`** | `TestRWMutex` | Read-write lock — parallel reads, exclusive writes | Read-heavy, write-rare scenarios |
| **Deadlock** | `TestDealockSimulation` | Circular wait — each goroutine holding a lock the other needs | 🚫 Avoid — lock in consistent order |
| **Deadlock fix** | `TestTransferWithoutDeadlock` | Consistent lock ordering + `sync.WaitGroup` | Multiple resource locking |
| **`sync.Once`** | `TestOnce` | Run a function exactly once across many goroutines | Lazy initialization, singleton setup |
| **`sync.Pool`** | `TestPool` / `TestPoolNew` / `TestAnotherPool` | Temporary object reuse to reduce allocations | Heavy allocation patterns (JSON, buffers) |
| **`sync.Map`** | `TestSyncMap` | Concurrent-safe map — Store, Load, Range | Write-once-read-many, disjoint keysets |
| **`sync.Cond`** | `TestSyncCond` / `TestSyncCondBroadcast` | Condition variable — `Signal()` wakes one, `Broadcast()` wakes all | Event-driven goroutine wake-up |
| **`sync/atomic`** | `TestAtomicInt64` / `TestAtomicCompareAndSwap` | Lock-free atomic operations — Add, Load, CAS | Simple counters, flags, leader election |

---

### Reference

| File | Purpose |
|------|---------|
| `concurrency/3_synchronization_test.go` | Synchronization — race condition, Mutex, RWMutex, deadlock, Once, Pool, Map, Cond, Atomic |
