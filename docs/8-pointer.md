## Go Pointer

A **pointer** holds the **memory address** of a value, rather than the value itself. Instead of passing data around, you can pass a reference to where the data lives — allowing functions and methods to modify the original value.

---

### Pass by Value (Copy)

By default, Go uses **pass by value** — assigning one variable to another creates a full copy of the data.

```go
address1 := Address{"Malang", "East Java", "Indonesia"}
address2 := address1      // copies the entire struct
address2.City = "Kediri"  // only affects address2

fmt.Println("address1 →", address1)   // {Malang East Java Indonesia}
fmt.Println("address2 →", address2)   // {Kediri East Java Indonesia}
```

```
== Pointer Example ==

== Pass by Value ==
address1 →  {Malang East Java Indonesia}
address2 →  {Kediri East Java Indonesia}
```

`address1` and `address2` are completely independent — modifying one doesn't affect the other.

---

### Pass by Reference (& and *)

Use `&` to get a **pointer** (memory address) of a variable. Use `*` to **dereference** — access the value the pointer points to.

```go
address3 := &address1        // address3 is a *Address (pointer)
address3.City = "Surabaya"   // modifies address1 through the pointer

fmt.Println("address1 →", address1)   // {Surabaya East Java Indonesia}
fmt.Println("address3 →", address3)   // &{Surabaya East Java Indonesia}
fmt.Println(address1 == *address3)    // true — same underlying value
```

```
== Pass by Reference ==
address1 →  {Surabaya East Java Indonesia}
address3 →  &{Surabaya East Java Indonesia}
true
false
```

> **Note:** In Go, struct fields can be accessed through a pointer **without explicit dereferencing** — `address3.City` works the same as `(*address3).City`. This is called **automatic dereferencing**.

| Operator | Name | Description |
|----------|------|-------------|
| `&value` | Address-of | Creates a pointer to the value |
| `*pointer` | Dereference | Gets the value at the pointer's address |

---

### `new()` Function

`new(Type)` creates a pointer to a **zero-valued** instance of the type. It's equivalent to `&Type{}`.

```go
address4 := new(Address)   // *Address pointing to Address{"", "", ""}
address5 := address4       // same pointer — no copy
address5.Country = "Indonesia"

fmt.Println(address4)      // &{  Indonesia}
fmt.Println(address5)      // &{  Indonesia}
fmt.Println(address4 == address5)   // true — same reference
```

```
== New Function ==
&{  Indonesia}
&{  Indonesia}
true
```

Since `address5` points to the same memory as `address4`, changing one changes both.

> **Note:** `new(Address)` and `&Address{}` are functionally identical. `&Address{}` is more common and readable — `new()` is rarely used in practice.

---

### Pointer in Function Parameters

When a function takes a pointer parameter (`*Type`), it can modify the original value — the function receives a reference, not a copy.

```go
func ChangeAddressToIndonesia(a *Address) {
    a.Country = "Indonesia"
}
```

Passing with a pointer:

```go
address6 := &Address{}
address6.Country = "Konoha"
ChangeAddressToIndonesia(address6)

fmt.Println("address6 →", address6)        // &{ Indonesia}
fmt.Println(address6.Country)              // Indonesia
```

Passing a value as a pointer using `&`:

```go
address7 := Address{}          // value, not pointer
address7.Country = "Konoha"
ChangeAddressToIndonesia2(&address7)   // pass address with &

fmt.Println("address7 →", address7)   // { Indonesia}
```

```
== Function in Pointer ==
address6 before →  &{ Konoha}
address6 after →  &{ Indonesia}
Indonesia
address7 before →  { Konoha}
address7 after →  { Indonesia}
```

| Approach | Declaration | Pass to Function | Modifies Original? |
|----------|-------------|------------------|--------------------|
| Pointer variable | `address6 := &Address{}` | `func(address6)` | ✅ Yes |
| Value variable | `address7 := Address{}` | `func(&address7)` | ✅ Yes (pass address) |
| Value (no pointer) | `addr := Address{}` | `func(addr)` | ❌ No (gets copy) |

---

### Pointer Receiver Method

Methods can also use **pointer receivers** (`*Type`), allowing them to modify the struct directly.

```go
type Man struct {
    Name string
}

func (m *Man) Married() {
    m.Name = "Mr. " + m.Name
}
```

Usage:

```go
man := &Man{"Dewa"}
fmt.Println("man before →", man.Name)   // Dewa

man.Married()
fmt.Println("man after →", man.Name)    // Mr. Dewa
```

```
== Method in Pointer ==
man before →  Dewa
man after →  Mr. Dewa
```

> **Note:** Without a pointer receiver (`*Man`), the method would modify a **copy** of the struct, and the change would be lost. Use pointer receivers when a method needs to mutate the receiver.

---

### Value Receiver vs Pointer Receiver

| Aspect | Value Receiver (`m Man`) | Pointer Receiver (`m *Man`) |
|--------|-------------------------|----------------------------|
| Original modified? | ❌ No — works on copy | ✅ Yes — works on original |
| Memory | Copies entire struct | Copies only pointer (8 bytes) |
| Use case | Read-only methods | Methods that mutate or handle large structs |

> **Best Practice:** If in doubt, use a **pointer receiver**. Most methods in Go that modify state use pointer receivers. For read-only methods on small structs, value receivers are fine.

---

### Key Operators Reference

| Operator | Example | Description |
|----------|---------|-------------|
| `&` | `ptr := &value` | Creates a pointer to `value` |
| `*` | `val := *ptr` | Dereferences — gets the value at the pointer |
| `*Type` | `func(a *Address)` | Declares a pointer type |
| `new()` | `ptr := new(Address)` | Creates a pointer to a zero-valued instance |
