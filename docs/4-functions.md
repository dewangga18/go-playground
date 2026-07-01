## Go Functions

A function is a reusable block of code that performs a specific task. Go functions are declared with `func`, can take parameters, return values, and can be passed around as values.

### Basic Function

A function with no parameters and no return value.

```go
func basicFunction() {
    fmt.Println("Hi!")
}
```

```
=== Basic Function ===
Hi!
```

| Component | Description |
|-----------|-------------|
| `func` | Keyword to declare a function |
| `basicFunction` | Function name |
| `()` | Empty parameter list |
| `{}` | Function body |

### Function Parameter

Pass data into a function via parameters. Each parameter has a name and a type.

```go
func functionParameter(firstName string, lastName string) {
    fmt.Println("Hello", firstName, lastName)
}
```

```
=== Function Parameter ===
Hello Aaron Evanjulio
```

> **Note:** If consecutive parameters share the same type, you can write it once: `firstName, lastName string`.

### Return Value

A function can return a value. The return type is specified after the parameter list.

```go
func returnValue(width int, height int) int {
    return width * height
}
```

```
=== Return Value Function ===
Area is 50
```

```go
area := returnValue(10, 5)
fmt.Println("Area is", area)   // 50
```

### Multiple Return Values

Go supports **multiple return values** — a common pattern for returning results alongside errors or metadata.

```go
func returnValue2(width int, height int) (int, string) {
    area := width * height
    if area >= 10 {
        return area, "Large"
    }
    return area, "Small"
}
```

```
=== Return Value Function 2 ===
Area is 300 and size is Large
Area is 300
```

Use `_` (underscore) to ignore a return value you don't need:

```go
area, size := returnValue2(20, 15)
fmt.Println("Area is", area, "and size is", size)   // Area is 300 and size is Large

area2, _ := returnValue2(15, 20)
fmt.Println("Area is", area2)                         // Area is 300
```

> **Note:** The `_` (blank identifier) is a write-only placeholder. Use it whenever a value is returned but not needed — Go will not compile if a declared variable is unused.

### Named Return Values

Return values can be **named** in the function signature. Named return values act as variables declared at the top of the function.

```go
func namedReturnValue(width int, height int) (area int, size string) {
    area = width * height
    if area >= 10 {
        size = "Large"
    }
    return area, size
}
```

```
=== Named Return Value Function ===
Area is 100 and size is Large
```

| Feature | Description |
|---------|-------------|
| `(area int, size string)` | Named returns — declared as local variables |
| `area = width * height` | Assign directly without `:=` |
| `return area, size` | Explicit return (values can be omitted with naked return) |

> **Note:** With named return values, you can use a **naked return** (`return` without values). It returns the current values of the named return variables. Use sparingly — explicit returns are clearer in complex functions.

### Variadic Function

A variadic function accepts a **variable number of arguments**. The `...` prefix before the type indicates variadic.

```go
func variadicFunction(numbers ...int) (total int) {
    for _, number := range numbers {
        total += number
    }

    average := total / len(numbers)
    return average
}
```

```
=== Variadic Function ===
Average is 30
Average is 30
```

Call with multiple arguments:

```go
average := variadicFunction(10, 20, 30, 40, 50)
fmt.Println("Average is", average)   // 30
```

Or pass an existing slice using spread syntax `...`:

```go
sliceWithVariadic := []int{10, 20, 30, 40, 50}
average = variadicFunction(sliceWithVariadic...)
fmt.Println("Average is", average)   // 30
```

| Syntax | Description |
|--------|-------------|
| `numbers ...int` | Accepts zero or more `int` arguments |
| `slice...` | Spreads a slice into individual arguments |

> **Note:** The variadic parameter must be the **last** parameter in the function signature.

### Function As Value

Functions in Go are **first-class citizens** — they can be assigned to variables and passed around like any other value.

```go
func functionAsValue(name string) string {
    return "Good bye, " + name + "!"
}
```

```
=== Function As Value ===
Good bye, Aaron!
```

Assign a function to a variable and call it:

```go
goodBye := functionAsValue
fmt.Println(goodBye("Aaron"))   // Good bye, Aaron!
```

### Function As Parameter

Functions can be passed as arguments to other functions. This enables **callback** patterns.

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
```

```
11 is Odd
12 is Even
```

> **Note:** `strconv.Itoa()` converts an `int` to its string representation. `Itoa` stands for "Integer to ASCII".

For cleaner signatures, define a **type alias** for the function type:

```go
type Filter func(int) string

func functionAsParams2(number int, filter Filter) {
    fmt.Println(filter(number))
}
```

```
13 is Odd
14 is Even
```

| Approach | Code | Description |
|----------|------|-------------|
| Inline | `filter func(int) string` | Function type written directly in parameter |
| Type alias | `type Filter func(int) string` | Named type for reuse |

```go
functionAsParams(11, filterOddNumber)   // 11 is Odd
functionAsParams(12, filterOddNumber)   // 12 is Even
functionAsParams2(13, filterOddNumber)  // 13 is Odd
functionAsParams2(14, filterOddNumber)  // 14 is Even
```

### Anonymous Function

A function without a name. Anonymous functions are useful for short-lived logic, closures, or inline operations.

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

An anonymous function that runs immediately after declaration.

```go
numbers := []int{1, 2, 3}

for _, number := range numbers {
    func() {
        fmt.Println(number)
    }()
}

// Output:
// 1
// 2
// 3
```

| Form | Description |
|------|-------------|
| `func() { ... }` | Assign to variable, call later |
| `func() { ... }()` | Call immediately (IIFE) |

> **Note:** Anonymous functions capture variables from the surrounding scope (closures). In the IIFE example above, `number` is captured from the `for` loop.

### Recursive Function

A function that **calls itself** is called a recursive function. Every recursion needs a **base case** to stop the recursion.

```go
func recursiveFactorialFunction(value int) int {
    if value == 1 {
        return value
    }
    return value * recursiveFactorialFunction(value - 1)
}
```

```
=== Recursive Factorial Function ===
3628800 Example from 10
```

```go
result := recursiveFactorialFunction(10)
fmt.Println(result, "Example from 10")   // 3628800 Example from 10
```

| Concept | Description |
|---------|-------------|
| **Base case** | `if value == 1 { return value }` — stops the recursion |
| **Recursive case** | `value * recursiveFactorialFunction(value - 1)` — calls itself with a reduced value |

How `factorial(10)` evaluates step by step:

```
factorial(10)
→ 10 * factorial(9)
→ 10 * 9 * factorial(8)
→ ...
→ 10 * 9 * 8 * 7 * 6 * 5 * 4 * 3 * 2 * 1  // base case reached
→ 3628800
```

> **Note:** Go does not have tail-call optimization. Deep recursion may cause stack overflow. For large iterations, prefer loops.

### Closure

A **closure** is a function that captures and remembers variables from its surrounding scope, even after that scope has exited.

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

```
=== Closure Function ===
Increment
Increment
Increment
3
```

Key points:
- The anonymous function assigned to `closure` **captures** the `counter` variable from `closureFunction()`
- Each call to `closure()` increments the captured `counter`
- The final value (`3`) proves the closure maintains state across calls

| Feature | Description |
|---------|-------------|
| **Captures variables** | Remembers variables from the outer scope |
| **State persists** | Variables persist between closure calls |
| **Shared state** | Multiple closure calls share the same captured variable |

> **Note:** Closures are commonly used for **stateful functions**, **callbacks**, and **function factories** (functions that return other functions with pre-configured state).
