## Go Collections

Three collection types in Go: fixed-size arrays, dynamic slices, and key-value maps.

### Array

Fixed-size sequence of elements, same type. Declared with `[n]T`.

```go
var names [3]string
names[0] = "Aaron"
names[1] = "Evanjulio"
names[2] = "Dewangga"

fmt.Println(names)            // [Aaron Evanjulio Dewangga]
fmt.Println(names[0])         // Aaron
```

#### Literal Initialization

Remaining elements get the **zero value** of the type.

```go
var values = [4]int{1, 2, 3}

fmt.Println(values)                    // [1 2 3 0]
fmt.Println("Length:", len(values))    // 4
```

#### Compiler-calculated Size (`...`)

```go
var computedCapacity = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

fmt.Println(computedCapacity)                    // [1 2 3 4 5 6 7 8 9 10]
fmt.Println("Capacity:", cap(computedCapacity))  // 10
```

| Syntax | Size |
|--------|------|
| `[n]T{}` | Explicit — set by `n` |
| `[...]T{}` | Inferred — counted by compiler |

### Slice

Dynamic view into an underlying array. Can grow and shrink.

#### Create from Array

`array[low:high]` creates a slice referencing a portion of the array.

```go
var days = [...]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

sliceList := days[1:4]
fmt.Println(sliceList)                       // [Tuesday Wednesday Thursday]
fmt.Println("Length:", len(sliceList))       // 3
fmt.Println("Capacity:", cap(sliceList))     // 6
```

| Component | Value | Explanation |
|-----------|-------|-------------|
| **Pointer** | `1` | Starting index (low) |
| **Length** | `3` | Elements in slice (`high - low`) |
| **Capacity** | `6` | Remaining from start to end of array (`len(array) - low`) |

#### Shorthand

| Syntax | Equivalent | Result |
|--------|-----------|--------|
| `s[low:]` | `s[low:len(s)]` | From `low` to end |
| `s[:high]` | `s[0:high]` | From start to `high` |
| `s[:]` | `s[0:len(s)]` | Entire array/slice |

#### Slice is a Reference

Modifying a slice **changes the underlying array**.

```go
colors := [...]string{"Orange", "Yellow", "Brown", "Green", "Purple", "Blue"}
tmpColors1 := colors[3:]

tmpColors1[0] = "Lime"

fmt.Println(tmpColors1)   // [Lime Purple Blue]
fmt.Println(colors)       // [Orange Yellow Brown Lime Purple Blue] ← original changed!
```

#### Append

`append()` adds elements. If capacity is exceeded, Go allocates a **new underlying array**.

```go
colors := [...]string{"Orange", "Yellow", "Brown", "Green", "Purple", "Blue"}
tmpColors1 := colors[3:]   // len=3, cap=3
tmpColors1[0] = "Lime"

// Append exceeds capacity → new array
tmpColors2 := append(tmpColors1, "White", "Black")
fmt.Println(tmpColors2)   // [Lime Purple Blue White Black]
fmt.Println(colors)       // [Orange Yellow Brown Lime Purple Blue] ← unchanged
```

After capacity exceeded, new slice is independent:

```go
tmpColors2[2] = "Cyan"
fmt.Println(tmpColors2)   // [Lime Purple Cyan White Black]
fmt.Println(colors)       // [Orange Yellow Brown Lime Purple Blue] ← unchanged
```

> **Note:** When `append()` exceeds capacity, Go typically doubles the capacity. Plan initial capacity with `make()` if size is known in advance.

#### Make Function

`make([]T, length, capacity)` — create a slice with specified length and capacity.

```go
newColors := make([]string, 2, 10)
newColors[0] = "Maroon"
newColors[1] = "Magenta"

fmt.Println(newColors)                  // [Maroon Magenta]
fmt.Println("Length:", len(newColors))  // 2
fmt.Println("Capacity:", cap(newColors)) // 10
```

Append uses remaining capacity without reallocation:

```go
newColors2 := append(newColors, "Navy")
fmt.Println(newColors2)                     // [Maroon Magenta Navy]
fmt.Println("Capacity:", cap(newColors2))   // 10 ← same array
```

### Array vs Slice

| Aspect | Array | Slice |
|--------|-------|-------|
| **Syntax** | `[n]T` — size in brackets | `[]T` — no size |
| **Size** | Fixed at compile time | Dynamic |
| **Append** | ❌ Not supported | ✅ `append()` |
| **Underlying** | Stores data directly | References an array |

```go
thisArray := [...]int{1, 2, 3}
thisSlice := []int{1, 2, 3}

fmt.Println(thisArray)   // [1 2 3]
fmt.Println(thisSlice)   // [1 2 3]
```

```go
thisSlice = append(thisSlice, 4)
fmt.Println(thisSlice)   // [1 2 3 4], len=4, cap=6

thisSlice = append(thisSlice, 5, 6, 7, 8)
fmt.Println(thisSlice)   // [1 2 3 4 5 6 7 8], len=8, cap=12
```

### Map

Unordered key-value collection. Key must be comparable.

#### Literal Declaration

```go
person := map[string]string{
    "name": "Aaron",
    "age":  "22",
    "city": "Malang",
}

fmt.Println(person, "Length:", len(person))   // map[age:22 city:Malang name:Aaron] Length: 3
```

#### Delete Key

```go
delete(person, "age")
fmt.Println("After delete:", person)   // map[city:Malang name:Aaron]
```

#### Non-existent Key

Returns the **zero value** of the value type — no error.

```go
wrongKey := person["jobs"]
fmt.Println("Call wrong key:", wrongKey)   // "" (empty string)
```

#### Create with `make()`

```go
device := make(map[string]any)
device["name"] = "iQOO"
device["os"] = "android"
device["ram"] = 8
device["rom"] = 256

fmt.Println(device)   // map[name:iQOO os:android ram:8 rom:256]
```

> **Note:** Map is a reference type. Assigning to another variable shares the same underlying data. Also **unordered** — iteration order not guaranteed.
