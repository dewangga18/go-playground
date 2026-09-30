## Go Database (`database/sql`)

Connections, pooling, and inserts using the examples in [`database/`](../database/).

## How It Works

```text
Go code -> database/sql -> driver -> database server
                          mysql    MySQL
                          lib/pq   PostgreSQL
```

Your code calls the shared `database/sql` API. The selected driver communicates with the server. The API stays the same across databases, but connection strings and SQL placeholders can differ.

## Drivers and DSNs

Blank imports (`_`) register the drivers without calling them directly:

```go
import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    _ "github.com/lib/pq"
)
```

A **DSN** (Data Source Name) supplies the connection details to `sql.Open()`:

| DBMS | Driver name | DSN used in this project |
|------|-------------|-------------------------|
| MySQL | `mysql` | `go:lang@tcp(localhost:3306)/test` |
| PostgreSQL | `postgres` | `postgres://go:lang@localhost:5432/test?sslmode=disable` |

Both examples use database `test` and account `go` / `lang`. See [Database Setup](0-database-setup.md) for CLI and service commands.

## Opening and Checking Connections

[`1_connection_test.go`](../database/1_connection_test.go) checks each database connection:

```go
db, err := sql.Open("mysql", "go:lang@tcp(localhost:3306)/test")
if err != nil {
    t.Fatal("Error opening MySQL connection:", err)
}
defer db.Close()

if err := db.Ping(); err != nil {
    t.Fatal("Error pinging MySQL:", err)
}
```

| Call | Purpose |
|------|---------|
| `sql.Open()` | Initializes a connection pool; may not connect yet |
| `db.Ping()` | Checks connectivity, establishing a connection if needed |
| `db.Close()` | Closes the pool when the test finishes |

> **Note:** A successful `sql.Open()` alone does not confirm that the server is reachable.

## Connection Pool Settings

`*sql.DB` manages a **pool of connections**, rather than a single connection. Each operation borrows a connection and returns it for reuse, reducing the need to open a new one for every query.

[`2_polling_test.go`](../database/2_polling_test.go) configures both connection helpers with the same limits:

```go
db.SetMaxIdleConns(10)
db.SetMaxOpenConns(100)
db.SetConnMaxIdleTime(3 * time.Minute)
db.SetConnMaxLifetime(time.Hour)
```

| Setting | Effect |
|---------|--------|
| `SetMaxIdleConns(10)` | Keeps up to 10 idle connections for reuse |
| `SetMaxOpenConns(100)` | Allows up to 100 open connections; callers wait when the pool is full |
| `SetConnMaxIdleTime(3 * time.Minute)` | Limits how long connections remain idle |
| `SetConnMaxLifetime(time.Hour)` | Limits how long connections can be reused |

`TestConnectionsMySQL` and `TestConnectionsPostgres` each call `Ping()` 100 times through one pool.

> **Note:** These loop examples ignore `Ping()` errors, so their printed success messages do not prove connectivity. Use the connection tests above to check it.

## Inserting Data with `ExecContext`

[`3_exec_test.go`](../database/3_exec_test.go) inserts into `customer`. Create the table once in `test` using the matching SQL below; the Go file's SQL comment is not executed by the tests.

MySQL:

```sql
CREATE TABLE customer (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);
```

PostgreSQL:

```sql
CREATE TABLE customer (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);
```

Both schemas generate `id` automatically and require a non-null `name`.

```go
db := GetConnectionsMySQL()
defer db.Close()

ctx := context.Background()
_, err := db.ExecContext(ctx, "INSERT INTO customer (name) VALUES (?)", "John Doe")
if err != nil {
    t.Fatal("Error inserting into MySQL:", err)
}
```

`ExecContext` runs statements such as `INSERT`, `UPDATE`, or `DELETE` without returning rows. It returns `sql.Result` and `error`; `_` discards the result here, while `t.Fatal` stops the test if the insert fails.

`context.Background()` has no timeout; see [Context](15-context.md) for cancellation and deadlines.

| DBMS | Placeholder | Value inserted by the test |
|------|-------------|----------------------------|
| MySQL | `?` | `John Doe` |
| PostgreSQL | `$1` | `Jane Doe` |

Placeholders mark where query arguments belong. Their syntax follows the database and driver; `database/sql` does not translate between `?` and `$1`.

MySQL assigns each `?` an argument from left to right. PostgreSQL uses numbered positions: `$1` refers to the first argument, `$2` to the second. Both examples below set customer `1`'s name to `Budi`:

```go
// MySQL
_, err := db.ExecContext(ctx, "UPDATE customer SET name = ? WHERE id = ?", "Budi", 1)

// PostgreSQL (alternative)
_, err := db.ExecContext(ctx, "UPDATE customer SET name = $1 WHERE id = $2", "Budi", 1)
```

> **Note:** In `VALUES (?)` or `VALUES ($1)`, the parentheses belong to SQL's `VALUES` syntax. The placeholder itself is just `?` or `$1`. Pass values as separate arguments rather than concatenating them into SQL.

### Run and Verify

From the project root:

```bash
go test ./database -v -count=1 -run '^TestExecMySQL$'
go test ./database -v -count=1 -run '^TestExecPostgres$'
```

`-run` selects the test; `-count=1` bypasses cached results. In DBeaver or the database CLI, query the matching `test` database:

```sql
SELECT id, name FROM customer ORDER BY id;
```

> **Note:** Each successful run adds one row and leaves it in the database. Duplicate names are allowed. The server must be running, and the Go account needs `INSERT` permission on `customer`.

## References

- [Go `database/sql` API](https://pkg.go.dev/database/sql)
- [Executing SQL statements in Go](https://go.dev/doc/database/change-data)
- [Database setup](0-database-setup.md)
- [Module and dependency commands](0-init.md)
