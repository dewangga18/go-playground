## Go Functions

Reusable blocks of code — params, returns, variadic, first-class functions, closures, and recursion.

### Basic Function

No parameters, no return value.

```go
func basicFunction() {
    fmt.Println("Hi!")
}
```

Output:

```
=== Basic Function ===
Hi!
```

### Function Parameter

```go
func functionParameter(firstName string, lastName string) {
    fmt.Println("Hello", firstName, lastName)
}
```

> **Note:** Same-type params can be shortened: `firstName, lastName string`.

### Return Value

Return type after parameter list.

```go
func returnValue(width int, height int) int {
    return width * height
}

area := returnValue(10, 5)
fmt.Println("Area is", area)   // 50
```

Output:

```
=== Return Value Function ===
Area is 50
```

### Multiple Return Values

```go
func returnValue2(width int, height int) (int, string) {
    area := width * height
    if area >= 10 {
        return area, "Large"
    }
    return area, "Small"
}

area, size := returnValue2(20, 15)
fmt.Println("Area is", area, "and size is", size)   // Area is 300 and size is Large

area2, _ := returnValue2(15, 20)
fmt.Println("Area is", area2)                         // Area is 300
```

> **Note:** `_` (blank identifier) to ignore return values. Go won't compile if a declared variable is unused.

### Named Return Values

Return values can be named — they act as local variables.

```go
func namedReturnValue(width int, height int) (area int, size string) {
    area = width * height
    if area >= 10 {
        size = "Large"
    }
    return area, size
}
```

| Feature | Description |
|---------|-------------|
| `(area int, size string)` | Named returns — declared as local variables |
| `area = width * height` | Assign directly without `:=` |
| `return area, size` | Explicit return |

> **Note:** Can use **naked return** (`return` without values) with named returns. Use sparingly — explicit returns are clearer.

### Variadic Function

Variable number of arguments using `...`.

```go
func variadicFunction(numbers ...int) (total int) {
    for _, number := range numbers {
        total += number
    }
    average := total / len(numbers)
    return average
}

average := variadicFunction(10, 20, 30, 40, 50)
fmt.Println("Average is", average)   // 30

// Spread a slice
sliceWithVariadic := []int{10, 20, 30, 40, 50}
average = variadicFunction(sliceWithVariadic...)
fmt.Println("Average is", average)   // 30
```

> **Note:** Variadic param must be the **last** parameter.

### Function As Value

Functions are first-class — assign to variables.

```go
func functionAsValue(name string) string {
    return "Good bye, " + name + "!"
}

goodBye := functionAsValue
fmt.Println(goodBye("Aaron"))   // Good bye, Aaron!
```

### Function As Parameter

Pass functions as arguments (callback pattern).

```go
func filterOddNumber(number int) string {
    if number%2 == 1 {
        return strconv.Itoa(number) + " is Odd"
    }
    return strconv.Itoa(number) + " is Even"
}

func functionAsParams(number int, filter func(int) string) {
    fmt.Println(filter(number))
}

functionAsParams(11, filterOddNumber)   // 11 is Odd
```

Type alias for cleaner signatures:

```go
type Filter func(int) string

func functionAsParams2(number int, filter Filter) {
    fmt.Println(filter(number))
}

functionAsParams2(13, filterOddNumber)  // 13 is Odd
```

> **Note:** `strconv.Itoa()` converts `int` to string. Itoa = "Integer to ASCII".

### Anonymous Function

Function without a name.

#### Assigned to Variable

```go
greet := func(name string) {
    fmt.Println("Hello,", name)
}

greet("John")   // Hello, John
```

#### With Return Value

```go
square := func(number int) int {
    return number * number
}

result := square(5)
fmt.Println(result)   // 25
```

#### Immediately Invoked (IIFE)

```go
numbers := []int{1, 2, 3}

for _, number := range numbers {
    func() {
        fmt.Println(number)
    }()
}
// Output: 1, 2, 3
```

### Recursive Function

Function that calls itself. Needs a **base case** to stop.

```go
func recursiveFactorialFunction(value int) int {
    if value == 1 {
        return value
    }
    return value * recursiveFactorialFunction(value - 1)
}

result := recursiveFactorialFunction(10)
fmt.Println(result, "Example from 10")   // 3628800 Example from 10
```

| Concept | Description |
|---------|-------------|
| **Base case** | `if value == 1` — stops recursion |
| **Recursive case** | `value * recursiveFactorialFunction(value - 1)` — calls itself |

Step by step:

```
factorial(10)
→ 10 * factorial(9)
→ 10 * 9 * factorial(8) ...
→ 10 * 9 * 8 * 7 * 6 * 5 * 4 * 3 * 2 * 1
→ 3628800
```

> **Note:** Go doesn't have tail-call optimization. Deep recursion may cause stack overflow — prefer loops for large iterations.

### Closure

Function that captures variables from its surrounding scope — state persists across calls.

```go
func closureFunction() {
    counter := 0

    closure := func() {
        fmt.Println("Increment")
        counter++
    }

    closure()
    closure()
    closure()

    fmt.Println(counter)
}
```

Output:

```
=== Closure Function ===
Increment
Increment
Increment
3
```

| Feature | Description |
|---------|-------------|
| **Captures variables** | Remembers `counter` from outer scope |
| **State persists** | Counter increments across calls |
| **Shared state** | All closure calls share the same captured variable |
