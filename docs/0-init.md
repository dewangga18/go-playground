## Go Modules
Go modules define a project's identity and manage its dependencies.

### Setup

```bash
go mod init <module-name>
```

## Go Build

Compile the Go source code into an executable binary.

> **Note:** Each project can only have one `main` function. If there are more than one, the build will error because the compiler won't know which is the entry point.

```bash
go build ./basics/0-hello-world.go
```

## Go Run

Compile and run Go source code in one step without producing a binary.

```bash
go run ./basics/0-hello-world.go
```
