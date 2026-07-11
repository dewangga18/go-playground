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

---

### testify — Assertions & Mocks

```go
import "github.com/stretchr/testify/assert"
import "github.com/stretchr/testify/require"
```

Testify is a popular Go testing toolkit that provides rich assertion functions (`assert` + `require`) and mocking (`testify/mock`).

**Install:**

```bash
go get github.com/stretchr/testify
```

**Two packages — `assert` vs `require`:**

| Package | On failure | Use case |
|---------|------------|----------|
| `assert` | Reports failure, **continues** test (like `t.Error`) | Check multiple assertions in one test |
| `require` | Reports failure, **stops** test immediately (like `t.Fatal`) | Guard conditions — no point continuing |

**Common assertions (`assert` / `require`):**

| Function | Checks | Example |
|----------|--------|---------|
| `Equal(t, want, got)` | Values are equal | `assert.Equal(t, 123, 123)` |
| `NotEqual(t, a, b)` | Values differ | `assert.NotEqual(t, 123, 456)` |
| `Nil(t, val)` | Value is nil | `assert.Nil(t, err)` |
| `NotNil(t, val)` | Value is not nil | `assert.NotNil(t, conn)` |
| `Error(t, err)` | `err` is non-nil | `assert.Error(t, err)` |
| `NoError(t, err)` | `err` is nil | `assert.NoError(t, err)` |
| `True(t, val)` | Value is true | `assert.True(t, ok)` |
| `False(t, val)` | Value is false | `assert.False(t, ok)` |
| `Empty(t, val)` | Length is 0 (or zero value) | `assert.Empty(t, users)` |
| `NotEmpty(t, val)` | Length > 0 | `assert.NotEmpty(t, users)` |
| `Contains(t, str, sub)` | String/slice contains item | `assert.Contains(t, "Aaron", "ar")` |
| `Len(t, obj, n)` | Length equals n | `assert.Len(t, users, 1)` |
| `IsType(t, typ, val)` | Value is expected type | `assert.IsType(t, [1]string{}, users)` |
| `Zero(t, val)` | Value is zero-value for its type | `assert.Zero(t, 0)` |
| `NotZero(t, val)` | Value is not zero-value | `assert.NotZero(t, 1)` |

**Example — `test/sample_testify_test.go`:**

```go
func TestFuncAssertion(t *testing.T) {
	assert.Equal(t, 123, 123)
	assert.NotEqual(t, 123, 456)
	assert.Nil(t, nil)
	assert.NotNil(t, 1)

	err := errors.New("something went wrong")
	assert.Error(t, err)
	assert.NoError(t, nil)

	assert.True(t, true)
	assert.False(t, false)
}

func TestRequireFunction(t *testing.T) {
	require.NoError(t, nil)
	require.Error(t, errors.New("empty"))
	require.Nil(t, nil)
	require.NotNil(t, 1)
	require.Equal(t, true, true)
	require.Len(t, []string{}, 0)
}
```

> **Key insight:** `assert` and `require` share the same function signatures. The only difference is behavior on failure — `assert` continues (like `t.Error`), `require` stops (like `t.Fatal`). Swap between them by changing the import. Use `assert` by default, `require` for setup guards.

**Testify v1.11.1** — installed as a direct dependency.

| Step | Description |
|------|-------------|
| `go get github.com/stretchr/testify` | Add testify to `go.mod` |
| `go mod tidy` | Clean up and download all transitive deps |
| `import "github.com/stretchr/testify/assert"` | Use in test files |

#### Mocking (`testify/mock`)

```go
import "github.com/stretchr/testify/mock"
```

Testify's `mock` package lets you create mock objects for interface-based dependencies.

**Pattern — interface → mock struct → test:**

```go
type UserFetcher interface {
	FetchUser(id int) (string, error)
}

func GetUserGreeting(fetcher UserFetcher, id int) string {
	name, err := fetcher.FetchUser(id)
	if err != nil {
		return "Hello, Guest!"
	}
	return "Hello, " + name + "!"
}

type MockUserFetcher struct {
	mock.Mock
}

func (m *MockUserFetcher) FetchUser(id int) (string, error) {
	args := m.Called(id)
	return args.String(0), args.Error(1)
}
```

**Using the mock in a test:**

```go
func TestGetUserGreeting(t *testing.T) {
	// success case — user found
	mockFetcher := new(MockUserFetcher)
	mockFetcher.On("FetchUser", 1).Return("Aaron", nil)

	result := GetUserGreeting(mockFetcher, 1)
	assert.Equal(t, "Hello, Aaron!", result)
	mockFetcher.AssertExpectations(t)

	// error case — user not found
	mockFetcher2 := new(MockUserFetcher)
	mockFetcher2.On("FetchUser", 999).Return("", errors.New("not found"))

	result2 := GetUserGreeting(mockFetcher2, 999)
	assert.Equal(t, "Hello, Guest!", result2)
	mockFetcher2.AssertExpectations(t)
}
```

**Key mock methods:**

| Method | Description |
|--------|-------------|
| `mockObj.On(method, args...).Return(values...)` | Set up expected call + return values |
| `mockObj.AssertExpectations(t)` | Verify all expected calls were actually made |
| `args.String(n)` | Get nth return value as string |
| `args.Error(n)` | Get nth return value as error |

**Note:** `testify/mock` requires the `github.com/stretchr/objx` package (added automatically as indirect dependency). Run `go mod tidy` after adding mock imports.
