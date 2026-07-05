## Go Interface

An **interface** in Go is a set of **method signatures**. Unlike other languages, Go uses **implicit implementation** — a struct automatically satisfies an interface if it has all the required methods, without needing an `implements` keyword.

---

### Interface Definition & Implementation

Define an interface using the `type` and `interface` keywords. A struct implements it implicitly by defining the matching methods.

```go
type HasName interface {
    GetName() string
}

type Person struct {
    name string
}

func (p Person) GetName() string {
    return p.name
}
```

A variable of the interface type can hold any value whose type implements that interface:

```go
var person HasName
person = Person{"Aaron Evanjulio"}
PrintName(person)
```

Output:
```
=== Interface Example ===
Aaron Evanjulio
```

> **Note:** Implementation is **implicit** — no `implements` keyword is needed. Any struct that has all the methods with matching signatures automatically satisfies the interface.

---

### Interface as Parameter

Interfaces can be used as parameter types. This allows a function to accept **any type** that satisfies the interface.

```go
func PrintName(h HasName) {
    fmt.Println(h.GetName())
}
```

Usage:

```go
PrintName(Person{"Aaron Evanjulio"})   // Aaron Evanjulio
```

| Without Interface | With Interface |
|-----------------|----------------|
| Must create separate functions for each struct | One function can accept many different structs |
| Less flexible | More modular and extensible |

---

### Empty Interface / `any`

The **empty interface** (`interface{}`) can hold **a value of any type**. Since Go 1.18, the alias `any` is available — shorter and more idiomatic.

```go
// Old version (still valid)
func Upss() interface{} {
    return "Upss old version"
}

// New version (Go 1.18+)
func Ups() any {
    return "Ups new version"
}
```

Usage:

```go
var emptyOld = Upss()
fmt.Println(emptyOld)

var emptyNew = Ups()
fmt.Println(emptyNew)
```

Output:
```
=== Empty Interface Example ===
Upss old version
Ups new version
```

> **Note:** `any` is an alias for `interface{}`. They are technically identical, but `any` is shorter and is the standard in modern Go code.

---

### Nil

`nil` is the **zero value** for certain types — it means "no value" or "empty reference." A function can return `nil` to signal that no valid value exists.

```go
func newExample(name string) map[string]string {
    if name == "" {
        return nil
    }
    return map[string]string{
        "name": name,
    }
}
```

When name is empty, `newExample("")` returns `nil`. Printing a nil map shows `map[]` — it has zero length but doesn't crash:

```go
example := newExample("")
fmt.Println(example)            // map[]

val := example["name"]
if val == "" {
    fmt.Println("Value is empty")   // output
}
```

When name is provided, it returns a normal map:

```go
example = newExample("Aaron")
fmt.Println(example["name"])    // Aaron
```

Output:
```
=== Nil Example ===
map[]
Value is empty
Aaron
```

> **Note:** `nil` is **not valid for all types**. It only applies to these types:
> | Type | Can be nil | Example |
> |------|------------|---------|
> | Pointer | ✅ | `var p *int = nil` |
> | Slice | ✅ | `var s []int = nil` |
> | Map | ✅ | `var m map[string]int = nil` |
> | Interface | ✅ | `var i HasName = nil` |
> | Channel | ✅ | `var ch chan int = nil` |
> | Function | ✅ | `var fn func() = nil` |
> |
> | Primitive (int, string, bool, float, etc.) | ❌ | Cannot be nil — use zero value instead |

---

### Type Assertion

**Type assertion** extracts the concrete value from an interface. Use `value.(Type)` to assert that an interface holds a specific type.

```go
func random() any {
    return "OK"
}

result := random()
resultString := result.(string)
fmt.Println(resultString)   // OK
```

```
=== Type Assertion Example ===
OK
```

> **Note:** If the assertion is wrong (e.g., asserting `int` on a `string` value), the program **panics**. Always be certain of the underlying type before asserting.

#### Type Switch (Safer)

A **type switch** is the safe alternative. It uses `switch value := result.(type)` to match against multiple possible types without panicking.

```go
switch value := result.(type) {
case string:
    fmt.Println("String", value)   // String OK
case int:
    fmt.Println("Int", value)
default:
    fmt.Println("Unknown", value)
}
```

| Approach | Syntax | Safe? | Use Case |
|----------|--------|-------|----------|
| Simple assertion | `result.(string)` | ❌ Panics on mismatch | When type is guaranteed |
| Type switch | `switch v := result.(type)` | ✅ No panic | When type is unknown |

> **Best Practice:** Use a **type switch** when you're not absolutely sure of the underlying type. Reserve simple assertions for situations where a wrong type should crash the program.

---

### When to Use What

| Concept | When to Use | Example |
|---------|-------------|---------|
| **Interface** | Need a method contract that multiple types must satisfy | `HasName`, `Reader`, `Writer` |
| **Empty Interface / `any`** | Need to hold a value of unknown type | Simple generic-like functions, placeholders |
| **Interface as Parameter** | Want a function to accept many different types | `PrintName(h HasName)` |

> **Best Practice:** Prefer specific interfaces (with methods) over `any` whenever possible. `any` removes type safety — errors only surface at runtime.
