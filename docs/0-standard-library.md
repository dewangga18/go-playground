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

### `container/list` — Doubly-Linked List

```go
import "container/list"
```

**Functions used:**

| Function | Description |
|----------|-------------|
| `list.New()` | Creates a new empty doubly-linked list |
| `PushBack(v)` | Adds an element to the back of the list |
| `PushFront(v)` | Adds an element to the front of the list |
| `Front()` | Returns the first element (`*Element`) — `nil` if empty |
| `Back()` | Returns the last element (`*Element`) — `nil` if empty |
| `Len()` | Returns the number of elements in the list |
| `InsertBefore(v, mark)` | Inserts a new element before the given element |
| `InsertAfter(v, mark)` | Inserts a new element after the given element |
| `MoveBefore(e, mark)` | Moves an existing element before another |
| `MoveAfter(e, mark)` | Moves an existing element after another |
| `Remove(e)` | Removes an element from the list |

**Element fields:**

| Field | Description |
|-------|-------------|
| `.Value` | The value stored in the element (`any`) — type-assert as needed |
| `.Next()` | Returns the next element or `nil` |
| `.Prev()` | Returns the previous element or `nil` |

**Example:**

```go
l := list.New()
l.PushBack("B")
l.PushBack("C")
l.PushBack("D")
l.PushFront("A")

fmt.Println("Length:", l.Len())     // 4

// Iterate forward
for e := l.Front(); e != nil; e = e.Next() {
    fmt.Println(e.Value)            // A, B, C, D
}

// Iterate backward
for e := l.Back(); e != nil; e = e.Prev() {
    fmt.Println(e.Value)            // D, C, B, A
}

k := list.New()
k.PushBack(1)
k.PushBack(2)

// Insert before / after
a := k.PushBack(4)
k.InsertBefore(3, a)
b := k.PushBack(5)
k.InsertAfter(6, b)                  // 1, 2, 3, 4, 5, 6

// Move elements
first := k.PushFront(0)
k.MoveBefore(first, k.Front())       // moves 0 before itself (no-op effectively)

last := k.PushBack(7)
k.MoveAfter(last, k.Back())          // moves 7 after itself (no-op effectively)

// Remove elements
k.Remove(k.Front())                  // removes first element
```

**Full output:**

```
Length: 4
Iteration Forward
A
B
C
D
Iteration Backward
D
C
B
A
1
2
3
4
5
6
move before function
0
1
2
3
4
5
6
move after function
0
1
2
3
4
5
6
7
remove function
1
2
3
4
5
6
7
```

> **Important:** `container/list` is **pointer-based**. Elements are accessed via `*Element` pointers (`Next()`, `Prev()`, `Front()`, `Back()`). There's no built-in way to deep-copy or clone a list — the only way is to iterate through the original and build a new one with `PushBack()`/`PushFront()`.

**Queue pattern (FIFO):**

```go
queue := list.New()

// Enqueue — add to back
queue.PushBack("job1")
queue.PushBack("job2")
queue.PushBack("job3")

// Dequeue — remove from front
for queue.Len() > 0 {
    e := queue.Front()
    fmt.Println(e.Value)   // job1, job2, job3
    queue.Remove(e)
}
```

**Stack pattern (LIFO):**

```go
stack := list.New()

// Push — add to front
stack.PushFront("a")
stack.PushFront("b")
stack.PushFront("c")

// Pop — remove from front
for stack.Len() > 0 {
    e := stack.Front()
    fmt.Println(e.Value)   // c, b, a
    stack.Remove(e)
}
```

**When to use `container/list` vs slice:**

| Scenario | Use | Why |
|----------|-----|-----|
| Frequent insert/delete in the **middle** | `container/list` | Slice needs shifting — expensive for large data |
| Queue (FIFO) or Stack (LIFO) | **Either** | List is cleaner. Slice also works (`append` + reslice) but needs index tracking |
| Random access by index (`list[500]`) | **Slice** | List must iterate from Front/Back — O(n) |
| Small data (< 100 items) | **Slice** | Simpler, performance difference is negligible |
| Cache (LRU, etc.) | `container/list` | Built-in move-to-front/back, remove — perfect for eviction tracking |

> **TL;DR:** Default to slice. `container/list` shines when you need frequent insert/delete at arbitrary positions, or built-in move operations (like LRU cache).

---

### `container/ring` — Circular List

```go
import "container/ring"
```

Circular ring — like a list that wraps around. **No start or end.** The `*Ring` pointer always points to "current" position, and you move with `Next()`/`Prev()` or `Move(n)`.

**Functions used:**

| Function | Description |
|----------|-------------|
| `ring.New(n)` | Creates a new ring with `n` zero-valued elements |
| `Len()` | Returns the number of elements in the ring |
| `Do(fn)` | Calls `fn` on every element — iterates forward from current position |
| `Move(n)` | Moves the ring pointer forward (`n > 0`) or backward (`n < 0`) — returns new `*Ring` |
| `Link(r)` | Links another ring `r` after the current element — merges two rings |
| `Unlink(n)` | Removes `n` elements after the current element (not including current) |

**Element field:**

| Field | Description |
|-------|-------------|
| `.Value` | The value stored in the element (`any`) — type-assert as needed |

**Example:**

```go
r := ring.New(5)

// Populate ring
for i := 0; i < r.Len(); i++ {
    r.Value = "Value " + strconv.Itoa(i+1)
    r = r.Next()
}

// Print all
r.Do(func(i any) {
    fmt.Println(i)           // Value 1, Value 2, ..., Value 5
})

// Move 2 positions forward
r = r.Move(2)
r.Do(func(i any) {
    fmt.Println(i)           // Value 3, Value 4, Value 5, Value 1, Value 2
})

// Link — merge another ring
r2 := ring.New(2)
r2.Value = "Value 6"
r2.Next().Value = "Value 7"
r.Link(r2)                   // inserts r2's elements after current position
r.Do(func(i any) {
    fmt.Println(i)           // Value 3, Value 6, Value 7, Value 4, Value 5, Value 1, Value 2
})

// Unlink — remove n elements after current
r.Unlink(1)                  // removes the element after current (Value 6)
r.Do(func(i any) {
    fmt.Println(i)           // Value 3, Value 7, Value 4, Value 5, Value 1, Value 2
})
```

**Full output:**

```
Value 1
Value 2
Value 3
Value 4
Value 5

Move 2
Value 3
Value 4
Value 5
Value 1
Value 2

Link
Value 3
Value 6
Value 7
Value 4
Value 5
Value 1
Value 2

Unlink
Value 3
Value 7
Value 4
Value 5
Value 1
Value 2
```

> **Note:** Unlike `container/list`, `container/ring` has **no zero-value**. Must create with `ring.New(n)`. The ring always has a current position — operations like `Link()` and `Unlink()` happen relative to that position. `Unlink()` does **not** remove the current element, only elements after it.

**When to use `container/ring`:**

| Scenario | Use | Why |
|----------|-----|-----|
| Fixed-size buffer (overwrite oldest) | `container/ring` | Circular — no need to track head/tail manually |
| Round-robin scheduler | `container/ring` | `Move(n)` advances to next participant naturally |
| Something simpler? | **Slice with index** | Rings are niche. Most cases work fine with a slice + modulo index |

> **TL;DR:** Rings are niche. Only reach for this when you truly need a circular buffer — else slice + `%` index is simpler.

---

### `sort` — Slice Sorting

```go
import "sort"
```

**Functions used:**

| Function | Description |
|----------|-------------|
| `Ints(slice)` | Sorts `[]int` in-place (ascending) |
| `Strings(slice)` | Sorts `[]string` in-place (ascending) |
| `Float64s(slice)` | Sorts `[]float64` in-place (ascending) |
| `Sort(data)` | Sorts any slice that implements `sort.Interface` (`Len`, `Less`, `Swap`) |
| `Slice(slice, less)` | Sorts any slice using a `less` function — no need to implement `sort.Interface` |
| `Reverse(data)` | Wraps a `sort.Interface` to sort in **descending** order — use with `sort.Sort()` |

**Type adapters (implement `sort.Interface`):**

| Type | For sorting |
|------|-------------|
| `sort.IntSlice` | `[]int` |
| `sort.StringSlice` | `[]string` |
| `sort.Float64Slice` | `[]float64` |

**Example — built-in types:**

```go
ages := []int{10, 20, 30, 5, 15}
sort.Ints(ages)
fmt.Println(ages)      // [5 10 15 20 30]

// Reverse
sort.Sort(sort.Reverse(sort.IntSlice(ages)))
fmt.Println(ages)      // [30 20 15 10 5]

names := []string{"John", "Doe", "Jane", "Bob"}
sort.Strings(names)
fmt.Println(names)     // [Bob Doe Jane John]

floats := []float64{1.0, 2.0, 3.0, 5.0, 1.5}
sort.Float64s(floats)
fmt.Println(floats)    // [1 1.5 2 3 5]
```

**Example — slice of structs (2 ways):**

**Way 1: Implement `sort.Interface`**

```go
type User struct {
    Name string
    Age  string
}

type UserSlice []User

func (u UserSlice) Len() int           { return len(u) }
func (u UserSlice) Less(i, j int) bool { return u[i].Age < u[j].Age }
func (u UserSlice) Swap(i, j int)      { u[i], u[j] = u[j], u[i] }

users := []User{
    {"John", "20"},
    {"Doe", "25"},
    {"Jane", "22"},
    {"Bob", "28"},
}

sort.Sort(UserSlice(users))
fmt.Println(users)     // [{John 20} {Jane 22} {Doe 25} {Bob 28}]
```

**Way 2: `sort.Slice` — simpler, no interface needed**

```go
sort.Slice(users, func(i, j int) bool {
    return users[i].Age < users[j].Age
})
fmt.Println(users)     // [{John 20} {Jane 22} {Doe 25} {Bob 28}]
```

> **Note:** `sort.Ints()`, `sort.Strings()`, `sort.Float64s()` modify the slice **in-place** — no return value. For custom types, `sort.Slice()` is more convenient than implementing `sort.Interface`. Both sort by `Age` as string (lexicographic order since `Age` is `string`).

---

### `time` — Time & Date

```go
import "time"
```

**Functions used:**

| Function | Description |
|----------|-------------|
| `time.Now()` | Returns current local time as `time.Time` |
| `time.Date(year, month, day, hour, min, sec, nsec, loc)` | Creates a `time.Time` at specific date/location |
| `time.Parse(layout, value)` | Parses a string into `time.Time` based on layout |
| `time.LoadLocation(name)` | Loads a timezone by name (e.g. `"Asia/Jakarta"`) |
| `time.ParseDuration(s)` | Parses a duration string like `"2h30m"` into `time.Duration` |

**Method on `time.Time`:**

| Method | Returns |
|--------|---------|
| `.Local()` | Convert to local timezone |
| `.UTC()` | Convert to UTC |
| `.In(loc)` | Convert to a specific `*time.Location` |
| `.Zone()` | `(name string, offset int)` — e.g. `("WIB", 25200)` |
| `.Year()` | `int` — year component |
| `.Month()` | `time.Month` — month name |
| `.Day()` | `int` — day of month |
| `.Hour()` | `int` — hour (0–23) |
| `.Unix()` | `int64` — Unix timestamp (seconds since epoch) |
| `.UnixNano()` | `int64` — Unix timestamp in nanoseconds |
| `.Format(layout)` | `string` — format time using a layout |

**Built-in layout constants:**

| Constant | Layout | Example |
|----------|--------|---------|
| `time.RFC3339` | `"2006-01-02T15:04:05Z07:00"` | `2022-07-22T07:32:22Z` |
| `time.RFC3339Nano` | `"2006-01-02T15:04:05.999999999Z07:00"` | `2022-07-22T07:32:22.123456Z` |
| `time.RFC1123Z` | `"Mon, 02 Jan 2006 15:04:05 -0700"` | `Thu, 22 Jul 2022 07:32:22 +0000` |
| `time.DateTime` | `"2006-01-02 15:04:05"` | `2022-07-22 07:32:22` |
| `time.DateOnly` | `"2006-01-02"` | `2022-07-22` |
| `time.TimeOnly` | `"15:04:05"` | `07:32:22` |

**Common custom layouts (reference time: `Mon Jan 2 15:04:05 MST 2006`):**

> Go uses a unique approach — you write the layout using the reference time `01/02 03:04:05PM '06 -0700`. The actual numbers don't matter, only their position matters.

| Layout | Use case | Example |
|--------|----------|---------|
| `"2006-01-02 15:04:05"` | MySQL / SQL datetime | `2024-07-07 19:33:00` |
| `"2006-01-02"` | Date only (e.g. birthdate, logs) | `2024-07-07` |
| `"15:04:05"` | Time only | `19:33:00` |
| `"2006-01-02 15:04:05 -0700"` | With timezone offset | `2024-07-07 19:33:00 +0700` |
| `"2006-01-02T15:04:05Z0700"` | ISO 8601 (no colon in offset) | `2024-07-07T19:33:00+0700` |
| `"2006-01-02 15:04:05 MST"` | With timezone name | `2024-07-07 19:33:00 WIB` |
| `"02-Jan-2006"` | DD-Mon-YYYY (log files, invoices) | `07-Jul-2024` |
| `"Mon, 02 Jan 2006"` | HTTP headers | `Sun, 07 Jul 2024` |
| `"Monday, 02 January 2006"` | Full human-readable | `Sunday, 07 July 2024` |

**Key layout parts:**

| Reference token | Means | Example output |
|----------------|-------|---------------|
| `2006` | Year (4 digits) | `2024` |
| `06` | Year (2 digits) | `24` |
| `01` or `1` | Month number | `07` or `7` |
| `January` or `Jan` | Month name | `July` or `Jul` |
| `02` or `2` | Day number | `07` or `7` |
| `Monday` or `Mon` | Day name | `Sunday` or `Sun` |
| `15` | Hour (24-hour) | `19` |
| `03` or `3` | Hour (12-hour) | `07` or `7` |
| `PM` or `pm` | AM/PM | `PM` or `pm` |
| `04` or `4` | Minutes | `33` or `33` |
| `05` or `5` | Seconds | `00` or `0` |
| `.000` or `.999` | Milliseconds | `.123` |
| `-0700` | Timezone offset (numeric) | `+0700` |
| `-07:00` | Timezone offset (with colon) | `+07:00` |
| `MST` | Timezone name | `WIB`, `UTC` |

**Example:**

```go
now := time.Now()

fmt.Println("now       :", now)                    // 2026-07-07 09:38:04.429564 +0700 WIB
fmt.Println("Local     :", now.Local())             // 2026-07-07 09:38:04.429564 +0700 WIB
fmt.Println("Zone      :", now.Zone())              // WIB 25200
fmt.Println("UTC       :", now.UTC())               // 2026-07-07 02:38:04.429564 +0000 UTC
fmt.Println("Unix      :", now.Unix())              // 1783391884
fmt.Println("UnixNano  :", now.UnixNano())           // 1783391884429564000

utc := time.Date(2022, time.July, 22, 5, 0, 0, 0, time.UTC)
wib := time.Date(2024, 07, 7, 7, 0, 0, 0, time.Local)

fmt.Println(utc.UTC())     // 2022-07-22 05:00:00 +0000 UTC
fmt.Println(wib.UTC())     // 2024-07-07 00:00:00 +0000 UTC

// Parse from string
parseRFC3339, _ := time.Parse(time.RFC3339, "2022-07-22T07:32:22Z")
fmt.Println("parseRFC3339:", parseRFC3339)           // 2022-07-22 07:32:22 +0000 UTC

// Parse with custom layout
parseMyTime, _ := time.Parse("2006-01-02 15:04:05", "2024-07-07 19:33:00")
fmt.Println("parseMyTime:", parseMyTime)             // 2024-07-07 19:33:00 +0000 UTC

// Format time to string
fmt.Println(now.Format("2006-01-02 15:04:05"))      // 2026-07-07 09:38:04

// Extract components
fmt.Println("Year:", utc.Year())    // 2022
fmt.Println("Month:", utc.Month())   // July
fmt.Println("Day:", utc.Day())     // 22
fmt.Println("Hour:", utc.Hour())    // 5
```

**Common timezone locations in Indonesia and abroad:**

```go
import "time"

// Indonesia
wib, _ := time.LoadLocation("Asia/Jakarta")      // UTC+7
wita, _ := time.LoadLocation("Asia/Makassar")   // UTC+8
wit, _ := time.LoadLocation("Asia/Jayapura")    // UTC+9

// International
utc, _ := time.LoadLocation("UTC")                // UTC±0
london, _ := time.LoadLocation("Europe/London")  // UTC±0 / +1 (DST)
tokyo, _ := time.LoadLocation("Asia/Tokyo")      // UTC+9
nyc, _ := time.LoadLocation("America/New_York") // UTC-5 / -4 (DST)
```

> **Note:** The reference time in Go is `Mon Jan 2 15:04:05 MST 2006` = `01/02 03:04:05PM '06 -0700`. It's easier to remember as "1 2 3 4 5 6 7" (month, day, hour, minute, second, year, timezone). For Indonesia, `Asia/Jakarta` is the standard IANA timezone — use `time.LoadLocation("Asia/Jakarta")` instead of hardcoding `+7`.

---

### `time.Duration` — Time Intervals

```go
import "time"
```

`time.Duration` is a type representing elapsed time in nanoseconds. Built from constants like `time.Second`, `time.Minute`, `time.Hour`.

**Duration constants:**

| Constant | Approx value |
|----------|-------------|
| `time.Nanosecond` | `1 ns` |
| `time.Microsecond` | `1000 ns` |
| `time.Millisecond` | `1,000,000 ns` |
| `time.Second` | `1,000,000,000 ns` |
| `time.Minute` | `60 s` |
| `time.Hour` | `60 min` |

**Methods on `time.Duration`:**

| Method | Returns |
|--------|---------|
| `.Nanoseconds()` | `int64` — duration as total nanoseconds |
| `.Microseconds()` | `int64` — duration as total microseconds |
| `.Milliseconds()` | `int64` — duration as total milliseconds |
| `.Seconds()` | `float64` — duration in seconds (with decimals) |
| `.Minutes()` | `float64` — duration in minutes |
| `.Hours()` | `float64` — duration in hours |

**Example:**

```go
duration1 := time.Second * 100         // 100 seconds
duration2 := time.Minute * 10          // 10 minutes
duration3 := time.Hour * 1             // 1 hour

fmt.Println("Seconds", duration1.Seconds())   // 100
fmt.Println("Minutes", duration2.Minutes())   // 10
fmt.Println("Hours", duration3.Hours())       // 1

// Arithmetic — durations support +, -, *, /
diff := duration3 - duration2 - duration1
fmt.Println("Duration", diff)                 // 48m20s

// Parse from string
parseDuration, _ := time.ParseDuration("2h30m")
fmt.Println("ParseDuration", parseDuration)            // 2h30m0s
fmt.Println("ParseDuration hour", parseDuration.Hours())  // 2.5
fmt.Println("ParseDuration min", parseDuration.Minutes()) // 150
```

> **Note:** Durations support arithmetic (`+`, `-`, `*`, `/`). When printed with `fmt.Println`, Go automatically formats them as human-readable strings like `48m20s` or `1h30m`. This works because `time.Duration` has a custom `.String()` method. `time.ParseDuration()` accepts strings like `"300ms"`, `"2h30m"`, `"1.5s"`, `"-10m"` — supports `ns`, `us`/`µs`, `ms`, `s`, `m`, `h`. For the full list, check out the [Go Standard Library Docs](https://pkg.go.dev/std).

---

### `reflect` — Runtime Reflection

```go
import "reflect"
```

Reflection lets you inspect and manipulate types/values **at runtime** — useful for generic utilities, serialization, validation, and testing.

**Functions used:**

| Function | Description |
|----------|-------------|
| `reflect.TypeOf(i)` | Returns the `reflect.Type` of the value — metadata about the type itself |
| `reflect.ValueOf(i)` | Returns the `reflect.Value` — the actual value with methods to read/modify it |

**Methods on `reflect.Type`:**

| Method | Returns |
|--------|---------|
| `.Name()` | `string` — type name (e.g. `"Sample"`, `"int"`) |
| `.Kind()` | `reflect.Kind` — the **underlying kind** (e.g. `struct`, `ptr`, `slice`, `map`, `string`, `int`) |
| `.NumField()` | `int` — number of fields (for structs) |
| `.Field(i)` | `reflect.StructField` — info about the i-th field (name, type, tags) |
| `.NumMethod()` | `int` — number of exported methods |
| `.Method(i)` | `reflect.Method` — info about the i-th method (name, type) |

**Methods on `reflect.Value`:**

| Method | Returns / Behavior |
|--------|-------------------|
| `.Kind()` | `reflect.Kind` — same as Type's Kind |
| `.Elem()` | `reflect.Value` — dereferences a pointer/interface |
| `.CanSet()` | `bool` — whether the value can be modified (only if **addressable**) |
| `.SetInt(i)` | Sets the int value — panics if type mismatch |
| `.SetString(s)` | Sets the string value |
| `.SetFloat(f)` | Sets the float value |
| `.FieldByName(name)` | `reflect.Value` — gets struct field by name |
| `.MethodByName(name)` | `reflect.Value` — gets method by name |
| `.Call(args)` | `[]reflect.Value` — calls the method with given arguments |
| `.Int()` | `int64` — reads the int value |
| `.String()` | `string` — reads the string value |

**`reflect.Kind` — the underlying type category:**

| Kind | Description |
|------|-------------|
| `struct` | Struct type |
| `ptr` | Pointer |
| `slice` | Slice |
| `map` | Map |
| `string` | String |
| `int`, `int8`, ..., `int64` | Signed integers |
| `float32`, `float64` | Floats |
| `bool` | Boolean |
| `func` | Function |
| `interface` | Interface |
| `array` | Array (fixed-size) |

> `Kind()` tells you **what** the type fundamentally is, regardless of custom type names. E.g. both `type Age int` and `int` have `Kind() == int`.

**Example:**

```go
// TypeOf & ValueOf
sample := Sample{"Uhuyy", "23"}
sampleType := reflect.TypeOf(sample)
sampleValue := reflect.ValueOf(sample)

fmt.Println(sampleType.Name())                                    // Sample
fmt.Println(sampleValue.FieldByName("Name").String())            // Uhuyy

// Kind
fmt.Println(sampleType.Kind())                                    // struct
fmt.Println(reflect.TypeOf(&sample).Kind())                       // ptr

var nums []int
fmt.Println(reflect.TypeOf(nums).Kind())                          // slice

var m map[string]int
fmt.Println(reflect.TypeOf(m).Kind())                             // map
```

**Struct Fields & Tags:**

```go
for i := 0; i < sampleType.NumField(); i++ {
    field := sampleType.Field(i)
    fmt.Printf("Field %d: %s (%s)\n", i, field.Name, field.Type)

    required := field.Tag.Get("required")
    max := field.Tag.Get("max")
    fmt.Printf("  required: %q, max: %q\n", required, max)
}
```

```
  Field 0: Name (string)
    required: "true", max: "10"
  Field 1: Age (string)
    required: "", max: ""
```

**Elem — dereference pointer:**

```go
num := 42
ptr := reflect.ValueOf(&num)
elem := ptr.Elem()                    // dereference → gets the int Value

fmt.Println(ptr.Kind())               // ptr
fmt.Println(elem.Kind())              // int
fmt.Println(elem.Int())               // 42
```

**CanSet & Set — modify values through pointer:**

```go
// Must pass pointer + use Elem() to get an addressable value
elem.SetInt(100)
fmt.Println(num)                      // 100

// Modify struct fields
p := Person{"Budi", 25}
pv := reflect.ValueOf(&p).Elem()      // must pass pointer!

fmt.Println(pv.FieldByName("Name").CanSet())   // true
pv.FieldByName("Name").SetString("Agus")
pv.FieldByName("Age").SetInt(30)
fmt.Println(p)                                  // {Agus 30}
```

**Methods — iterate and call dynamically:**

```go
calc := Calculator{Value: 10}
calcType := reflect.TypeOf(calc)
calcValue := reflect.ValueOf(calc)

fmt.Println(calcType.NumMethod())               // 2

// List methods
for i := 0; i < calcType.NumMethod(); i++ {
    method := calcType.Method(i)
    fmt.Println(method.Name, method.Type)       // e.g. Add (func(main.Calculator, int) int)
}

// Call methods dynamically
result := calcValue.MethodByName("Add").Call([]reflect.Value{reflect.ValueOf(5)})
fmt.Println(result[0].Int())                    // 15

result = calcValue.MethodByName("Mul").Call([]reflect.Value{reflect.ValueOf(3)})
fmt.Println(result[0].Int())                    // 30
```

**Structs used in examples:**

```go
type Sample struct {
    Name string `required:"true" max:"10"`
    Age  string
}

type Person struct {
    Name string
    Age  int
}

type Calculator struct {
    Value int
}

func (c Calculator) Add(n int) int { return c.Value + n }
func (c Calculator) Mul(n int) int { return c.Value * n }
```

> **Key rule for `Set`:** You can only `Set` a value that is **addressable** — meaning it came from a pointer, a slice element, a map entry, or a field of an addressable struct. A value from `reflect.ValueOf(someVar)` (value, not pointer) is **never** addressable. Always use `reflect.ValueOf(&x).Elem()` to get a settable value.

> **When to use reflect:** Validation libraries, ORMs/serializers, generic pretty-printers, testing utilities. Go's static typing usually makes reflection unnecessary for application code. Use sparingly — reflection is slower, less type-safe, and harder to read than explicit code.

---

### `regexp` — Regular Expressions

```go
import "regexp"
```

**Functions used:**

| Function | Description |
|----------|-------------|
| `Compile(pattern)` | Compiles a regex pattern — returns `(*Regexp, error)`. Always check the error! |
| `MustCompile(pattern)` | Compiles a regex pattern or **panics** — use when you're certain the pattern is valid (e.g. hardcoded literals) |
| `MatchString(pattern, s)` | Checks if the pattern matches **anywhere** in the string — returns `bool` |
| `FindString(s)` | Returns the **first** match as a string, or empty string if no match |
| `FindAllString(s, n)` | Returns **all** matches as `[]string`. `n = -1` for all, `n >= 0` to limit results |
| `ReplaceAllString(s, repl)` | Replaces all matches with `repl`. Supports `$1`, `$2`, etc. for capture group references |
| `Split(s, n)` | Splits string by the pattern (like `strings.Split` but uses regex). `n = -1` for all, `n >= 0` to limit |
| `FindStringSubmatch(s)` | Returns full match + capture groups as `[]string` — `[0]` = full match, `[1]` = first group, etc. |

**Methods on `*Regexp`:**

| Method | Returns |
|--------|---------|
| `.SubexpIndex(name)` | `int` — index of a named capture group (`(?P<name>...)`), for use with `FindStringSubmatch` results |

**Example — Compile vs MustCompile:**

```go
// Compile — returns error for invalid patterns
re, err := regexp.Compile(`golang`)
if err != nil {
    fmt.Println("Error:", err)
    return
}

// MustCompile — panics on invalid pattern, use for hardcoded patterns
re2 := regexp.MustCompile(`golang`)
```

**Example — MatchString:**

```go
text := "golang regexp is fun and golang is awesome"

fmt.Println(regexp.MustCompile(`golang`).MatchString(text))   // true
fmt.Println(regexp.MustCompile(`java`).MatchString(text))     // false
```

**Example — FindString (first match):**

```go
fmt.Println(regexp.MustCompile(`golang`).FindString(text))    // golang

reDigit := regexp.MustCompile(`\d+`)
fmt.Println(reDigit.FindString("order 99 price 500"))         // 99
```

**Example — FindAllString (all matches):**

```go
all := regexp.MustCompile(`golang`).FindAllString(text, -1)
fmt.Println(all)              // [golang golang]
fmt.Println(len(all))         // 2

// Limit results
limited := regexp.MustCompile(`golang`).FindAllString(text, 1)
fmt.Println(limited)          // [golang]
```

**Example — ReplaceAllString (with capture groups):**

```go
replaced := regexp.MustCompile(`golang`).ReplaceAllString(text, "Go")
fmt.Println(replaced)         // Go regexp is fun and Go is awesome

// Replace digits
replacedDigit := regexp.MustCompile(`\d+`).ReplaceAllString("phone 123, zip 456", "***")
fmt.Println(replacedDigit)    // phone ***, zip ***

// Capture group references ($1, $2, ...)
emailText := "user@example.com, admin@test.org, invalid-email"
reEmail := regexp.MustCompile(`(\w+)@(\w+\.\w+)`)
masked := reEmail.ReplaceAllString(emailText, "$1 at $2")
fmt.Println(masked)           // user at example.com, admin at test.org, invalid-email
```

**Example — Split:**

```go
csvLine := "a,b,c,d,e"
parts := regexp.MustCompile(`,`).Split(csvLine, -1)
fmt.Println(parts)            // [a b c d e]

// With limit — stops after n parts
limitedParts := regexp.MustCompile(`,`).Split(csvLine, 3)
fmt.Println(limitedParts)     // [a b c,d,e]

// Split on whitespace (handles multiple spaces)
words := regexp.MustCompile(`\s+`).Split("hello   world  foo", -1)
fmt.Println(words)            // [hello world foo]
```

**Example — FindStringSubmatch (capture groups):**

```go
logLine := "ERROR 2024-07-07 15:30:00 Connection timeout"
reLog := regexp.MustCompile(`(\w+) (\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}) (.+)`)
matches := reLog.FindStringSubmatch(logLine)

fmt.Println(matches[0])       // ERROR 2024-07-07 15:30:00 Connection timeout  (full match)
fmt.Println(matches[1])       // ERROR
fmt.Println(matches[2])       // 2024-07-07 15:30:00
fmt.Println(matches[3])       // Connection timeout
```

**Named capture groups:**

```go
reNamed := regexp.MustCompile(`(?P<name>\w+)@(?P<domain>\w+\.\w+)`)
emailMatch := reNamed.FindStringSubmatch("user@example.com")

fmt.Println("Name:", emailMatch[reNamed.SubexpIndex("name")])      // user
fmt.Println("Domain:", emailMatch[reNamed.SubexpIndex("domain")])  // example.com
```

**Full output of the example code:**

```
=== Compile — compile pattern (returns error if invalid) ===
Compiled: golang

=== MustCompile — compile or panic (use when pattern is certain) ===
MustCompiled: golang

=== MatchString — check if pattern matches anywhere ===
true
false

=== FindString — first match ===
golang
99

=== FindAllString — all matches (n = -1 for all) ===
[golang golang]
Count: 2
[golang]

=== ReplaceAllString — replace matches with new string ===
Go regexp is fun and Go is awesome
phone ***, zip ***
user at example.com, admin at test.org, invalid-email

=== Split — split string by pattern ===
[a b c d e]
[a b c,d,e]
[hello world foo]

=== FindStringSubmatch — match with capture groups ===
[ERROR 2024-07-07 15:30:00 Connection timeout ERROR 2024-07-07 15:30:00 Connection timeout]
Level: ERROR
Time: 2024-07-07 15:30:00
Message: Connection timeout
Name: user
Domain: example.com
```

> **Important:** `Compile()` returns an error for invalid patterns — use this when patterns come from user input. `MustCompile()` panics instead — use for hardcoded constants where a panic means a bug in your code. Named groups (`(?P<name>...)`) are more readable than positional groups — use `SubexpIndex()` to look up the index by name.
