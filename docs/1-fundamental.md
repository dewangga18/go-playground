## Fundamental Data Types

Go provides a rich set of built-in data types. Each type has a fixed size and range, ensuring consistent behavior across platforms.

### Numeric Types

```go
var a int8 = 127
var b uint8 = 255
var c float32 = 3.14
var d complex64 = 1 + 2i

fmt.Println(a, b, c, d)
```

#### Integers

| Type | Size | Range |
|------|------|-------|
| `int8` | 8-bit | -128 to 127 |
| `int16` | 16-bit | -32,768 to 32,767 |
| `int32` | 32-bit | -2,147,483,648 to 2,147,483,647 |
| `int64` | 64-bit | -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807 |
| `uint8` | 8-bit | 0 to 255 |
| `uint16` | 16-bit | 0 to 65,535 |
| `uint32` | 32-bit | 0 to 4,294,967,295 |
| `uint64` | 64-bit | 0 to 18,446,744,073,709,551,615 |

#### Floats

| Type | Size | Description |
|------|------|-------------|
| `float32` | 32-bit | IEEE 754 single-precision (~6 decimal digits) |
| `float64` | 64-bit | IEEE 754 double-precision (~15 decimal digits) |

#### Complex Numbers

| Type | Size | Description |
|------|------|-------------|
| `complex64` | 64-bit | Complex number with `float32` real and imaginary parts |
| `complex128` | 128-bit | Complex number with `float64` real and imaginary parts |

#### Aliases

| Alias | Underlying Type | Description |
|-------|-----------------|-------------|
| `byte` | `uint8` | Represents a single byte of data |
| `rune` | `int32` | Represents a Unicode code point |
| `int` | `int32` or `int64` | Platform-dependent; at least 32-bit |
| `uint` | `uint32` or `uint64` | Platform-dependent; at least 32-bit |

> **Note:** The size of `int` and `uint` depends on the architecture — 32-bit on 32-bit systems, 64-bit on 64-bit systems. They are the most efficient integer types for the target platform.

### Boolean Types

A `bool` type represents a boolean value — either `true` or `false`.

```go
var trueConstant bool = true
var falseConstant bool = false

fmt.Println(trueConstant)   // true
fmt.Println(falseConstant)  // false
fmt.Println(1 == 1)         // true
fmt.Println(1 != 1)         // false
```

### String Types

A `string` is a sequence of bytes. Strings are immutable in Go.

```go
var str1 string = "Hello"

fmt.Println(str1)               // Hello
fmt.Println(len(str1))          // 5
fmt.Println(str1[0])            // 72 (byte value, not char)
```

> **Note:** Indexing a string returns a byte value, not a character. For Unicode characters, use `rune`.

### Variable Declaration

Go provides several ways to declare variables.

#### var

Explicit declaration with type or inferred type.

```go
var a int8 = 127
var name = "Go"  // type inferred
```

#### const

Immutable values. Cannot be reassigned after declaration.

```go
const e = 42        // untyped constant
const f float64 = 2.718
const pi = 3.14159
```

#### Short Declaration (`:=`)

Shorthand for `var` inside functions. Type is inferred.

```go
count := 10
tmp := tempA
```

#### Multiple Declaration

Declare multiple variables in a single `var` block. Type can be specified per field or inferred.

```go
var (
	firstName = "Aaron"
	lastName  = "Evanjulio"
	version int = 1
)
```

#### Swap Pattern

Swap two variables using a temporary variable.

```go
tmp := tempA
tempA = tempB
tempB = tmp
```
