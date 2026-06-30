## Go Fundamental 

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

### Numeric Operations

Go supports standard arithmetic operators for numeric types.

#### Arithmetic Operators

| Operator | Description | Example (`a=10, b=10, c=5, d=2, e=3`) |
|----------|-------------|----------------------------------------|
| `+` | Addition | `c + d` → `7` |
| `-` | Subtraction | `d - e` → `-1` |
| `*` | Multiplication | `c * d` → `10` |
| `/` | Division | `a / b` → `1` |
| `%` | Modulo (remainder) | `c % d` → `1` |

Arithmetic follows standard operator precedence — multiplication and division are evaluated before addition and subtraction.

```go
const a = 10
const b = 10
const c = 5
const d = 2
const e = 3

var i = a/b + c*d - e   // 10/10 + 5*2 - 3 = 1 + 10 - 3 = 8
fmt.Println("result:", i)   // 8
```

#### Augmented Assignment

Modify a variable by applying an operation and reassigning the result in one step.

```go
i := 8

i += 5    // i = i + 5  → 13
fmt.Println("i += 5:", i)

i -= 5    // i = i - 5  → 8
fmt.Println("i -= 5:", i)

i *= 5    // i = i * 5  → 40
fmt.Println("i *= 5:", i)

i /= 5    // i = i / 5  → 8
fmt.Println("i /= 5:", i)

i %= 5    // i = i % 5  → 3
fmt.Println("i %= 5:", i)
```

> **Note:** The `%` (modulo) operator and its augmented variant `%=` only work with integer types.

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

### Type Conversion

Convert between types explicitly. Go does not support implicit type conversion.

#### Byte to String

Indexing a string returns a byte (`uint8`). Store it in a variable first, then convert using `string()`.

```go
var firstName = "Aaron"
var byteVal uint8 = firstName[0]
var byteToStr = string(byteVal)

fmt.Println("Value of byteVal from firstName[0]:", byteVal)
fmt.Println("Convert byte to string:", byteToStr)
```

> **Note:** `string(byteValue)` treats the byte as a Unicode code point. This works for ASCII characters because their byte value equals their code point.

#### Integer Conversion

Convert between integer types using `T(value)` syntax. Beware of overflow when converting to a smaller type.

```go
var val32 int32 = 32769
var val64 int64 = int64(val32)    // 32769 — safe
var val16 int16 = int16(val32)    // overflow! 32769 > int16 max (32767)

fmt.Println("int32 to int64:", val64)
fmt.Println("int32 to int16:", val16, "// number overflow")
```

| Conversion | Result | Notes |
|-----------|--------|-------|
| `int32` → `int64` | Safe | Wider type fits all values |
| `int32` → `int16` | Overflow | Narrower type truncates bits |

> **Note:** Always check the target type's range before converting to a narrower type to avoid silent overflow.

### Type Declaration

Create a new type from an existing type using `type`. The new type has the same underlying structure but is treated as a distinct type.

```go
type WhatsappNumber string

var w1 WhatsappNumber = "08123456789"
var w2 string = "08123456789"

fmt.Println(w1)
fmt.Println(WhatsappNumber(w2))
fmt.Println("is equal? ", w1 == WhatsappNumber(w2)) // true
```

> **Note:** A type declaration creates a distinct type, not an alias. To convert between the original type and the declared type, use explicit conversion `T(value)`.
