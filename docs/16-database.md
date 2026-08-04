## Go Database (`database/sql`)

How Go connects to relational databases — documented from the code in the `database/` folder.

---

## How it works

```
App (database/sql code)
    │
    ▼
Database Interface (database/sql)
    │
    ▼
Database Driver (mysql, lib/pq, ...)
    │
    ▼
DBMS (MySQL, PostgreSQL, ...)
```

- Your code only talks to the `database/sql` **interface** — never directly to the DBMS.
- The **driver** translates `database/sql` calls into the DBMS's wire protocol.
- A full list of available drivers: https://go.dev/wiki/SQLDrivers

---

## Drivers used in this project

| DBMS | Driver | Import | Why |
|------|--------|--------|-----|
| MySQL | `github.com/go-sql-driver/mysql` | `_ "github.com/go-sql-driver/mysql"` | Pure Go, de-facto standard MySQL driver |
| PostgreSQL | `github.com/lib/pq` | `_ "github.com/lib/pq"` | Pure Go (no CGo), standard `database/sql` driver |

> Drivers are imported with a **blank identifier** (`_ "..."`) — they register themselves with `database/sql` via `init()`. You never call them directly.

```go
import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    _ "github.com/lib/pq"
)
```

---

## DSN formats

A **DSN** (Data Source Name) is the connection string passed to `sql.Open()`.

| DBMS | DSN format | Example |
|------|------------|---------|
| MySQL | `user:pass@tcp(host:port)/db` | `go:lang@tcp(localhost:3306)/test` |
| PostgreSQL (lib/pq) | `postgres://user:pass@host:port/dbname?sslmode=disable` | `postgres://go:lang@localhost:5432/test?sslmode=disable` |

> ⚠️ The MySQL-style DSN (`user:pass@tcp(host:port)/db`) does **not** work for PostgreSQL — each driver expects its own format.

---

## `sql.Open()` vs `db.Ping()`

```go
db, err := sql.Open("mysql", "go:lang@tcp(localhost:3306)/test")
```

| Call | What it does |
|------|--------------|
| `sql.Open()` | **Only validates the DSN format** — it does **not** connect to the server. Returns `(*sql.DB, error)` |
| `db.Ping()` | **Actually connects** to the DBMS and returns an error if the server is unreachable / credentials are wrong |
| `db.Close()` | Closes the connection pool — always `defer` it |

```go
func TestConnectionMySQL(t *testing.T) {
    db, err := sql.Open("mysql", "go:lang@tcp(localhost:3306)/test")
    if err != nil {
        t.Fatal("Error opening MySQL connection:", err) // DSN invalid
    }
    defer db.Close()

    if err := db.Ping(); err != nil {
        t.Fatal("Error pinging MySQL:", err) // server unreachable / bad creds
    }
    t.Log("MySQL connection successful")
}
```

> A `sql.Open()` that returns `nil` error does **not** mean the DB is reachable — always follow up with `db.Ping()` (or a query).

---

## Connection pool settings

`database/2_polling_test.go` configures the pool the same way for both DBMS — the example below shows MySQL, but `GetConnectionsPostgres()` is identical except for the driver name and DSN.

```go
func GetConnectionsMySQL() *sql.DB {
    db, err := sql.Open("mysql", "go:lang@tcp(localhost:3306)/test")
    if err != nil {
        panic(err)
    }

    db.SetMaxIdleConns(10)                 // max idle connections kept open in the pool
    db.SetMaxOpenConns(100)                // max concurrent open connections
    db.SetConnMaxIdleTime(3 * time.Minute) // how long a connection may sit idle before being closed
    db.SetConnMaxLifetime(time.Hour)       // how long a connection may live before being recycled

    return db
}
```

| Method | Default | What it controls |
|--------|---------|------------------|
| `SetMaxIdleConns(n)` | `2` | Max connections held in the idle pool (0 = no idle conns, `-1` = unlimited) |
| `SetMaxOpenConns(n)` | unlimited | Max connections open at once (0 = unlimited). When full, new requests **wait** |
| `SetConnMaxIdleTime(d)` | unlimited | Idle connections are closed after `d` — frees resources during quiet periods |
| `SetConnMaxLifetime(d)` | unlimited | Connections are recycled after `d` — good for avoiding stale connections / DB restarts |

> **Why pool settings matter:** opening a fresh connection per query is expensive. The pool reuses idle connections (`SetMaxIdleConns`) up to a cap (`SetMaxOpenConns`), and `SetConnMaxLifetime` prevents long-lived connections from going stale (e.g. after the server restarts).

---

## The polling test pattern

`database/2_polling_test.go` opens one pool and fires **100 pings** through it — demonstrating that `database/sql` reuses pooled connections instead of reconnecting each time. There's a MySQL and a Postgres variant; the MySQL one is shown here:

```go
func TestConnectionsMySQL(t *testing.T) {
    db := GetConnectionsMySQL()
    defer db.Close()

    for i := 1; i <= 100; i++ {
        db.Ping()
        fmt.Println("Success ")
    }
}
```

---

## Related docs

- [`docs/0-database-setup.md`](0-database-setup.md) — creating databases & tables from the CLI (MySQL & PostgreSQL), service management
- [`docs/0-init.md`](0-init.md) — `go mod` commands for adding driver dependencies
