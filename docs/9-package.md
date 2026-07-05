## Go Package & Import

Custom packages (`helper`, `db`, `blank`) live under `basics/example-package/`.

### Package Import

Import path starts from module root (`basic-module`), not project root.

```go
import(
    "fmt"
    "basic-module/basics/example-package/helper"
    "basic-module/basics/example-package/db"
)
```

### Access Modifier (Exported vs Unexported)

Visibility is determined by the **first letter**:

- **Uppercase** → exported (public) — accessible from other packages
- **Lowercase** → unexported (private) — package-local only

```go
// example-package/helper/helper.go
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
// example-package/db/db.go
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
    _ "basic-module/basics/example-package/blank" // trigger init only
)
```

```go
// example-package/blank/blank.go
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
