## Standard Library Reference

Progressive log of Go standard library packages encountered during learning. Each entry includes what the package does, the functions you've used, and where you first encountered it.

---

### `fmt`

**First encountered:** `1-basic-fundamental.md` (Go Fundamental)

Package `fmt` implements formatted I/O. It's the most commonly used package for printing output and reading input.

| Function | Format | Description |
|----------|--------|-------------|
| `fmt.Println()` | `fmt.Println(a, b, c)` | Print values separated by spaces, ending with a newline |

> **Note:** `fmt.Println` automatically adds spaces between arguments and appends a newline at the end.

---

### `strconv`

**First encountered:** `4-functions.md` (Go Functions)

Package `strconv` implements conversions between strings and basic data types. The name stands for **"string conversion"**.

| Function | Description | Example |
|----------|-------------|---------|
| `strconv.Itoa()` | Integer to ASCII — converts `int` to `string` | `strconv.Itoa(42)` → `"42"` |

```go
number := 11
result := strconv.Itoa(number) + " is Odd"
fmt.Println(result)   // 11 is Odd
```

> **Note:** `Itoa` stands for **I**nteger **to** **A**SCII. The reverse operation is `strconv.Atoi()` (string to int), but we haven't used it yet.

---

*More packages will be added here as you progress.*
