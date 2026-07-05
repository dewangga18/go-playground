## Go Fundamental

All the built-in data types Go has — numeric, boolean, string — plus how to declare variables and convert between types.

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

| Type | Size | Precision |
|------|------|-----------|
| `float32` | 32-bit | IEEE 754 single (~6 decimal digits) |
| `float64` | 64-bit | IEEE 754 double (~15 decimal digits) |

#### Complex Numbers

| Type | Size | Description |
|------|------|-------------|
| `complex64` | 64-bit | `float32` real + imaginary |
| `complex128` | 128-bit | `float64` real + imaginary |

#### Aliases

| Alias | Underlying | Notes |
|-------|------------|-------|
| `byte` | `uint8` | Represents a single byte |
| `rune` | `int32` | Represents a Unicode code point |
| `int` | `int32` or `int64` | Platform-dependent (at least 32-bit) |
| `uint` | `uint32` or `uint64` | Platform-dependent (at least 32-bit) |

> **Note:** `int`/`uint` size depends on architecture — 32-bit on 32-bit systems, 64-bit on 64-bit.

### Numeric Operations

#### Arithmetic Operators

| Op | Description | Example (`a=10, b=10, c=5, d=2, e=3`) |
|----|-------------|----------------------------------------|
| `+` | Addition | `c + d` → `7` |
| `-` | Subtraction | `d - e` → `-1` |
| `*` | Multiplication | `c * d` → `10` |
| `/` | Division | `a / b` → `1` |
| `%` | Modulo | `c % d` → `1` |

Standard precedence applies: `*`/`/`/`%` before `+`/`-`.

```go
const a = 10; const b = 10; const c = 5; const d = 2; const e = 3

var i = a/b + c*d - e   // 1 + 10 - 3 = 8
fmt.Println("result:", i)   // 8
```

#### Augmented Assignment

```go
i := 8
i += 5    // i = i + 5  → 13
i -= 5    // i = i - 5  → 8
i *= 5    // i = i * 5  → 40
i /= 5    // i = i / 5  → 8
i %= 5    // i = i % 5  → 3
```

> **Note:** `%` and `%=` only work with integer types.

#### Unary Operators

| Op | Name | Example |
|----|------|---------|
| `-` | Negation | `-a` |
| `++` | Increment | `a++` |
| `--` | Decrement | `b--` |

```go
var a = 10; var b = -10
fmt.Println("a =", a)    // 10
a++
fmt.Println("a++ =", a)  // 11
b--
fmt.Println("b-- =", b)  // -11
```

> **Note:** `++` and `--` are **statements** in Go, not expressions. Can't use inside other expressions like `c = a++ + b`.

### Boolean Types

```go
var trueConstant bool = true
var falseConstant bool = false

fmt.Println(trueConstant)   // true
fmt.Println(falseConstant)  // false
fmt.Println(1 == 1)         // true
fmt.Println(1 != 1)         // false
```

#### Comparison Operators

| Op | Description | Example (`a=1, b=2`) |
|----|-------------|----------------------|
| `>` | Greater than | `a > b` → `false` |
| `<` | Less than | `a < b` → `true` |
| `>=` | Greater or equal | `a >= b` → `false` |
| `<=` | Less or equal | `a <= b` → `true` |
| `==` | Equal | `a == b` → `false` |
| `!=` | Not equal | `a != b` → `true` |

Works with numeric types and strings. `<`, `>`, `<=`, `>=` only work with ordered types.

#### Logical Operators

| Op | Name | Description | Example |
|----|------|-------------|---------|
| `&&` | AND | `true` if **both** are `true` | `true && false` → `false` |
| `\|\|` | OR | `true` if **at least one** is `true` | `true \|\| false` → `true` |
| `!` | NOT | Inverts boolean | `!true` → `false` |

Short-circuit evaluation: `&&` stops if left is `false`, `||` stops if left is `true`.

### String Types

Strings are sequences of bytes. **Immutable** in Go.

```go
var str1 string = "Hello"

fmt.Println(str1)               // Hello
fmt.Println(len(str1))          // 5
fmt.Println(str1[0])            // 72 (byte value, not character)
```

> **Note:** Indexing a string returns a **byte** (`uint8`), not a character. For Unicode characters, use `rune`.

### Variable Declaration

#### `var` — explicit

```go
var a int8 = 127
var name = "Go"  // type inferred
```

#### `const` — immutable

```go
const e = 42        // untyped constant
const pi = 3.14159
```

#### `:=` — short declaration

Shorthand for `var` inside functions. Type inferred.

```go
count := 10
tmp := tempA
```

#### Multiple declaration

```go
var (
    firstName = "Aaron"
    lastName  = "Evanjulio"
    version int = 1
)
```

#### Swap pattern

```go
tmp := tempA
tempA = tempB
tempB = tmp
```

### Type Conversion

Go doesn't support implicit conversion — must be explicit with `T(value)`.

#### Byte to String

```go
var firstName = "Aaron"
var byteVal uint8 = firstName[0]
var byteToStr = string(byteVal)

fmt.Println("Value of byteVal from firstName[0]:", byteVal)
fmt.Println("Convert byte to string:", byteToStr)
```

> **Note:** `string(byteValue)` treats the byte as a Unicode code point. Works for ASCII because byte value = code point.

#### Integer Conversion

```go
var val32 int32 = 32769
var val64 int64 = int64(val32)    // 32769 — safe
var val16 int16 = int16(val32)    // overflow! 32769 > int16 max (32767)
```

| Conversion | Result | Notes |
|-----------|--------|-------|
| `int32` → `int64` | Safe | Wider type fits all values |
| `int32` → `int16` | Overflow | Narrower type truncates bits |

> **Note:** Always check target type's range before converting to narrower type — silent overflow.

### Type Declaration

Create a new distinct type from an existing type using `type`. Not an alias — explicit conversion required.

```go
type WhatsappNumber string

var w1 WhatsappNumber = "08123456789"
var w2 string = "08123456789"

fmt.Println(w1)
fmt.Println(WhatsappNumber(w2))
fmt.Println("is equal? ", w1 == WhatsappNumber(w2)) // true
```
