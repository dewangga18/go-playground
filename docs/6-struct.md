## Go Struct

Custom data type that groups related fields together. Can also have methods via receivers.

### Struct Definition

```go
type Customer struct {
    name, city string
    age        int
}
```

> **Note:** Same-type fields on one line: `name, city string`.

### Struct Literal

#### Named Fields

```go
customer1 := Customer{
    name: "Aaron Evanjulio",
    city: "Malang",
    age:  22,
}
```

Output:

```
=== Struct Example ===
{Aaron Evanjulio Malang 22}
Aaron Evanjulio
```

#### Positional Fields

Values in declaration order — less verbose, but fragile.

```go
customer2 := Customer{"Evanjulio Dewangga", "Kediri", 22}
```

| Form | Example | Pros | Cons |
|------|---------|------|------|
| **Named** | `Customer{name: "Aaron", age: 22}` | Self-documenting, order-independent | More verbose |
| **Positional** | `Customer{"Aaron", "Malang", 22}` | Concise | Fragile — breaks if field order changes |

### Access Fields

```go
fmt.Println(customer1.name)   // Aaron Evanjulio
```

### Struct Method

Method on struct via **receiver**.

```go
func (cust Customer) sayHello() {
    fmt.Println("Hi,", cust.name)
}

customer1.sayHello()   // Hi, Aaron Evanjulio
customer2.sayHello()   // Hi, Evanjulio Dewangga
```

| Component | Description |
|-----------|-------------|
| `(cust Customer)` | **Receiver** — binds method to `Customer` |
| `sayHello()` | Method name |
| `cust.name` | Access field via receiver |

> **Note:** Methods are just functions with a receiver. No inheritance — use **composition** instead.
