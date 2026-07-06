## Go Modules

Module management, build, and run commands for a Go project.

- `go mod init <module-name>` — initializes a Go project
- `basic-module` → the module name used in import paths

## Go Build

```bash
go build ./basics/0-hello-world.go
```

- Compiles source into an executable binary
- Only **1 `main` function** allowed per project — error if more exist

## Go Run

```bash
go run ./basics/0-hello-world.go
```

- Compile + run in one step, no binary produced

### With Arguments

```bash
go run ./libraries/3-os.go arg1 arg2
```

- Arguments passed after the file are available via `os.Args`
- `os.Args[0]` = the program path, `os.Args[1:]` = the actual arguments

### With Flags (using `flag` package)

```bash
go run ./libraries/4-flag.go -host=localhost -port=8080 -user=root -password=123456
```

- Alternative to raw `os.Args` — parse named flags with `-key=value` syntax
- Use `flag.String()`, `flag.Int()`, etc. to declare flags, then `flag.Parse()`
- Flags can be passed in **any order** — no need to track positional indexes
- See [`flag` package docs](/docs/0-standard-library.md#flag--command-line-flag-parsing) for details

## External Dependencies

### `go get`

```bash
go get github.com/joho/godotenv
```

- Downloads and adds an external package to `go.mod`

### `go mod tidy`

```bash
go mod tidy
```

- Cleans up `go.mod` and downloads any missing dependencies
- Run this after adding imports from external packages

### `.env` File

```bash
# .env
TEXT=hello_world
```

- Used with `godotenv` to load environment variables from a file
- Place `.env` in the project root
- Call `godotenv.Load(".env")` at startup to populate `os.Getenv()`
