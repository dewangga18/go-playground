## Standard Library

A quick reference for Go standard library packages used throughout these notes.<br/>

---

### `fmt` — Formatted I/O

```go
import "fmt"
```

**Common Functions**

| Function | Description |
|----------|-------------|
| `Print()` | Prints without a newline |
| `Println()` | Prints with a newline |
| `Printf()` | Prints using a format string |
| `Sprint()` | Returns a string |
| `Sprintf()` | Returns a formatted string |

**Example**

```go
fmt.Println("Hello", "World")

name := "John"
age := 20

fmt.Printf("%s is %d years old\n", name, age)
```
<br/>

---


### `strconv` — String Conversions

```go
import "strconv"
```

**Common Functions**

| Function | Description |
|----------|-------------|
| `Itoa()` | Converts `int` to `string` |
| `Atoi()` | Converts `string` to `int` |
| `ParseBool()` | Converts `string` to `bool` |
| `ParseFloat()` | Converts `string` to `float64` |

**Example**

```go
number := 10

text := strconv.Itoa(number)

fmt.Println(text)
```

**Notes**

- `Itoa` = **Integer to ASCII**
- `Atoi` = **ASCII to Integer**