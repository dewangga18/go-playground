## Go Control Flow

Branching (`if`, `switch`) and looping (`for`) constructs in Go.

### If Expression

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

| Condition | Grade |
|-----------|-------|
| `score >= 80` | A |
| `score >= 70` | B |
| `score >= 60` | C |
| `score >= 50` | D |
| else | E |

> **Note:** First matching condition executes. `else` is optional.

#### Short Statement

Variable declared here is scoped to `if`/`else` block only.

```go
examplePassword := "123456"

if length := len(examplePassword); length >= 8 {
    fmt.Println("Password is valid")
} else {
    fmt.Println("Password is too short")   // output
}
```

### Switch Expression

Cleaner than long `if-else if` chains. **No fall-through by default** — unlike C/Java.

#### Switch with No Expression

Evaluates each `case` as boolean. First match executes.

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

#### Switch with Expression

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

> **Note:** No `break` needed — only matching case runs. Use `fallthrough` explicitly if needed.

#### Short Statement in Switch

```go
examplePassword := "123456"

switch length := len(examplePassword); length >= 8 {
case true:
    fmt.Println("Password is valid")
default:
    fmt.Println("Password is too short")   // output
}
```

### For Loop

Only `for` in Go — no `while` or `do-while`. But `for` covers all use cases.

#### For as While

```go
counter := 1
for counter <= 10 {
    fmt.Println("Iteration:", counter)
    counter++
}
```

#### For with Initialization

Classic C-style.

```go
for counter := 1; counter <= 10; counter++ {
    fmt.Println("Iteration:", counter)
}
```

> **Note:** `counter` is scoped to the `for` block only.

#### For Range

Iterate over collections (slice, array, map, string). Returns (index, value).

```go
colors := []string{"Red", "Yellow", "Green", "Blue"}

for i, color := range colors {
    fmt.Println("Index:", i, "Color:", color)
}
```

Use `_` to ignore index: `for _, color := range colors`.

#### Break

Exit loop immediately.

```go
for _, color := range colors {
    if color == "Green" {
        break
    }
    fmt.Println("Color:", color)
}
// Output: Color: Red, Color: Yellow
```

#### Continue

Skip current iteration.

```go
for _, color := range colors {
    if color == "Green" {
        continue
    }
    fmt.Println("Color:", color)
}
// Output: Color: Red, Color: Yellow, Color: Blue
```

| Keyword | Effect |
|---------|--------|
| `break` | Exits loop entirely |
| `continue` | Skips to next iteration |
