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
| `Atoi()` | Converts `string` to `int` — returns `(int, error)` |
| `Itoa()` | Converts `int` to `string` |
| `ParseBool()` | Converts `string` to `bool` — returns `(bool, error)`. Accepts `"1"`, `"t"`, `"T"`, `"TRUE"`, `"true"`, `"True"`, `"0"`, `"f"`, `"F"`, `"FALSE"`, `"false"`, `"False"` |
| `FormatBool()` | Converts `bool` to `string` — returns `"true"` or `"false"` |
| `ParseFloat()` | Converts `string` to `float64` — returns `(float64, error)`. Second param is bitSize (`32` or `64`) |
| `FormatFloat()` | Converts `float64` to `string` with formatting — params: `(value, fmt byte, prec int, bitSize int)` |

**Example:**

```go
// ParseBool — string to bool
result, err := strconv.ParseBool("true")
fmt.Println(result)      // true

parseBool := strconv.FormatBool(result)
fmt.Println(parseBool)   // true

// Atoi / Itoa — string/int conversion
num, _ := strconv.Atoi("123")
fmt.Println(num)         // 123

text := strconv.Itoa(123)
fmt.Println(text)        // "123"

// ParseFloat / FormatFloat — float conversion
f, _ := strconv.ParseFloat("123.45", 64)
fmt.Println(f)           // 123.45

formatted := strconv.FormatFloat(f, 'f', 2, 64)
fmt.Println(formatted)   // 123.45
```

**FormatFloat format byte (`fmt`) options:**

| Byte | Description | Example |
|------|-------------|---------|
| `'f'` | Decimal notation | `123.45` |
| `'e'` | Scientific notation (lowercase) | `1.234500e+02` |
| `'E'` | Scientific notation (uppercase) | `1.234500E+02` |

**Notes:**
- `Itoa` = **Integer to ASCII**, `Atoi` = **ASCII to Integer**
- `ParseBool()` is **case-insensitive** — accepts `"true"`, `"TRUE"`, `"True"`, `"t"`, `"1"`, etc.
- `Parse` functions return `(value, error)` — always check the error. `Format` functions return just `string` (no error).

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

---

### `os` — Operating System

```go
import "os"
```

**Functions used:**

| Function | Description |
|----------|-------------|
| `Args` | Command-line arguments (variable) — returns `[]string` |
| `Hostname()` | Returns the hostname of the machine |
| `Getenv()` | Gets an environment variable by key (returns empty string if missing) |
| `LookupEnv()` | Gets an env var with a boolean indicating if it exists |
| `Setenv()` | Sets an environment variable |
| `Unsetenv()` | Unsets/deletes an environment variable |
| `Environ()` | Returns all environment variables as `[]string` in `"KEY=VALUE"` format |

**Example — Command-Line Arguments:**

```go
args := os.Args

fmt.Println("Arguments:", len(args))
for i, arg := range args {
    fmt.Println("Index:", i, "Arg:", arg)
}
```

Run with: `go run main.go arg1 arg2`

**Example — Environment Variables:**

```go
e := os.Getenv("SAMPLE_ENV")
fmt.Println("SAMPLE_ENV:", e)

value, isExist := os.LookupEnv("SAMPLE_ENV")
fmt.Println("Value:", value, "Exists:", isExist)

os.Setenv("SAMPLE_ENV", "hi_env")
value, isExist = os.LookupEnv("SAMPLE_ENV")
fmt.Println("Value:", value, "Exists:", isExist)

os.Unsetenv("SAMPLE_ENV")
value, isExist = os.LookupEnv("SAMPLE_ENV")
fmt.Println("Value:", value, "Exists:", isExist)
```

**Example — Hostname:**

```go
host, err := os.Hostname()
if err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Println("Hostname:", host)
}
```

> **Note:** `Args[0]` is the program name itself — actual arguments start at `Args[1]`. Use `LookupEnv()` when you need to distinguish between an empty env var and a missing one.

---

### `flag` — Command-Line Flag Parsing

```go
import "flag"
```

**Functions used:**

| Function | Description |
|----------|-------------|
| `String()` | Declares a string flag with name, default value, and description |
| `Int()` | Declares an integer flag with name, default value, and description |
| `Parse()` | Parses the command-line flags — call after declaring all flags |

**Example:**

```go
host := flag.String("host", "localhost", "host description")
port := flag.Int("port", 8080, "port description")
user := flag.String("user", "admin", "user description")
password := flag.String("password", "123456", "password description")

flag.Parse()

fmt.Println("Host:", *host)
fmt.Println("Port:", *port)
fmt.Println("User:", *user)
fmt.Println("Password:", *password)
```

Run with: `go run main.go -host=localhost -port=8080 -user=root -password=123456`

> **Note:** `flag.String()` and `flag.Int()` return **pointers** — dereference with `*` to get the value. `flag.Parse()` must be called after declaring all flags and before accessing their values. Flags can be passed in any order — no need to match positional indexes like `os.Args`.

---

---

### `strings` — String Manipulation

```go
import "strings"
```

**Functions used:**

| Function | Description |
|----------|-------------|
| `Contains()` | Checks if a string contains a substring |
| `Count()` | Counts non-overlapping occurrences of a substring |
| `Index()` | Returns the index of the first occurrence of a substring (-1 if not found) |
| `Repeat()` | Repeats a string n times |
| `Replace()` | Replaces the first n occurrences of a substring |
| `ReplaceAll()` | Replaces all occurrences of a substring |
| `Split()` | Splits a string by a separator into a slice |
| `Title()` | Converts each word to title case (first letter uppercase) |
| `ToLower()` | Converts to lowercase |
| `ToUpper()` | Converts to uppercase |
| `HasPrefix()` | Checks if a string starts with a prefix |
| `HasSuffix()` | Checks if a string ends with a suffix |
| `TrimSpace()` | Removes leading and trailing whitespace |

**Example:**

```go
s := "hello string"

fmt.Println(strings.Contains(s, "string"))   // true
fmt.Println(strings.Count(s, "l"))             // 2
fmt.Println(strings.Index(s, "string"))        // 6
fmt.Println(strings.Repeat("ha", 5))           // hahahahaha
fmt.Println(strings.Replace(s, "o", "x", 1))   // hellx string
fmt.Println(strings.ReplaceAll(s, "o", "x"))   // hellx string
fmt.Println(strings.Split(s, "o"))             // [hell  string]
fmt.Println(strings.Title(s))                  // Hello String
fmt.Println(strings.ToLower("HELLO"))          // hello
fmt.Println(strings.ToUpper("hello"))          // HELLO
fmt.Println(strings.HasPrefix(s, "hello"))     // true
fmt.Println(strings.HasSuffix(s, "string"))    // true

s2 := "      password          "
fmt.Println("'" + strings.TrimSpace(s2) + "'")   // 'password'
```

> **Note:** Strings are **immutable** in Go — all `strings` functions return a **new string**, they never modify the original.

---

### `math` — Math Operations

```go
import "math"
```

**Functions used:**

| Function | Description | Example |
|----------|-------------|---------|
| `Abs()` | Absolute value | `Abs(-10.5)` → `10.5` |
| `Max()` | Returns the larger of two values | `Max(10, 20)` → `20` |
| `Min()` | Returns the smaller of two values | `Min(10, 20)` → `10` |
| `Round()` | Rounds to nearest integer (half up) | `Round(3.6)` → `4` |
| `Ceil()` | Rounds up | `Ceil(3.2)` → `4` |
| `Floor()` | Rounds down | `Floor(3.8)` → `3` |
| `Pow()` | Power (x^y) | `Pow(2, 3)` → `8` |
| `Sqrt()` | Square root | `Sqrt(9)` → `3` |
| `Mod()` | Modulo (like `%` but works with floats too) | `Mod(5, 2)` → `1`, `Mod(5.5, 2)` → `1.5` |

**Constants:**

| Constant | Value |
|----------|-------|
| `Pi` | `3.141592653589793` |

**Example:**

```go
fmt.Println(math.Abs(-10.5))  // 10.5
fmt.Println(math.Max(10, 20)) // 20
fmt.Println(math.Min(10, 20)) // 10
fmt.Println(math.Round(3.6))  // 4
fmt.Println(math.Ceil(3.2))   // 4
fmt.Println(math.Floor(3.8))  // 3
fmt.Println(math.Pow(2, 3))   // 8
fmt.Println(math.Sqrt(9))     // 3
fmt.Println(math.Mod(5, 2))   // 1
fmt.Println(math.Mod(5.5, 2)) // 1.5
fmt.Println(math.Pi)          // 3.141592653589793
```

> **Note:** `math` functions work with `float64`. Use `Mod()` instead of `%` when working with floats — `%` only works with integers.

---

### `math/rand/v2` — Random Number Generation

```go
import "math/rand/v2"
```

**Functions used:**

| Function | Description |
|----------|-------------|
| `Int()` | Returns a random `int` (non-negative) |
| `IntN(n)` | Returns a random `int` in `[0, n)` |
| `Float64()` | Returns a random `float64` in `[0.0, 1.0)` |

**Example:**

```go
fmt.Println("Random Int:", rand.Int())       // e.g. 2470256555260306322
fmt.Println("Random IntN (0-9):", rand.IntN(10)) // e.g. 4
fmt.Println("Random Float:", rand.Float64())  // e.g. 0.4031342119625486

// Random float in custom range [min, max)
min := 10.0
max := 20.0
fmt.Println("Random FloatN:", rand.Float64()*(max-min)) // e.g. 3.5
```

> **Note:** `math/rand/v2` is the newer Go 1.22+ version of the rand package — uses different algorithms than the original `math/rand`. Functions like `IntN()` don't exist in `math/rand` — that version uses `Intn()` (lowercase 'n').

---

> **Note:** There may be other packages I haven't documented here. For the full list, check out the [Go Standard Library Docs](https://pkg.go.dev/std).
