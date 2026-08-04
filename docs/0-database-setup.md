## Database Setup — MySQL & PostgreSQL

Setup guide for creating databases and tables from the CLI. This is a **prerequisite** doc — once your DB and tables exist, see [`docs/16-database.md`](16-database.md) for connecting from Go (`database/sql`).

---

## MySQL

**1. Log in to the MySQL CLI**

```bash
mysql -u root -p
```

(Enter your password when prompted)

**2. Create the database (if it doesn't exist yet)**

```sql
CREATE DATABASE database_name;
```

**3. Switch to that database**

```sql
USE database_name;
```

**4. Create a table**

```sql
CREATE TABLE customers (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

**5. Verify the table was created**

```sql
SHOW TABLES;
DESCRIBE customers;
```

---

## PostgreSQL

**1. Log in to psql**

```bash
psql -U postgres
```

Or connect directly to a specific database:

```bash
psql -U postgres -d database_name
```

**2. Create the database (if it doesn't exist yet)** — run from inside psql:

```sql
CREATE DATABASE customer;
```

**3. Switch to that database**

```
\c customer
```

**4. Create a table**

```sql
CREATE TABLE customers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

**5. Verify the table was created**

```
\dt
```

Inspect the column structure:

```
\d users
```

---

## Key differences to remember

| Operation | MySQL | PostgreSQL |
|-----------|-------|------------|
| Auto-increment | `AUTO_INCREMENT` | `SERIAL` |
| Switch database | `USE db;` | `\c db` |
| List databases | `SHOW DATABASES;` | `\l` |
| List tables | `SHOW TABLES;` | `\dt` |
| Table details | `DESCRIBE users;` | `\d users` |
| Exit | `EXIT;` or `\q` | `\q` |

---

## Managing the services (macOS)

```bash
# MySQL (Oracle installer)
alias mysql-start='sudo launchctl load /Library/LaunchDaemons/com.oracle.oss.mysql.mysqld.plist'
alias mysql-stop='sudo launchctl unload /Library/LaunchDaemons/com.oracle.oss.mysql.mysqld.plist'
alias mysql-status='pgrep -x mysqld >/dev/null && echo "MySQL Active" || echo "MySQL Offline"'

# PostgreSQL (Homebrew)
alias pg-start='brew services start postgresql@16'
alias pg-stop='brew services stop postgresql@16'
alias pg-status='brew services list | grep postgresql'
```

Check which ports are listening:

```bash
lsof -iTCP:3306 -sTCP:LISTEN   # MySQL
lsof -iTCP:5432 -sTCP:LISTEN   # PostgreSQL
```

---

## Connecting from Go

Once the database and tables are ready, connect from Go with `database/sql`:

```go
// MySQL — DSN: user:pass@tcp(host:port)/db
sql.Open("mysql", "go:lang@tcp(localhost:3306)/test")

// PostgreSQL (lib/pq) — URL DSN: postgres://user:pass@host:port/dbname
sql.Open("postgres", "postgres://go:lang@localhost:5432/test?sslmode=disable")
```

> Full details on drivers, connection pooling, and code examples live in [`docs/16-database.md`](16-database.md) and the `database/` folder.
