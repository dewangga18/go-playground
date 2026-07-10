## Go Unit Test

```go
import "testing"
```

Unit testing in Go is built-in — no external framework required. Test files end in `_test.go` and contain functions prefixed with `Test`.

**Run commands:**

| Command | Description |
|---------|-------------|
| `go test` | Run all tests in the current package |
| `go test -v` | Run tests with verbose output (shows each test name) |
| `go test -v -run TestSquare` | Run a specific test function |
| `go test ./...` | Run tests in all packages (recursive) |
| `go test -count=1` | Skip caching — run tests fresh every time |

**Basic test structure:**

```go
package test

import "testing"

func TestSomething(t *testing.T) {
    result := MyFunction(input)
    expected := 42

    if result != expected {
        t.Errorf("Expected %d, got %d", expected, result)
    }
}
```

---

### Failure Options

The `testing.T` type provides 6 methods to signal test failure. They fall into two categories:

| Method | Logs? | Stops? | Behavior |
|--------|-------|--------|----------|
| `t.Fail()` | ❌ | ❌ | Marks test as FAILED, execution **continues** |
| `t.FailNow()` | ❌ | ✅ | Marks test as FAILED, execution **stops** immediately (calls `runtime.Goexit`) |
| `t.Error(args...)` | ✅ | ❌ | Logs a message (`t.Log`), then calls `t.Fail()` — continues |
| `t.Errorf(fmt, args...)` | ✅ | ❌ | Logs a formatted message (`t.Logf`), then calls `t.Fail()` — continues |
| `t.Fatal(args...)` | ✅ | ✅ | Logs a message (`t.Log`), then calls `t.FailNow()` — stops immediately |
| `t.Fatalf(fmt, args...)` | ✅ | ✅ | Logs a formatted message (`t.Logf`), then calls `t.FailNow()` — stops immediately |

**Memory aid:**

```
Error[f] = Log + Fail     → continues  (non-fatal)
Fatal[f] = Log + FailNow  → stops      (fatal)
Fail / FailNow             → silent     (no log output)
```

---

### When to Use Each Method

#### `t.Errorf` — Best for value assertions (most common)

Use when checking **expected vs actual values**. The test continues, so all assertions run — you see every failure at once.

```go
func TestUserAge(t *testing.T) {
    user := GetUser(1)

    // Each assertion runs independently — you see ALL failures in one run
    if user.Name != "Budi" {
        t.Errorf("Name = %q, want %q", user.Name, "Budi")
    }
    if user.Age != 25 {
        t.Errorf("Age = %d, want %d", user.Age, 25)
    }
    if user.Email != "budi@example.com" {
        t.Errorf("Email = %q, want %q", user.Email, "budi@example.com")
    }
}
```

> ✅ **Use `Errorf` when:** you want to report every assertion failure in a single test run. This is the **default choice** for most tests.

---

#### `t.Error` — Same as Errorf, without format string

Use when you have pre-formatted messages or want to pass multiple values without `%` formatting.

```go
if result != expected {
    t.Error("Result mismatch — expected:", expected, "got:", result)
}
```

> ✅ **Use `Error` when:** you don't need format-string features. Think of it as `fmt.Print` vs `fmt.Printf`.

---

#### `t.Fatalf` / `t.Fatal` — Stop on critical failure

Use when **continuing would cause panics or meaningless errors**. Stops the test immediately.

```go
func TestUserSetup(t *testing.T) {
    user, err := CreateTestUser()
    if err != nil {
        t.Fatalf("Failed to create test user: %v", err)
    }

    // These only run if CreateTestUser succeeded
    t.Logf("User created: %+v", user)
    if user.Name != "Expected" {
        t.Errorf("Name = %q", user.Name)
    }
}
```

> ✅ **Use `Fatal`/`Fatalf` when:** (1) Setup failed and subsequent code requires the setup to work, (2) The system state is corrupted and further assertions are meaningless, (3) You'd get a nil pointer panic or a panic from accessing uninitialized data.

---

#### `t.Fail` — Silent failure without stopping

Low-level method — marks the test as failed but produces **no log output**. Rarely used directly.

```go
if len(result) != expectedLen {
    t.Fail() // fails the test silently
    // execution continues — you could check more conditions
}
```

> ✅ **Use `Fail` when:** you're building a custom assertion helper that logs its own message. Not useful in normal test code — prefer `Error`/`Errorf` for the log output.

---

#### `t.FailNow` — Silent failure with immediate stop

Marks the test as failed and **stops immediately** (calls `runtime.Goexit`). No log output. Rarely used directly.

```go
if err != nil {
    t.Helper()        // marks this as a helper function
    t.FailNow()       // stop — caller's test helper will handle logging
}
```

> ✅ **Use `FailNow` when:** you're writing a **custom test helper** that logs its own message but needs to stop execution. In normal test code, prefer `Fatal`/`Fatalf`.

---

### Quick Decision Guide

| Situation | Use |
|-----------|-----|
| Checking expected values — want to see all failures at once | **`Errorf`** or **`Error`** |
| Setup fails — later steps depend on it | **`Fatalf`** or **`Fatal`** |
| Checking multiple independent fields on the same object | **`Errorf`** (all fields reported) |
| Opening a file / connecting to a DB in a test | **`Fatalf`** (can't proceed if this fails) |
| Writing a custom assertion helper | **`Fail`** or **`FailNow`** (helper logs its own message) |
| A single assertion failure makes the rest meaningless | **`Fatalf`** |
| You want to check as many things as possible per test run | **`Errorf`** |
| You want to stop at the first failure (fail-fast) | **`Fatalf`** |

---

### Examples from the Codebase

The file `test/sample_calc_test.go` demonstrates all 6 methods in action:

```go
package test

import "testing"

// TestSquare demonstrates t.Errorf — formatted, non-fatal.
func TestSquare(t *testing.T) {
    result := Square(5)
    expected := 25
    if result != expected {
        t.Errorf("Expected %d, got %d", expected, result)
    }
}

// TestNegativeSquare demonstrates t.Error — non-formatted, non-fatal.
func TestNegativeSquare(t *testing.T) {
    result := Square(-5)
    expected := 25
    if result != expected {
        t.Error("Expected", expected, "got", result)
    }
}

// TestZeroSquare demonstrates t.Fatalf — formatted, fatal (stops).
func TestZeroSquare(t *testing.T) {
    result := Square(0)
    expected := 4 // deliberately wrong — this WILL fail
    if result != expected {
        t.Fatalf("Expected %d, got %d — stopping test now", expected, result)
    }
    t.Log("This line will never print")
}

// TestFatal demonstrates t.Fatal — non-formatted, fatal (stops).
func TestFatal(t *testing.T) {
    if 1 != 2 {
        t.Fatal("1 should equal 2 — cannot proceed, aborting")
    }
    t.Log("This line will never print")
}

// TestFailBehavior shows t.Fail() (continues) vs t.FailNow() (stops).
func TestFailBehavior(t *testing.T) {
    // t.Fail() — marks failed but continues
    if 1 != 2 {
        t.Fail()
    }
    t.Log("After t.Fail() — execution continues")

    // t.FailNow() — marks failed AND stops
    if 2 != 3 {
        t.FailNow()
    }
    t.Log("After t.FailNow() — this will never run")
}
```

Run these with:

```bash
go test -v ./test/
```

---

### Best Practices

#### 1. Default to `Errorf`

Most assertions should use `Errorf`. This way, a single test function reports **every** failure, not just the first one. You fix more in one iteration.

```go
// ✅ Good — three independent checks, all reported
t.Errorf("Name = %q", name)
t.Errorf("Age = %d", age)
t.Errorf("Email = %q", email)
```

#### 2. Use `Fatalf` for setup guards

When a failure means the rest of the test is meaningless, use `Fatalf`. This avoids cascading failures that confuse the real issue.

```go
// ✅ Good — stop immediately if setup fails
db, err := connectDB()
if err != nil {
    t.Fatalf("DB connection failed: %v", err)
}
```

#### 3. Write descriptive messages

Include **what** was expected vs **what** was got. This makes failure output immediately actionable.

```go
// ✅ Good
t.Errorf("Add(2, 3) = %d, want %d", result, 5)

// ❌ Less helpful
if result != 5 {
    t.Errorf("wrong result")
}
```

#### 4. Name tests clearly

Test names describe **what is being tested** and **what the expected outcome is**.

```go
// ✅ Clear
func TestDivideByZero(t *testing.T)
func TestUserRegistration_EmptyEmail(t *testing.T)

// ❌ Vague
func TestDivide(t *testing.T)
func TestUser(t *testing.T)
```

#### 5. Use `t.Helper()` in assertion helpers

If you write a custom assertion function, call `t.Helper()` so error messages point to the **caller's** line number, not the helper.

```go
func assertEqual(t *testing.T, got, want any) {
    t.Helper() // ← makes failure point to the caller's line
    if got != want {
        t.Errorf("got %v, want %v", got, want)
    }
}
```

#### 6. Prefer `Errorf` over `Error` (unless you have a reason)

`Errorf` is more common because format strings make messages clearer. Use `Error` only when you specifically want `fmt.Print`-style argument joining.

#### 7. Rule of thumb

> **`Errorf` everywhere, `Fatalf` only when the test cannot possibly continue safely.**

---

### File-Level Test Options

| Annotation | Effect |
|------------|--------|
| `//go:build !integration` | Build constraint — exclude file from certain builds |
| `func TestMain(m *testing.M)` | Test **main** — runs once for the package. Use for global setup/teardown |
| `t.Skip("reason")` | Skip a test at runtime — useful for conditional tests (e.g. requires network) |

**Example — `TestMain` for global setup:**

```go
func TestMain(m *testing.M) {
    // Global setup
    setupDB()

    // Run all tests in the package
    code := m.Run()

    // Global teardown
    teardownDB()

    os.Exit(code)
}
```

---

### Reference

| File | Purpose |
|------|---------|
| `test/sample_calc_test.go` | All 6 failure methods with inline explanations |
| `test/sample_calc.go` | Simple `Square()` function used in the example tests |
