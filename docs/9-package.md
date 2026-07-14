## Go Package & Import

Every `.go` file starts with `package <name>` — it's the first line of code. This declares which **group** the file belongs to.

```go
package helper   // this file belongs to the "helper" group
package db       // this file belongs to the "db" group
package main     // this file belongs to the "main" group (executable)
```

**Think of a package like a folder/category for your code:**

| Aspect | Explanation |
|--------|-------------|
| **Same package** = same group | Files in the same folder usually share the same `package` name. They can use each other's functions **without import** — including lowercase (private) ones |
| **Different package** = different group | If you want to use code from another package, you **must import** it first. And you can only access **uppercase (exported)** symbols |
| **Folder ≈ Package** | Convention: the package name matches the folder name. `helper/` folder → `package helper` |

**Example — same package, no import needed:**

```go
// File: myfolder/calc.go
package myfolder

func add(a, b int) int {     // lowercase → private to this package
	return a + b
}

func Add(a, b int) int {     // uppercase → exported
	return a + b
}
```

```go
// File: myfolder/main.go  ← same package!
package myfolder

func Process() {
	result := add(2, 3)  // ✅ works — same package, even though lowercase!
	result2 := Add(2, 3) // ✅ works too
}
```

**Different package — must import:**

```go
// File: cmd/app.go  ← different package!
package main

import "goplayground/myfolder"  // must import to use code from myfolder

func main() {
	// result := myfolder.add(2, 3)  // ❌ compilation error — lowercase = private
	result := myfolder.Add(2, 3)     // ✅ works — uppercase = exported
}
```

> **Key rule:** `add` (lowercase) is visible **only inside** the `myfolder` package. `Add` (uppercase) is visible **to everyone** who imports `myfolder`. This is how Go enforces encapsulation — no `private`/`public` keywords, just case.

---

Custom packages (`helper`, `db`, `blank`) live under `basics/example_package/`.

### Package Import

Import path starts from module root (`goplayground`), not project root.

```go
import(
    "fmt"
    "goplayground/basics/example_package/helper"
    "goplayground/basics/example_package/db"
)
```

### Access Modifier (Exported vs Unexported)

Visibility is determined by the **first letter**:

- **Uppercase** → exported (public) — accessible from other packages
- **Lowercase** → unexported (private) — package-local only

```go
// example_package/helper/helper.go
package helper

var ApplicationName = "Helper" // ✅ exported
var version = "1.1.0"          // ❌ unexported — only visible inside helper

func Square(a int) int {       // ✅ exported
    return a * a
}
```

Usage:

```go
fmt.Println(helper.Square(5))        // 25 ✅ works
fmt.Println(helper.ApplicationName)  // Helper ✅ works
// fmt.Println(helper.version)       // ❌ compilation error — private
```

> **Note:** No `private`/`public` keywords like Java/C#. **Case is the only mechanism.** Capital = exported, lowercase = private.

### Package Init

`init()` runs **automatically** when a package is imported — before `main()`.

```go
// example_package/db/db.go
package db

var connection string

func init() {
    connection = "MySql"
}

func GetDB() string {
    return connection
}
```

Output:

```
=== Package Init Example ===
Database is  MySql
```

Key `init()` behavior:

| Aspect | Behavior |
|--------|----------|
| **When** | Auto-runs on first import |
| **Order** | Before `main()`, in dependency order |
| **Multiple** | >1 `init()` per package allowed (runs in file order) |
| **Return** | Cannot return values |
| **Call** | Cannot be called explicitly — only Go runtime |

### Blank Identifier Import

Sometimes you need a package's `init()` to run without using its exported symbols directly. Use `_` as the package alias.

```go
import(
    _ "goplayground/basics/example_package/blank" // trigger init only
)
```

```go
// example_package/blank/blank.go
package blank

import "fmt"

func init() {
    fmt.Println("Blank package initialized")
}
```

Output (appears first before all other output):

```
Blank package initialized
```

> **Note:** Without `_`, Go refuses to compile — unused imports not allowed. `_` tells the compiler: "I know this isn't used directly, but I need its init side effect."

### Import Patterns

| Pattern | When to Use |
|---------|-------------|
| **Regular** (`"pkg"`) | Need exported functions/variables from that package |
| **Blank** (`_ "pkg"`) | Need `init()` side effects only |
| **Named** (`alias "pkg"`) | Want a shorter alias or avoid name conflicts |
