## Go Third Party Packages

Third-party packages I've integrated into this project — installation, setup, and usage notes.

---

### godotenv — Load `.env` Files

```go
import "github.com/joho/godotenv"
```

Load environment variables from a `.env` file into `os.Getenv()` at runtime.

**Install:**

```bash
go get github.com/joho/godotenv
```

**Setup — Create `.env` file in project root:**

```bash
# .env
TEXT=hello_world
```

**Usage:**

```go
err := godotenv.Load(".env")
if err != nil {
    fmt.Println("Error loading .env file")
}

e := os.Getenv("TEXT")
fmt.Println("TEXT:", e)   // TEXT: hello_world
```

| Step | Description |
|------|-------------|
| `godotenv.Load(\".env\")` | Reads the `.env` file and sets variables in `os.Getenv()` |
| `os.Getenv(\"TEXT\")` | Access the loaded variable anywhere in the program |

> **Note:** `godotenv.Load()` doesn't overwrite existing environment variables by default. Use `godotenv.Overload()` to force override. The `.env` file should not be committed to version control — add it to `.gitignore`.
