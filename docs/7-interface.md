## Go Interface

Interface = set of **method signatures**. Go uses **implicit implementation** — no `implements` keyword needed.

### Interface Definition & Implementation

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

`Person` automatically satisfies `HasName` — just needs `GetName() string` method.

---

### Interface as Parameter

Function accepting an interface works with **any type** that satisfies it.

```go
func PrintName(h HasName) {
    fmt.Println(h.GetName())
}

var person HasName
person = Person{"Aaron Evanjulio"}
PrintName(person)
```

Output:

```
=== Interface Example ===
Aaron Evanjulio
```

| Without Interface | With Interface |
|-----------------|----------------|
| Separate functions for each struct | One function for many structs |

---

### Empty Interface / `any`

`interface{}` (old) / `any` (Go 1.18+) — holds **any type**.

```go
func Ups() any {
    return "Ups new version"
}
```

> **Note:** `any` is an alias for `interface{}`. `any` is preferred in modern Go.

---

### Nil

`nil` = zero value for certain types. Means "no value."

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

Output:

```
=== Nil Example ===
map[]
Value is empty
Aaron
```

`nil` only applies to these types:

| Can be nil | Can NOT be nil |
|------------|---------------|
| Pointer, Slice, Map, Interface, Channel, Function | int, string, bool, float, etc. (zero value instead) |

---

### Type Assertion

Extract concrete value from an interface.

```go
func random() any {
    return "OK"
}

result := random()
resultString := result.(string)
fmt.Println(resultString)   // OK
```

> **Note:** Wrong type assertion **panics**. Use type switch for safety.

#### Type Switch (Safer)

```go
switch value := result.(type) {
case string:
    fmt.Println("String", value)
case int:
    fmt.Println("Int", value)
default:
    fmt.Println("Unknown", value)
}
```

| Approach | Safe? | Use Case |
|----------|-------|----------|
| Simple assertion `result.(string)` | ❌ Panics on mismatch | When type is guaranteed |
| Type switch `switch v := result.(type)` | ✅ No panic | When type is unknown |
