## Go Database (`database/sql`)

Connections, pooling, inserts, and queries using the examples in [`database/`](../database/).

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

## Querying Rows with `QueryContext`

[`4_query_test.go`](../database/4_query_test.go) reads the `customer` table with a `SELECT`:

```go
db := GetConnectionsMySQL()
defer db.Close()

ctx := context.Background()
script := "SELECT id, name FROM customer"

rows, err := db.QueryContext(ctx, script)
if err != nil {
    panic(err)
}

for rows.Next() {
    var id, name string
    if err := rows.Scan(&id, &name); err != nil {
        t.Fatal("Error scanning row:", err)
    }
    fmt.Println("id ", id)
    fmt.Println("name ", name)
}

if err := rows.Err(); err != nil {
    t.Fatal("Error iterating rows:", err)
}

defer rows.Close()
```

`TestQueryPostgres` runs the same code through `GetConnectionsPostgres()`. The query has no arguments, so the identical SQL works for both DBMS and needs no placeholders.

| Call | Purpose |
|------|---------|
| `db.QueryContext(ctx, script)` | Runs a `SELECT` and returns `*sql.Rows` |
| `rows.Next()` | Advances to the next row; returns `false` at the end **or** on an error |
| `rows.Scan(&id, &name)` | Copies the current row's columns into the variables, left to right |
| `rows.Err()` | Reports the error that stopped iteration, if any |
| `rows.Close()` | Releases `rows` and returns its connection to the pool |

Where `ExecContext` runs statements without rows and returns `sql.Result`, `QueryContext` returns `*sql.Rows`, a cursor over the result set. A query error is reported by the `err` from `QueryContext` itself, which the tests pass to `panic`.

### Checking `rows.Err()` After the Loop

`rows.Next()` returns `false` both when every row has been read and when iteration stopped early on an error. Only `rows.Err()` tells the two cases apart:

```go
for rows.Next() {
    // ...
}
if err := rows.Err(); err != nil {
    t.Fatal("Error iterating rows:", err)
}
```

> **Note:** Skipping `rows.Err()` can hide a mid-iteration failure such as a dropped connection; the loop would simply end early, as if the table were shorter.

The `Scan` destinations decide the Go types. Both `id` and `name` scan into `string` here: the driver converts each column's value to fit the destination, so the integer `id` prints as text. A destination that cannot hold a column's value makes `Scan` return an error.

> **Note:** In the tests, `defer rows.Close()` sits after the loop; deferring it immediately after `QueryContext` succeeds is the usual habit so every path releases the connection. `rows` keeps its connection checked out of the pool until `Close` is called.

> **Note:** The tests read whatever `customer` contains at the time. The insert tests above are an easy way to add rows first.

### Run and Verify

From the project root:

```bash
go test ./database -v -count=1 -run '^TestQueryMySQL$'
go test ./database -v -count=1 -run '^TestQueryPostgres$'
```

`-v` shows the `id` and `name` lines printed by the loop. The server must be running, and the Go account needs `SELECT` permission on `customer`.

## Automatic Time Parsing and NULL Columns

[`5_auto_parse_time_test.go`](../database/5_auto_parse_time_test.go) works with a new `orders` table in `test` that adds a nullable `updated_at` column beside a required `created_at`:

| Column | MySQL | PostgreSQL |
|--------|-------|------------|
| `created_at` | `TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP` | same |
| `updated_at` | `TIMESTAMP NULL DEFAULT NULL` | same |

Both tests create the table with `CREATE TABLE IF NOT EXISTS`, so no manual setup is needed.

### Own Connection Pools

The file defines its own helpers, `GetConnectionsMySQLParseTime()` and `GetConnectionsPostgresParseTime()`, separate from the pools in `2_polling_test.go`:

```go
// MySQL: parseTime=true makes the driver return time.Time
db, err := sql.Open("mysql", "go:lang@tcp(localhost:3306)/test?parseTime=true")

// PostgreSQL: lib/pq parses timestamps into time.Time on its own
db, err := sql.Open("postgres", "postgres://go:lang@localhost:5432/test?sslmode=disable")
```

| DBMS | Extra DSN parameter | Effect |
|------|---------------------|--------|
| MySQL | `parseTime=true` | `DATETIME`/`TIMESTAMP` columns scan into `time.Time` instead of `[]byte` |
| PostgreSQL | none | `lib/pq` already decodes `timestamp` columns into `time.Time` |

> **Note:** MySQL converts `time.Time` arguments and results with the DSN's `loc` setting, which defaults to `UTC`. Add `&loc=Local` to keep your machine's wall time. PostgreSQL `timestamp` (without time zone) stores the wall time as written, and `lib/pq` labels the scanned value `UTC`.

### Scanning a Nullable Time Column

A plain `time.Time` destination fails when the value is `NULL`. `sql.NullTime` wraps a `time.Time` with a `Valid` flag instead:

```go
var updatedAt sql.NullTime // NULL becomes Valid=false, not an error
if err := rows.Scan(&id, &name, &createdAt, &updatedAt); err != nil {
    t.Fatal("Error scanning row:", err)
}

updated := "NULL"
if updatedAt.Valid {
    updated = updatedAt.Time.Format("2006-01-02 15:04:05 MST")
}
```

`Valid == false` means the column was `NULL`; reading `.Time` without checking `Valid` gives the zero time.

### All `sql.Null...` Types

Every nullable SQL type has a matching wrapper. Each struct holds the value plus a `Valid bool`:

| Type | Wrapped Go type | Typical SQL columns |
|------|-----------------|---------------------|
| `sql.NullString` | `string` | `VARCHAR`, `CHAR`, `TEXT` |
| `sql.NullBool` | `bool` | `BOOLEAN`, `TINYINT(1)` |
| `sql.NullInt16` | `int16` | `SMALLINT` |
| `sql.NullInt32` | `int32` | `INT`, `integer` |
| `sql.NullInt64` | `int64` | `BIGINT`, `SERIAL`, `BIGSERIAL` |
| `sql.NullByte` | `byte` (`uint8`) | `TINYINT UNSIGNED` |
| `sql.NullFloat64` | `float64` | `DOUBLE`, `REAL`, `DECIMAL` |
| `sql.NullTime` | `time.Time` | `TIMESTAMP`, `DATETIME`, `DATE` |

There is also a generic form that covers the same cases with one name. It stores the value in `V` instead of a type-specific field (`String`, `Int64`, `Time`, ...):

```go
var name sql.Null[string]
if err := rows.Scan(&name); err != nil {
    t.Fatal("Error scanning row:", err)
}
if name.Valid {
    fmt.Println(name.V)
} else {
    fmt.Println("name is NULL")
}
```

`T` should be one of the types accepted by `driver.Value`.

> **Note:** All of these implement `driver.Valuer`, so they can be passed as query arguments — that is what the `sql.NullTime{}` insert above relies on.

`sql.NullTime` also works as a query argument. Its `Value()` method returns `nil` when `Valid` is `false`, so the driver writes `NULL`:

```go
script := "INSERT INTO orders (name, created_at, updated_at) VALUES (?, ?, ?)"
_, err := db.ExecContext(ctx, script, "Keyboard", time.Now(), sql.NullTime{})
```

Each test run inserts two rows — one with a `NULL` `updated_at` and one with a real time — then prints both states.

### Converting to Plain Go Types

Each wrapper documents its own conversion in its doc comment on pkg.go.dev ([`NullString`](https://pkg.go.dev/database/sql#NullString), [`NullTime`](https://pkg.go.dev/database/sql#NullTime), and so on): the pattern is always `if x.Valid { use x.<Field> }`. There is no `ToString()`-style method — conversion means reading the value field yourself:

| From | Value field | Fallback when `NULL` |
|------|-------------|----------------------|
| `sql.NullString` | `.String` | `""` |
| `sql.NullBool` | `.Bool` | `false` |
| `sql.NullInt16` | `.Int16` | `0` |
| `sql.NullInt32` | `.Int32` | `0` |
| `sql.NullInt64` | `.Int64` | `0` |
| `sql.NullByte` | `.Byte` | `0` |
| `sql.NullFloat64` | `.Float64` | `0` |
| `sql.NullTime` | `.Time` | `time.Time{}` (January 1, year 1) |
| `sql.Null[T]` | `.V` | zero value of `T` |

Because every wrapper has the same shape (value field + `Valid`), the generic form converts through one small helper:

```go
func orZero[T any](v sql.Null[T]) T {
    if v.Valid {
        return v.V
    }
    var zero T
    return zero
}

name := orZero(stringName) // sql.Null[string] -> string
```

For a `time.Time` specifically, prefer an explicit fallback over the zero value:

```go
value := time.Unix(0, 0)
if updatedAt.Valid {
    value = updatedAt.Time
}
```

> **Note:** Scanning `NULL` into a plain non-pointer type fails with an error such as `converting NULL to string is unsupported`. A pointer destination is the other NULL-aware option: with `var name *string`, `rows.Scan(&name)` sets `name` to `nil` for `NULL` and to a filled `*string` otherwise. `sql.Null...` stays the idiomatic choice because `Valid` makes the check explicit at the call site.

### Run and Verify

From the project root:

```bash
go test ./database -v -count=1 -run '^TestAutoParseTimeMySQL$'
go test ./database -v -count=1 -run '^TestAutoParseTimePostgres$'
```

Each successful run adds two rows to `orders`. In DBeaver or the database CLI:

```sql
SELECT id, name, created_at, updated_at FROM orders ORDER BY id;
```

> **Note:** The server must be running, and the Go account needs `CREATE`, `INSERT`, and `SELECT` permission on `test`.

## References

- [Go `database/sql` API](https://pkg.go.dev/database/sql)
- [`sql.Null...` types on pkg.go.dev](https://pkg.go.dev/database/sql#NullString)
- [`Scanner` interface](https://pkg.go.dev/database/sql#Scanner) and [`driver.Valuer` interface](https://pkg.go.dev/database/sql/driver#Valuer)
- [Executing SQL statements in Go](https://go.dev/doc/database/change-data)
- [Querying data in Go](https://go.dev/doc/database/querying)
- [Database setup](0-database-setup.md)
- [Module and dependency commands](0-init.md)
