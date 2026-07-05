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
