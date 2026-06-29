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
