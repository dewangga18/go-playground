## Go Struct

A `struct` is a **collection of fields** that groups related data together. It's Go's way of creating custom data types.

### Struct Definition

Define a struct using the `type` keyword followed by the struct name and fields.

```go
type Customer struct {
    name, city string
    age        int
}
```

| Component | Description |
|-----------|-------------|
| `type Customer struct` | Declares a new struct type named `Customer` |
| `name, city string` | Two string fields (same-type shorthand) |
| `age int` | Integer field |

> **Note:** Fields with the same type can be declared on one line (`name, city string`) instead of separate lines (`name string; city string`).

### Struct Literal

Create a struct value using a **struct literal**. There are two forms: **named fields** and **positional**.

#### Named Fields

Specify field names explicitly. Order doesn't matter.

```go
customer1 := Customer{
    name: "Aaron Evanjulio",
    city: "Malang",
    age:  22,
}
```

```
=== Struct Example ===
{Aaron Evanjulio Malang 22}
Aaron Evanjulio
```

#### Positional Fields

Omit field names and provide values **in declaration order**. Less verbose but less explicit.

```go
customer2 := Customer{"Evanjulio Dewangga", "Kediri", 22}
```

```
=== Struct Example ===
{Evanjulio Dewangga Kediri 22}
Evanjulio Dewangga
```

| Form | Example | Pros | Cons |
|------|---------|------|------|
| **Named** | `Customer{name: "Aaron", age: 22}` | Self-documenting, order-independent | More verbose |
| **Positional** | `Customer{"Aaron", "Malang", 22}` | Concise | Fragile — breaks if field order changes |

### Access Fields

Access individual fields using dot notation.

```go
fmt.Println(customer1.name)   // Aaron Evanjulio
fmt.Println(customer2.name)   // Evanjulio Dewangga
```

### Struct Method

Go allows you to define methods on structs using a **receiver**. A receiver is like a parameter that binds the method to the type.

```go
func (cust Customer) sayHello() {
    fmt.Println("Hi,", cust.name)
}
```

```
=== Struct Method Example ===
Hi, Aaron Evanjulio
Hi, Evanjulio Dewangga
```

| Component | Description |
|-----------|-------------|
| `(cust Customer)` | **Receiver** — binds method to `Customer` type |
| `sayHello()` | Method name |
| `cust.name` | Accesses the `name` field via the receiver |

```go
customer1.sayHello()   // Hi, Aaron Evanjulio
customer2.sayHello()   // Hi, Evanjulio Dewangga
```

> **Note:** Methods are just functions with a receiver. Unlike classes in OOP languages, Go structs don't support inheritance — use **composition** instead.
