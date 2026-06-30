## Go Control Flow

Control flow statements determine the order in which code is executed. Go provides standard branching constructs like `if` and `switch`, with some unique features like short statements.

### If Expression

The `if` statement evaluates a boolean condition. Parentheses around the condition are optional, but braces `{}` are required.

#### Basic If

```go
score := 60
grade := ""

if score >= 80 {
    grade = "A"
} else if score >= 70 {
    grade = "B"
} else if score >= 60 {
    grade = "C"
} else if score >= 50 {
    grade = "D"
} else {
    grade = "E"
}

fmt.Println("Your score:", score)   // 60
fmt.Println("Your grade is", grade) // C
```

| Condition | grade |
|-----------|-------|
| `score >= 80` | A |
| `score >= 70` | B |
| `score >= 60` | C |
| `score >= 50` | D |
| else | E |

> **Note:** The first matching condition is executed. The rest are skipped. An `else` clause is optional and runs if none of the conditions match.

#### Short Statement

Go allows a short statement to be executed before the condition. Variables declared here are scoped to the `if` block.

```go
examplePassword := "123456"

if length := len(examplePassword); length >= 8 {
    fmt.Println("Password is valid")
} else {
    fmt.Println("Password is too short")   // output
}
```

| Component | Description |
|-----------|-------------|
| `length := len(examplePassword)` | Short statement — runs before condition |
| `length >= 8` | Boolean condition |
| `length` variable | Scoped to `if`/`else` block only |

> **Note:** The `length` variable from the short statement is only accessible inside the `if` and `else` blocks. It does not leak to the outer scope.

### Switch Expression

The `switch` statement is a cleaner way to write long `if-else if` chains. Go's switch is unique — it **does not fall through** by default.

#### Switch with No Expression

When no expression is provided, `switch` evaluates each `case` as a boolean condition. The first matching case executes.

```go
score := 90
grade := ""

switch {
case score >= 90:
    grade = "A"
case score >= 80:
    grade = "B"
case score >= 70:
    grade = "C"
case score >= 60:
    grade = "D"
case score >= 50:
    grade = "D"
default:
    grade = "E"
}

fmt.Println("Your score:", score)   // 90
fmt.Println("Your grade is", grade) // A
```

| Score Range | Grade |
|-------------|-------|
| `>= 90` | A |
| `>= 80` | B |
| `>= 70` | C |
| `>= 60` | D |
| `>= 50` | D |
| else | E |

#### Switch with Expression

Switch can also evaluate a specific value and match against cases.

```go
switch grade {
case "A":
    fmt.Println("Excellent!")   // output
case "B":
    fmt.Println("Good!")
case "C":
    fmt.Println("Average!")
case "D":
    fmt.Println("Poor!")
default:
    fmt.Println("Fail!")
}
```

> **Note:** Unlike many languages, Go's `switch` does **not** need `break` statements. Only the matching case runs — no fall-through. Use `fallthrough` explicitly if you want it.

#### Short Statement in Switch

Like `if`, `switch` also supports a short statement before the condition.

```go
examplePassword := "123456"

switch length := len(examplePassword); length >= 8 {
case true:
    fmt.Println("Password is valid")
default:
    fmt.Println("Password is too short")   // output
}
```

| Component | Description |
|-----------|-------------|
| `length := len(examplePassword)` | Short statement — runs before switch |
| `length >= 8` | Condition to evaluate |
| `case true / default` | Matching cases |

### For Loop

Go only has `for` as its looping construct — no `while` or `do-while`. But `for` can be written in several flexible forms to cover all use cases.

#### For as While

Go doesn't have a separate `while` keyword. Use `for` with just a condition.

```go
counter := 1

for counter <= 10 {
    fmt.Println("Iteration:", counter)
    counter++
}

// Output: Iteration: 1 through 10
```

| Form | Equivalent To |
|------|---------------|
| `for condition {}` | `while (condition)` in other languages |

#### For with Initialization

The classic C-style `for` loop with init, condition, and post statement.

```go
for counter := 1; counter <= 10; counter++ {
    fmt.Println("Iteration:", counter)
}

// Output: Iteration: 1 through 10
```

| Part | Description |
|------|-------------|
| `counter := 1` | Initialization — runs once before loop |
| `counter <= 10` | Condition — checked before each iteration |
| `counter++` | Post statement — runs after each iteration |

> **Note:** The variable declared in initialization (`counter`) is scoped to the `for` block only.

#### For Range

Iterate over elements of a collection (slice, array, map, string). Returns **index** and **value** for each iteration.

```go
colors := []string{"Red", "Yellow", "Green", "Blue"}

for i, color := range colors {
    fmt.Println("Index:", i, "Color:", color)
}

// Output:
// Index: 0 Color: Red
// Index: 1 Color: Yellow
// Index: 2 Color: Green
// Index: 3 Color: Blue
```

| Component | Description |
|-----------|-------------|
| `i` | Index (position in collection) |
| `color` | Value at that position |
| `range colors` | Iterates over each element |

> **Note:** Use `_` (underscore) to ignore the index if you only need the value: `for _, color := range colors`.

#### Break

Exit the loop immediately, skipping remaining iterations.

```go
for _, color := range colors {
    if color == "Green" {
        break
    }
    fmt.Println("Color:", color)
}

// Output:
// Color: Red
// Color: Yellow
```

`break` stops the loop when `"Green"` is encountered — `"Green"` and `"Blue"` are not printed.

#### Continue

Skip the current iteration and move to the next one.

```go
for _, color := range colors {
    if color == "Green" {
        continue
    }
    fmt.Println("Color:", color)
}

// Output:
// Color: Red
// Color: Yellow
// Color: Blue
```

`continue` skips the rest of the loop body for `"Green"`, but continues with the next element (`"Blue"`).

| Keyword | Effect |
|---------|--------|
| `break` | Exits the loop entirely |
| `continue` | Skips to the next iteration |

### If vs Switch vs For

| Aspect | If | Switch | For |
|--------|-----|--------|-----|
| **Purpose** | Branch on condition | Branch on multiple values | Repeat code |
| **Syntax** | `if condition {}` | `switch { case ...: }` | `for init; cond; post {}` |
| **Short statement** | ✅ Supported | ✅ Supported | ✅ In init |
| **Range** | ❌ | ❌ | ✅ `for range` |
| **Break/Continue** | ❌ | ✅ `break` (optional) | ✅ Both |
