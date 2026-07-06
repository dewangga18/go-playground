## Standard Library

Quick reference for Go standard library packages I've encountered so far.

---

### `fmt` — Formatted I/O

```go
import "fmt"
```

**Functions used:**

| Function | Description |
|----------|-------------|
| `Print()` | Print without newline |
| `Println()` | Print with newline |
| `Printf()` | Print with format string |
| `Sprint()` | Returns a string |
| `Sprintf()` | Returns a formatted string |

**Formatting Verbs (used with `Printf`/`Sprintf`):**

| Verb | Description | Example |
|------|-------------|---------|
| `%s` | String | `"hello"` |
| `%d` | Integer (decimal) | `42` |
| `%t` | Boolean | `true` |
| `%f` | Float (default precision) | `123.456000` |
| `%.Nf` | Float with `N` decimal places | `%.2f` → `123.46` |
| `%e` | Float in scientific notation (lowercase `e`) | `1.234568e+02` |
| `%E` | Float in scientific notation (uppercase `E`) | `1.234568E+02` |

**Example:**

```go
fmt.Println("Hello", "World")

name := "John"
age := 20

fmt.Printf("%s is %d years old\n", name, age)

// Formatting verbs
number := 12345.6789
fmt.Printf("%f\n", number)    // 12345.678900
fmt.Printf("%.2f\n", number)  // 12345.68
fmt.Printf("%e\n", number)    // 1.234568e+04
fmt.Printf("%E\n", number)    // 1.234568E+04
```

---

### `strconv` — String Conversions

```go
import "strconv"
```

**Functions used:**

| Function | Description |
|----------|-------------|
| `Itoa()` | Converts `int` to `string` |
| `Atoi()` | Converts `string` to `int` |
| `ParseBool()` | Converts `string` to `bool` |
| `ParseFloat()` | Converts `string` to `float64` |

**Example:**

```go
number := 10
text := strconv.Itoa(number)
fmt.Println(text)
```

**Notes:**
- `Itoa` = **Integer to ASCII**
- `Atoi` = **ASCII to Integer**

---

### `errors` — Error Creation

```go
import "errors"
```

**Functions used:**

| Function | Description |
|----------|-------------|
| `New()` | Creates a new error with a message string |
| `Is()` | Checks if an error matches a target error (supports wrapping) |

**Example:**

```go
return 0, errors.New("cannot divide by zero")
```

**Custom Sentinel Errors:**

```go
var (
    ValidationError = errors.New("validation error")
    NotFoundError   = errors.New("data not found")
)
```

**Error Checking with `switch`:**

```go
func checkErrors(err error) {
    switch err {
    case ValidationError:
        fmt.Println("Validation Error")
    case NotFoundError:
        fmt.Println("Not Found Error")
    case nil:
        fmt.Println("Success")
    default:
        fmt.Println("Unknown Error")
    }
}
```

**Error Checking with `errors.Is()`:**

```go
if errors.Is(err, ValidationError) {
    fmt.Println("Validation Error")
} else if errors.Is(err, NotFoundError) {
    fmt.Println("Not Found Error")
}
```

**Notes:**
- `error` is a built-in interface, not a package. `errors` is the package for creating and working with errors.
- `errors.New()` creates a simple error. For structured errors, implement the `Error()` interface on a custom struct.
- `errors.Is()` is preferred over `==` because it unwraps the error chain and works with wrapped errors.
