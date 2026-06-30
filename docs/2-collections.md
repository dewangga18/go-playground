## Go Collections

Go has built-in collection types for grouping multiple values. Unlike many languages, Go separates fixed-size arrays from dynamic slices.

### Array

An array is a **fixed-size** sequence of elements of the same type.

#### Declaration & Assignment

Declare with `[n]T` where `n` is the size and `T` is the element type. Assign values by index.

```go
var names [3]string
names[0] = "Aaron"
names[1] = "Evanjulio"
names[2] = "Dewangga"

fmt.Println(names)            // [Aaron Evanjulio Dewangga]
fmt.Println(names[0])         // Aaron
```

#### Literal Initialization

Initialize directly using a literal. Remaining elements get the **zero value** of the type.

```go
var values = [4]int{1, 2, 3}

fmt.Println(values)              // [1 2 3 0]
fmt.Println("Length:", len(values))     // 4
fmt.Println("Capacity:", cap(values))   // 4
```

> **Note:** `len()` returns the number of elements, `cap()` returns the capacity. For arrays, both are equal to the declared size.

#### Compiler-calculated Size (`...`)

Use `...` to let the compiler count the number of elements in the literal.

```go
var computedCapacity = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

fmt.Println(computedCapacity)               // [1 2 3 4 5 6 7 8 9 10]
fmt.Println("Capacity:", cap(computedCapacity))  // 10
```

| Syntax | Example | Size |
| --- | --- | --- |
| `[n]T{}` | `[4]int{1, 2, 3}` | Explicit — set by `n` |
| `[...]T{}` | `[...]int{1, 2, 3, 4, 5}` | Inferred — counted by compiler |

### Slice

A slice is a **dynamic** view into an underlying array. Unlike arrays, slices can grow and shrink.

#### Create Slice from Array

Use `array[low:high]` to create a slice. The slice references a portion of the array.

```go
var days = [...]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

sliceList := days[1:4]
fmt.Println(sliceList)                     // [Tuesday Wednesday Thursday]
fmt.Println("Length:", len(sliceList))     // 3  (high - low)
fmt.Println("Capacity:", cap(sliceList))   // 6  (len(array) - low)
```

| Component | Value | Explanation |
| --- | --- | --- |
| **Pointer** | `1` | Starting index in the array (low) |
| **Length** | `3` | Number of elements in the slice (`high - low`) |
| **Capacity** | `6` | Remaining elements from start to end of array (`len(array) - low`) |

#### Slice Shorthand

| Syntax | Equivalent To | Result |
| --- | --- | --- |
| `s[low:]` | `s[low:len(s)]` | From `low` to the end |
| `s[:high]` | `s[0:high]` | From the start to `high` |
| `s[:]` | `s[0:len(s)]` | The entire array/slice |

```go
days := [...]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

tmp1 := days[4:]   // [Friday Saturday Sunday]
tmp2 := days[:4]   // [Monday Tuesday Wednesday Thursday]
tmp3 := days[:]    // [Monday Tuesday Wednesday Thursday Friday Saturday Sunday]
```

#### Slice is a Reference

Modifying a slice **changes the underlying array**. This is because a slice is a view, not a copy.

```go
colors := [...]string{"Orange", "Yellow", "Brown", "Green", "Purple", "Blue"}
tmpColors1 := colors[3:]

tmpColors1[0] = "Lime"

fmt.Println(tmpColors1)   // [Lime Purple Blue]
fmt.Println(colors)       // [Orange Yellow Brown Lime Purple Blue] ← original changed!
```

#### Append

Use `append()` to add elements. If the slice has enough capacity, it modifies the same underlying array. If capacity is exceeded, `append()` creates a **new underlying array**.

```go
colors := [...]string{"Orange", "Yellow", "Brown", "Green", "Purple", "Blue"}
tmpColors1 := colors[3:]    // [Green Purple Blue], len=3, cap=3
tmpColors1[0] = "Lime"     // same modification from section above

// Append exceeds capacity → creates new array
tmpColors2 := append(tmpColors1, "White", "Black")
fmt.Println(tmpColors2)   // [Lime Purple Blue White Black]
fmt.Println(colors)       // [Orange Yellow Brown Lime Purple Blue] ← unchanged
```

```go
// After capacity exceeded, the new slice is independent
tmpColors2[2] = "Cyan"
fmt.Println(tmpColors2)   // [Lime Purple Cyan White Black]
fmt.Println(colors)       // [Orange Yellow Brown Lime Purple Blue] ← unchanged
```

> **Note:** When `append()` exceeds capacity, Go allocates a new underlying array — typically doubling the capacity. The old array is unaffected.

#### Make Function

`make([]T, length, capacity)` creates a slice with a specified length and capacity.

```go
newColors := make([]string, 2, 10)
newColors[0] = "Maroon"
newColors[1] = "Magenta"

fmt.Println(newColors)              // [Maroon Magenta]
fmt.Println("Length:", len(newColors))     // 2
fmt.Println("Capacity:", cap(newColors))   // 10
```

Appending to a slice created with `make()` uses the remaining capacity without allocating a new array.

```go
newColors2 := append(newColors, "Navy")
fmt.Println(newColors2)               // [Maroon Magenta Navy]
fmt.Println("Length:", len(newColors2))     // 3
fmt.Println("Capacity:", cap(newColors2))   // 10 ← same array, capacity not exceeded
```

### Array vs Slice

| Aspect | Array | Slice |
| --- | --- | --- |
| **Syntax** | `[n]T` — size in brackets | `[]T` — no size |
| **Size** | Fixed at compile time | Dynamic |
| **Append** | ❌ Not supported | ✅ `append()` |
| **Underlying** | Stores data directly | References an underlying array |

```go
thisArray := [...]int{1, 2, 3}
thisSlice := []int{1, 2, 3}

fmt.Println(thisArray)   // [1 2 3]
fmt.Println(thisSlice)   // [1 2 3]
```

```go
// thisArray = append(thisArray, 4)  // ❌ compiler error — arrays can't grow
thisSlice = append(thisSlice, 4)

fmt.Println(thisSlice)   // [1 2 3 4], length=4, capacity=6 (doubled)
```

Appending multiple elements at once:

```go
thisSlice = append(thisSlice, 5, 6, 7, 8)
fmt.Println(thisSlice)   // [1 2 3 4 5 6 7 8], length=8, capacity=12
```

> **Note:** When a slice's capacity is full, `append()` doubles the capacity. This means `append()` may be expensive for large slices — plan your initial capacity with `make()` if you know the size in advance.

### Map

A map is an **unordered** collection of key-value pairs. The key type must be comparable (`==`, `!=`), and all keys must be the same type.

#### Literal Declaration

Declare with `map[K]V{...}` where `K` is the key type and `V` is the value type.

```go
person := map[string]string{
    "name": "Aaron",
    "age":  "22",
    "city": "Malang",
}

fmt.Println(person, "Length:", len(person))   // map[age:22 city:Malang name:Aaron] Length: 3
fmt.Println("Name:", person["name"])          // Aaron
fmt.Println("Age:", person["age"])            // 22
fmt.Println("City:", person["city"])          // Malang
```

#### Delete Key

Use `delete(map, key)` to remove a key-value pair.

```go
delete(person, "age")
fmt.Println("After delete:", person)   // map[city:Malang name:Aaron]
```

#### Accessing Non-existent Key

Accessing a key that doesn't exist returns the **zero value** of the value type — no error.

```go
wrongKey := person["jobs"]
fmt.Println("Call wrong key:", wrongKey)   // "" (empty string)
```

#### Create with `make()`

`make(map[K]V)` creates an empty map. Use `any` as the value type to store mixed types.

```go
device := make(map[string]any)
device["name"] = "iQOO"
device["os"] = "android"
device["ram"] = 8
device["rom"] = 256

fmt.Println(device)   // map[name:iQOO os:android ram:8 rom:256]
```

| Key Type | Value Type | Description |
| --- | --- | --- |
| `string` | `string` | Simple string-to-string map (`person`) |
| `string` | `any` | Mixed value types (`device`) |

> **Note:** Map is a reference type. Assigning a map to another variable shares the same underlying data. Maps are also **unordered** — iteration order is not guaranteed.