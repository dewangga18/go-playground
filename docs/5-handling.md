## Go Error Handling

No `try-catch` in Go. Uses `defer`, `panic`, `recover` instead.

### Defer

Schedules a function call to run **after** the surrounding function completes — whether normal return or panic.

```go
func exampleDefer() {
    defer fmt.Println("end session defer")

    fmt.Println("start session")
}
```

Output:

```
=== Defer Handling ===
start session
end session defer
```

| Behavior | Description |
|----------|-------------|
| **Execution timing** | Runs when surrounding function returns or panics |
| **Multiple defers** | **LIFO** order (last deferred runs first) |
| **Arguments** | Evaluated at time of `defer`, not when it runs |

> **Note:** Commonly used for cleanup — closing files, releasing mutexes, teardown messages.

### Panic

Stops normal execution, runs deferred calls, then propagates upward.

#### Panic and Defer

Deferred functions **always run**, even on panic.

```go
func examplePanicAndDefer(err bool) {
    defer fmt.Println("end session defer")

    if err {
        panic("something went wrong")
    }

    fmt.Println("function end")
}
```

Normal (`err = false`):

```
=== Panic and Defer Handling ===
function end
end session defer
```

Panic (`err = true`):

```
=== Panic and Defer Handling ===
end session defer
panic: something went wrong
```

| Scenario | Defer runs? | After panic? |
|----------|-------------|--------------|
| Normal return | ✅ Yes | — |
| Panic | ✅ Yes (before propagation) | ❌ Program crashes |

### Recover

`recover()` regains control of a panicking goroutine. **Only works inside a deferred function.**

```go
func deferAndRecover() {
    fmt.Println("end session defer")
    msg := recover()
    fmt.Println("recover :", msg)
}

func examplePanicAndRecover() {
    defer deferAndRecover()
    panic("something went wrong")
}
```

Output:

```
=== Panic and Recover Handling ===
end session defer
recover : something went wrong
```

| Function | Purpose |
|----------|---------|
| `defer` | Schedule cleanup that always runs |
| `panic` | Signal unrecoverable error |
| `recover` | Catch panic gracefully |

> **Note:** `recover()` returns the panic value. Returns `nil` if no panic. Only useful inside `defer`.

### When to Use What

| Situation | Approach |
|-----------|----------|
| **Cleanup** (close file, unlock mutex) | `defer` |
| **Fatal error** (can't continue) | `panic` |
| **Prevent crash** at boundary (HTTP handler, goroutine) | `recover` in `defer` |
| **Expected errors** (invalid input, not found) | Return `error` (not covered here) |
