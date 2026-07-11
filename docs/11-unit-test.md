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
| `func TestMain(m *testing.M)` | Test **main** — runs once for the **entire package**. Use for global setup/teardown |
| `t.Skip("reason")` | Skip a test at runtime — useful for conditional tests (e.g. requires network) |

---

#### `TestMain` — Global Setup/Teardown

`TestMain` runs **once** for all tests in a package — before any test, and after all tests finish.

```go
// File: test/sample_main_test.go
package test

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("Starting Test")   // ← runs ONCE before all tests
	m.Run()                          // ← runs all test functions
	fmt.Println("Test Finished")    // ← runs ONCE after all tests
}
```

**Without `TestMain`**, each test runs independently and there's no hook for global setup.

**With `TestMain`**, you can:

| Use case | Example |
|----------|---------|
| Open DB connection | `setupDB()` before `m.Run()` |
| Seed test data | `seedData()` before `m.Run()` |
| Start mock server | `testServer.Start()` before `m.Run()` |
| Clean up resources | `testServer.Close()` after `m.Run()` |

Output when running all tests:

```
Starting Test
=== RUN   TestSquare
--- PASS: TestSquare (0.00s)
=== RUN   TestFuncAssertion
--- PASS: TestFuncAssertion (0.00s)
... (all tests run here) ...
Test Finished
```

> **Note:** `m.Run()` returns an exit code (`int`). In real projects, you should pass it to `os.Exit(code)` so the shell reports the correct pass/fail status. The example above omits it for simplicity.

---

#### `t.Skip` — Conditional Skipping

Skip a test at runtime when a condition isn't met.

```go
// File: test/sample_skip_test.go
package test

import (
	"runtime"
	"testing"
)

func TestSkipFunction(t *testing.T) {
	goos := runtime.GOOS
	if goos != "linux" {
		t.Skip("Skipping this test because it's not implemented for", goos)
	}

	// This code only runs on Linux
	t.Log("This will not be printed on macOS or Windows")
}
```

Output on macOS:

```
=== RUN   TestSkipFunction
    sample_skip_test.go:11: Skipping this test because it's not implemented for darwin
--- SKIP: TestSkipFunction (0.00s)
```

> `t.Skip` calls `t.SkipNow()` internally (stops the test immediately). Use it for: OS-specific tests, network-dependent tests, or tests that need external services that aren't available.

---

### Subtests (`t.Run`)

Subtests let you group related test cases inside one test function. Each subtest runs independently — setup/teardown can be shared, but one subtest failure doesn't stop others.

```go
import "testing"
```

**Basic subtest structure:**

```go
t.Run("subtest-name", func(t *testing.T) {
    // test logic here
})
```

**Example — `test/sample_sub_test.go`:**

```go
package test

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSubFunction(t *testing.T) {
	t.Run("sub1", func(t *testing.T) {
		res := Square(2)
		assert.Equal(t, 4, res)
	})

	t.Run("sub2", func(t *testing.T) {
		res := Square(3)
		assert.Equal(t, 9, res)
	})
}
```

Output:

```
=== RUN   TestSubFunction
=== RUN   TestSubFunction/sub1
=== RUN   TestSubFunction/sub2
--- PASS: TestSubFunction (0.00s)
    --- PASS: TestSubFunction/sub1 (0.00s)
    --- PASS: TestSubFunction/sub2 (0.00s)
```

**Benefits of subtests:**

| Benefit | Explanation |
|---------|-------------|
| **Shared setup** | Common variables/code before `t.Run()` — runs once, not per subtest |
| **Independent failures** | One subtest can fail, others still run |
| **Clear naming** | Output shows `TestName/subname` — easy to identify failures |
| **Run a single subtest** | `go test -v -run TestSubFunction/sub1` |
| **Table-driven + subtests** | Combine table tests with `t.Run(name, fn)` for the cleanest pattern |

> Subtests are especially powerful with **table-driven tests**: define a slice of test cases, then loop and call `t.Run()` for each case. Each row in the table becomes a named subtest.

---

### Table-Driven Tests

Table-driven tests are the idiomatic Go way to test multiple inputs/outputs with minimal repetition. Define a **table** (slice of test cases), then loop and run each case as a subtest.

**Pattern:**

```go
func TestSomething(t *testing.T) {
    tests := []struct {
        name    string   // subtest name
        input   int
        want    int
    }{
        {name: "case 1", input: 2, want: 4},
        {name: "case 2", input: 3, want: 9},
    }

    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            got := MyFunction(tc.input)
            assert.Equal(t, tc.want, got)
        })
    }
}
```

**Example — `test/sample_table_test.go`:**

```go
package test

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestTableFunction(t *testing.T) {
	squareTests := []struct {
		name    string
		expect  int
		request int
	}{
		{name: "Test1", expect: 1,  request: 1},
		{name: "Test2", expect: 4,  request: 2},
		{name: "Test3", expect: 9,  request: 3},
		{name: "Test4", expect: 16, request: 4},
		{name: "Test5", expect: 25, request: 5},
	}

	for _, test := range squareTests {
		t.Run(test.name, func(t *testing.T) {
			got := Square(test.request)
			assert.Equal(t, test.expect, got)
		})
	}
}
```

Output:

```
=== RUN   TestTableFunction
=== RUN   TestTableFunction/Test1
=== RUN   TestTableFunction/Test2
...
--- PASS: TestTableFunction (0.00s)
    --- PASS: TestTableFunction/Test1 (0.00s)
    --- PASS: TestTableFunction/Test2 (0.00s)
    --- PASS: TestTableFunction/Test3 (0.00s)
    --- PASS: TestTableFunction/Test4 (0.00s)
    --- PASS: TestTableFunction/Test5 (0.00s)
```

**Run a single case:**

```bash
go test -v -run TestTableFunction/Test3
```

| Benefit | Why |
|---------|-----|
| **Add cases easily** | Just add a new row to the struct slice |
| **No copy-paste** | One test function, one loop — handles all cases |
| **Clear failure output** | Test name tells you exactly which case failed |
| **Combines with testify** | `assert.Equal` makes assertions one-liners |

> Table-driven tests are the **standard pattern** in Go. Even the standard library uses them extensively. Practice this pattern until it becomes second nature.

---

### Benchmark — Performance Testing

Benchmarks measure how fast and how memory-efficient your code is. They use `testing.B` instead of `testing.T`.

**Run commands:**

| Command | Description |
|---------|-------------|
| `go test -bench=.` | Run all benchmarks |
| `go test -bench=. -benchmem` | Run benchmarks + show memory allocations (`B/op`, `allocs/op`) |
| `go test -bench=. -run=^$` | Run benchmarks **only** (skip all tests) |
| `go test -bench=BenchmarkSquare` | Run specific benchmark |

**Basic benchmark — `test/sample_brenchmark_test.go`:**

```go
package test

import "testing"

func BenchmarkSquare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Square(10)
	}
}
```

`b.N` is adjusted by Go runtime — starts small, then increases until the benchmark gets a stable measurement.

**Benchmark with subtests:**

```go
func BenchmarkSquareSub(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.Run("Square 5", func(b *testing.B) {
			Square(5)
		})
		b.Run("Square 3", func(b *testing.B) {
			Square(3)
		})
	}
}
```

**Table-driven benchmark:**

```go
func BenchmarkTableSquare(b *testing.B) {
	testCase := []struct {
		Name  string
		Input int
	}{
		{Name: "Square 5", Input: 5},
		{Name: "Square 3", Input: 3},
		{Name: "Square 2", Input: 2},
		{Name: "Square 1", Input: 1},
	}

	for _, tc := range testCase {
		b.Run(tc.Name, func(b *testing.B) {
			Square(tc.Input)
		})
	}
}
```

**Sample output:**

```
BenchmarkSquare-8                    1000000000               0.3269 ns/op          0 B/op          0 allocs/op
BenchmarkSquareSub/Square_5-8        1000000000               0.0000001 ns/op       0 B/op          0 allocs/op
BenchmarkSquareSub/Square_3-8        1000000000               0.0000000 ns/op       0 B/op          0 allocs/op
BenchmarkTableSquare/Square_5-8      1000000000               0.0000001 ns/op       0 B/op          0 allocs/op
BenchmarkTableSquare/Square_3-8      1000000000               0.0000001 ns/op       0 B/op          0 allocs/op
BenchmarkTableSquare/Square_2-8      1000000000               0.0000000 ns/op       0 B/op          0 allocs/op
BenchmarkTableSquare/Square_1-8      1000000000               0.0000001 ns/op       0 B/op          0 allocs/op
```

**What the columns mean:**

| Column | Meaning | Example |
|--------|---------|---------|
| `BenchmarkSquare-8` | Benchmark name + number of CPU cores | `-8` = 8 cores |
| `1000000000` | Number of iterations (`b.N`) | 1 billion iterations |
| `0.3269 ns/op` | Nanoseconds **per operation** (lower = faster) | 0.3 nanoseconds — very fast! |
| `0 B/op` | Bytes allocated **per operation** (lower = better) | 0 — no allocations |
| `0 allocs/op` | Number of memory allocations per operation | 0 — no heap allocs |

> **Benchmark rules of thumb:** `go test -bench=. -benchmem -run=^$` is the standard command — benchmarks only, with memory stats. If `ns/op` is suspiciously low (like `0.0000001`), the compiler may have optimized away the function call — use `result := Square(n); _ = result` to prevent that.

---

### Reference

| File | Purpose |
|------|---------|
| `test/sample_calc_test.go` | All 6 failure methods with inline explanations |
| `test/sample_calc.go` | Simple `Square()` function used in the example tests |
| `test/sample_testify_test.go` | Testify `assert` + `require` + mock demonstrations |
| `test/sample_main_test.go` | `TestMain` — global setup/teardown hook |
| `test/sample_skip_test.go` | `t.Skip` — conditional test skipping |
| `test/sample_sub_test.go` | Subtests with `t.Run` — nested test groups |
| `test/sample_table_test.go` | Table-driven test with testify + subtests |
| `test/sample_brenchmark_test.go` | Benchmark — performance testing with `testing.B` |

---

### Testify — Rich Assertions

Testify (`github.com/stretchr/testify`) is a Go testing toolkit that provides expressive assertion functions, reducing boilerplate compared to raw `if` + `t.Errorf`.

```go
import "github.com/stretchr/testify/assert"
import "github.com/stretchr/testify/require"
```

**`assert` vs `require` — same API, different failure behavior:**

| Package | On failure | Analogy |
|---------|------------|---------|
| `assert` | Logs + continues test | Like `t.Error` / `t.Errorf` |
| `require` | Logs + stops test immediately | Like `t.Fatal` / `t.Fatalf` |

**Common assertions (available in both `assert` and `require`):**

| Function | What it checks |
|----------|----------------|
| `Equal(t, want, got)` | Values are deeply equal |
| `NotEqual(t, a, b)` | Values differ |
| `Nil(t, val)` | Value is nil |
| `NotNil(t, val)` | Value is not nil |
| `Error(t, err)` | `err != nil` |
| `NoError(t, err)` | `err == nil` |
| `True(t, val)` | `val == true` |
| `False(t, val)` | `val == false` |
| `Empty(t, val)` | Length is 0 or zero-value |
| `NotEmpty(t, val)` | Length > 0 |
| `Contains(t, s, sub)` | Contains substring/item |
| `Len(t, obj, n)` | `len(obj) == n` |
| `IsType(t, typ, val)` | Value is expected type |
| `Zero(t, val)` | Value is zero-value (0, "", nil, false) |
| `NotZero(t, val)` | Value is not zero-value |

**Example — `test/sample_testify_test.go`:**

```go
func TestFuncAssertion(t *testing.T) {
	assert.Equal(t, 123, 123)
	assert.NotEqual(t, 123, 456)
	assert.Nil(t, nil)
	assert.NotNil(t, 1)

	err := errors.New("something went wrong")
	assert.Error(t, err)
	assert.NoError(t, nil)

	assert.True(t, true)
	assert.False(t, false)

	users := [1]string{}
	assert.Empty(t, users)

	users[0] = "Aaron"
	assert.NotEmpty(t, users)

	assert.Contains(t, "Aaron", "ar")
	assert.NotContains(t, "Aaron", "zz")
	assert.Len(t, users, 1)
	assert.IsType(t, [1]string{}, users)

	assert.Zero(t, 0)
	assert.Zero(t, "")
	assert.NotZero(t, 1)
}
```

**Recommendations:**

| Situation | Use |
|-----------|-----|
| Default choice | `assert` — continue on failure, see all failures |
| Setup guard / critical precondition | `require` — stop immediately |
| You want standard library only | `t.Error` / `t.Fatalf` — no external deps needed |

> Testify is **optional** — Go's built-in `testing` package is already powerful. Use testify when you want cleaner code (>15 assertions in one file), especially for table-driven tests with `assert.Equal` instead of `if got != want { t.Errorf(...) }`.

---

#### Mocking with Testify (`testify/mock`)

Testify also provides a mocking framework via `github.com/stretchr/testify/mock`. It lets you create mock objects that simulate external dependencies.

```go
import "github.com/stretchr/testify/mock"
```

**Pattern — interface → mock struct → test:**

```go
// 1. Define the interface
// File: test/sample_testify_test.go
type UserFetcher interface {
	FetchUser(id int) (string, error)
}

// 2. Function that depends on the interface
func GetUserGreeting(fetcher UserFetcher, id int) string {
	name, err := fetcher.FetchUser(id)
	if err != nil {
		return "Hello, Guest!"
	}
	return "Hello, " + name + "!"
}

// 3. Mock struct that implements the interface
type MockUserFetcher struct {
	mock.Mock
}

func (m *MockUserFetcher) FetchUser(id int) (string, error) {
	args := m.Called(id)
	return args.String(0), args.Error(1)
}
```

**4. Using the mock in a test:**

```go
func TestGetUserGreeting(t *testing.T) {
	// success case — user found
	mockFetcher := new(MockUserFetcher)
	mockFetcher.On("FetchUser", 1).Return("Aaron", nil)

	result := GetUserGreeting(mockFetcher, 1)
	assert.Equal(t, "Hello, Aaron!", result)
	mockFetcher.AssertExpectations(t)

	// error case — user not found
	mockFetcher2 := new(MockUserFetcher)
	mockFetcher2.On("FetchUser", 999).Return("", errors.New("not found"))

	result2 := GetUserGreeting(mockFetcher2, 999)
	assert.Equal(t, "Hello, Guest!", result2)
	mockFetcher2.AssertExpectations(t)
}
```

**Key mock methods:**

| Method | Description |
|--------|-------------|
| `mockObj.On(method, args...).Return(values...)` | Set up expected call + return values |
| `mockObj.AssertExpectations(t)` | Verify all expected calls were actually made |
| `mockObj.AssertCalled(t, method, args...)` | Assert a specific method was called with specific args |
| `args.Get(0).(string)` | Get return value by index with type assertion |
| `args.String(0)` | Shorthand for string return value |
| `args.Error(1)` | Shorthand for error return value |

> **Mocking rule of thumb:** Mock **interfaces**, not implementations. Your function should accept an interface so you can swap real implementations with mocks in tests. This makes tests fast, isolated, and reliable.
