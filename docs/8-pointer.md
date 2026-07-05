## Go Pointer

Pointer = **memory address** of a value. Instead of copying data, pass a reference to where the data lives.

### Pass by Value (Copy)

By default, Go copies data on assignment.

```go
address1 := Address{"Malang", "East Java", "Indonesia"}
address2 := address1      // copies the entire struct
address2.City = "Kediri"  // only affects address2

fmt.Println("address1 →", address1)   // {Malang East Java Indonesia}
fmt.Println("address2 →", address2)   // {Kediri East Java Indonesia}
```

`address1` and `address2` are independent — modifying one doesn't affect the other.

### Pass by Reference (`&` and `*`)

`&` = get pointer (memory address). `*` = dereference (get value at pointer).

```go
address3 := &address1        // address3 is *Address (pointer)
address3.City = "Surabaya"   // modifies address1 through pointer

fmt.Println("address1 →", address1)   // {Surabaya East Java Indonesia}
fmt.Println(address1 == *address3)    // true — same value
```

> **Note:** Go auto-dereferences struct fields — `address3.City` works same as `(*address3).City`.

| Operator | Name | Description |
|----------|------|-------------|
| `&value` | Address-of | Creates pointer to value |
| `*pointer` | Dereference | Gets value at pointer's address |

### `new()` Function

`new(Type)` creates a pointer to a **zero-valued** instance. Same as `&Type{}`.

```go
address4 := new(Address)
address5 := address4       // same pointer — no copy
address5.Country = "Indonesia"

fmt.Println(address4 == address5)   // true — same reference
```

> **Note:** `&Address{}` is more common and readable. `new()` is rarely used.

### Pointer in Function Parameters

Pointer param (`*Type`) lets the function modify the original value.

```go
func ChangeAddressToIndonesia(a *Address) {
    a.Country = "Indonesia"
}
```

```go
address6 := &Address{}
address6.Country = "Konoha"
ChangeAddressToIndonesia(address6)
fmt.Println(address6.Country)   // Indonesia

// Or pass address of a value variable:
address7 := Address{Country: "Konoha"}
ChangeAddressToIndonesia(&address7)
fmt.Println(address7.Country)   // Indonesia
```

| Approach | Declaration | Pass to Function | Modifies Original? |
|----------|-------------|------------------|--------------------|
| Pointer variable | `address6 := &Address{}` | `func(address6)` | ✅ Yes |
| Value + `&` | `address7 := Address{}` | `func(&address7)` | ✅ Yes |
| Value (no ptr) | `addr := Address{}` | `func(addr)` | ❌ No (copy) |

### Pointer Receiver Method

Pointer receiver (`*Type`) lets methods modify the struct directly.

```go
type Man struct {
    Name string
}

func (m *Man) Married() {
    m.Name = "Mr. " + m.Name
}

man := &Man{"Dewa"}
man.Married()
fmt.Println(man.Name)    // Mr. Dewa
```

> **Note:** Without `*Man`, `Married()` would modify a **copy** — changes lost.

### Value Receiver vs Pointer Receiver

| Aspect | Value (`m Man`) | Pointer (`m *Man`) |
|--------|----------------|--------------------|
| Original modified? | ❌ No — copy | ✅ Yes |
| Memory | Copies entire struct | Copies only pointer (8 bytes) |
| Use case | Read-only, small structs | Mutation, large structs |

### Key Operators

| Op | Example | Description |
|----|---------|-------------|
| `&` | `ptr := &value` | Address-of — creates pointer |
| `*` | `val := *ptr` | Dereference — gets value at pointer |
| `*Type` | `func(a *Address)` | Pointer type declaration |
| `new()` | `ptr := new(Address)` | Pointer to zero-valued instance |
