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

**Example:**

```go
fmt.Println("Hello", "World")

name := "John"
age := 20

fmt.Printf("%s is %d years old\n", name, age)
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

**Example:**

```go
return 0, errors.New("cannot divide by zero")
```

**Notes:**
- `error` is a built-in interface, not a package. `errors` is the package for creating and working with errors.
- `errors.New()` creates a simple error. For structured errors, implement the `Error()` interface on a custom struct.
