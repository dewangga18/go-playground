## Go Error Handling

Go does not have `try-catch` like other languages. Instead, it uses three built-in mechanisms: **`defer`**, **`panic`**, and **`recover`** — working together to handle exceptional situations.

### Defer

`defer` schedules a function call to run **after** the surrounding function completes — whether it finishes normally or panics.

```go
func exampleDefer() {
    defer fmt.Println("end session defer")

    fmt.Println("start session")
}
```

```
=== Defer Handling ===
start session
end session defer
```

| Behavior | Description |
|----------|-------------|
| **Execution timing** | Runs when the surrounding function returns or panics |
| **Multiple defers** | Executed in **LIFO** order (last deferred runs first) |
| **Arguments evaluated immediately** | Deferred function's arguments are evaluated at the time `defer` is called, not when it runs |

> **Note:** `defer` is commonly used for cleanup operations — closing files, releasing mutexes, or printing teardown messages.

### Panic

`panic` stops the normal execution flow. When a function panics, it stops executing, runs its deferred calls, and propagates the panic upward.

#### Panic and Defer

Deferred functions **always run**, even when a panic occurs.

```go
func examplePanicAndDefer(err bool) {
    defer fmt.Println("end session defer")

    if err {
        panic("something went wrong")
    }

    fmt.Println("function end")
}
```

When `err = false` — normal execution:

```
=== Panic and Defer Handling ===
function end
end session defer
```

When `err = true` — panic occurs:

```
=== Panic and Defer Handling ===
end session defer
panic: something went wrong
```

| Scenario | Defer runs? | After panic? |
|----------|-------------|--------------|
| Normal return | ✅ Yes | — |
| Panic | ✅ Yes (before propagation) | ❌ No — program crashes |

> **Note:** A panic without recovery will crash the program. Use `recover` to handle panics gracefully.

### Recover

`recover()` regains control of a panicking goroutine. It **only works inside a deferred function** — outside defer, it returns `nil`.

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

```
=== Panic and Recover Handling ===
end session defer
recover : something went wrong
```

`recover()` catches the panic value (`"something went wrong"`) and prevents the program from crashing. Execution continues after the function that panicked.

| Function | Purpose |
|----------|---------|
| `defer` | Schedule cleanup code that always runs |
| `panic` | Signal an unrecoverable error |
| `recover` | Catch and handle a panic gracefully |

> **Note:** `recover()` returns the value passed to `panic()`. If there is no panic, `recover()` returns `nil`. Only use recover in `defer` — anywhere else it does nothing.

### When to Use What

| Situation | Approach |
|-----------|----------|
| **Cleanup** (close file, unlock mutex) | `defer` |
| **Fatal error** (can't continue) | `panic` |
| **Prevent crash** at boundary (HTTP handler, goroutine) | `recover` in `defer` |
| **Expected errors** (invalid input, not found) | Return `error` (not covered here — see next chapter) |
