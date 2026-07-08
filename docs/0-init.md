## Go Modules

Module management, build, and run commands for a Go project.

- `go mod init <module-name>` — initializes a Go project
- `basic-module` → the module name used in import paths

### `go mod download`

```bash
go mod download
```

- Downloads all module dependencies into the local cache (`~/go/pkg/mod`)
- Does **not** compile or run — just fetches the source
- Useful in CI/CD pipelines to cache dependencies before building
- ⚠️ **No need to set `$GOPATH` manually** — Go modules work from any directory; cache defaults to `$HOME/go/pkg/mod`

### `go mod verify`

```bash
go mod verify
```

- Verifies the checksums of downloaded modules against `go.sum`
- Detects tampered or corrupted cached dependencies
- Reports `all modules verified` if everything is correct

### `go clean -modcache`

```bash
go clean -modcache
```

- Removes the entire module download cache (`~/go/pkg/mod`)
- Useful when dependencies are corrupted or you want a fresh download
- Run `go mod download` after this to re-fetch dependencies
- ⚠️ **No need to set `$GOPATH` manually** — with Go modules, `$GOPATH` defaults to `$HOME/go` automatically

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

### `go list -m -u all`

```bash
go list -m -u all
```

- Lists all module dependencies and checks for available updates (`-u`)
- For each module, shows `current version` and `latest version` side by side
- Only checks modules that are already in `go.mod`

### `go get -u ./...`

```bash
go get -u ./...
```

- Updates **all** dependencies to their latest available versions
- `./...` means all packages in the current module tree
- Updates `go.mod` and `go.sum` automatically
- ⚠️ **Caution:** May introduce breaking changes — prefer updating specific packages: `go get package@latest`

### `.env` File

```bash
# .env
TEXT=hello_world
```

- Used with `godotenv` to load environment variables from a file
- Place `.env` in the project root
- Call `godotenv.Load(".env")` at startup to populate `os.Getenv()`
