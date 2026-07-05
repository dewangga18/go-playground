## Go Error

Go's error handling pattern: return `error` as the last return value, check with `if err != nil`.

### Basic Error

`errors.New()` creates a simple error message.

```go
func divide(num1, div int) (int, error) {
    if div == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return num1 / div, nil
}
```

Caller checks with `if err != nil`:

```go
res, err := divide(10, 0)
if err != nil {
    fmt.Println("ERROR:", err)
} else {
    fmt.Println("RESULT:", res)
}

res2, err2 := divide(10, 2)
if err2 != nil {
    fmt.Println("ERROR:", err2)
} else {
    fmt.Println("RESULT:", res2)
}
```

Output:

```
=== Error Example ===
ERROR: cannot divide by zero
RESULT: 5
```

| Pattern | Description |
|---------|-------------|
| `return 0, errors.New("msg")` | Return error on failure |
| `return result, nil` | Return `nil` error on success |
| `if err != nil` | Check and handle error |

> **Note:** `error` is a built-in interface in Go. `nil` means no error. This pattern is used everywhere in Go — not `try-catch`.

### Custom Error

Create custom error types by implementing the `Error()` method on a struct.

```go
type validationError struct {
    Message string
}

func (ve *validationError) Error() string {
    return ve.Message
}

type notFoundError struct {
    Message string
}

func (ve *notFoundError) Error() string {
    return ve.Message
}
```

Use them in functions:

```go
func saveData(name string, data any) error {
    if name == "" {
        return &validationError{
            Message: "name is required",
        }
    }

    if name != "root" {
        return &notFoundError{
            Message: "username not found",
        }
    }

    fmt.Println("Saving data:", name)
    return nil
}
```

Output:

```
=== Custom Error Example ===
ERROR: name is required
Saving data: root
ERROR: username not found
```

| Error type | Condition | Returns |
|------------|-----------|---------|
| `validationError` | `name == ""` | "name is required" |
| `notFoundError` | `name != "root"` | "username not found" |
| `nil` | `name == "root"` | Success — data saved |

### Error Type Assertion

Use a type switch to handle different error types differently.

```go
switch err := err.(type) {
case *validationError:
    fmt.Println("Validation Error:", err.Message)
case *notFoundError:
    fmt.Println("Not Found Error:", err.Message)
default:
    fmt.Println("Unknown Error:", err)
}
```

Output:

```
Validation Error: name is required
```

| Approach | Use Case |
|----------|----------|
| `errors.New("msg")` | Simple, one-off errors |
| `&validationError{...}` + `Error()` method | Structured errors with extra fields |
| `switch err.(type)` | Handle different error types differently |

> **Note:** Always return **pointer** (`&validationError{}`) when implementing custom errors. Value receivers are also fine, but returning pointer is more consistent with Go conventions.
