/*
orders table used by these tests (created automatically by the tests):

mysql
CREATE TABLE IF NOT EXISTS orders (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT NULL
);

psql
CREATE TABLE IF NOT EXISTS orders (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT NULL
);
*/

package database

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

const createOrdersMySQL = `CREATE TABLE IF NOT EXISTS orders (
	id INT AUTO_INCREMENT PRIMARY KEY,
	name VARCHAR(100) NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NULL DEFAULT NULL
)`

const createOrdersPostgres = `CREATE TABLE IF NOT EXISTS orders (
	id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	name VARCHAR(100) NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NULL DEFAULT NULL
)`

// GetConnectionsMySQLParseTime adds parseTime=true to the DSN so MySQL
// returns DATETIME/TIMESTAMP columns as time.Time instead of []byte.
func GetConnectionsMySQLParseTime() *sql.DB {
	db, err := sql.Open("mysql", "go:lang@tcp(localhost:3306)/test?parseTime=true")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)                 // minimum idle connections
	db.SetMaxOpenConns(100)                // maximum open connections
	db.SetConnMaxIdleTime(3 * time.Minute) // maximum idle time
	db.SetConnMaxLifetime(time.Hour)       // maximum connection lifetime

	return db
}

// GetConnectionsPostgresParseTime is a separate pool for this file. lib/pq
// already parses timestamp columns into time.Time, so no extra DSN
// parameter is needed.
func GetConnectionsPostgresParseTime() *sql.DB {
	db, err := sql.Open("postgres", "postgres://go:lang@localhost:5432/test?sslmode=disable")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)                 // minimum idle connections
	db.SetMaxOpenConns(100)                // maximum open connections
	db.SetConnMaxIdleTime(3 * time.Minute) // maximum idle time
	db.SetConnMaxLifetime(time.Hour)       // maximum connection lifetime

	return db
}

func TestAutoParseTimeMySQL(t *testing.T) {
	db := GetConnectionsMySQLParseTime()
	defer db.Close()

	ctx := context.Background()

	if _, err := db.ExecContext(ctx, createOrdersMySQL); err != nil {
		t.Fatal("Error creating MySQL orders table:", err)
	}

	// sql.NullTime{} has Valid=false, so the driver writes NULL
	script := "INSERT INTO orders (name, created_at, updated_at) VALUES (?, ?, ?)"
	if _, err := db.ExecContext(ctx, script, "Keyboard", time.Now(), sql.NullTime{}); err != nil {
		t.Fatal("Error inserting NULL updated_at into MySQL:", err)
	}

	if _, err := db.ExecContext(ctx, script, "Mouse", time.Now(), time.Now()); err != nil {
		t.Fatal("Error inserting into MySQL:", err)
	}

	rows, err := db.QueryContext(ctx, "SELECT id, name, created_at, updated_at FROM orders ORDER BY id")
	if err != nil {
		t.Fatal("Error querying MySQL:", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var createdAt time.Time
		var updatedAt sql.NullTime // NULL becomes Valid=false, not an error
		if err := rows.Scan(&id, &name, &createdAt, &updatedAt); err != nil {
			t.Fatal("Error scanning row:", err)
		}

		updated := "NULL"
		if updatedAt.Valid {
			updated = updatedAt.Time.Format("2006-01-02 15:04:05 MST")
		}
		fmt.Printf("id %d | name %s | created_at %s | updated_at %s\n",
			id, name, createdAt.Format("2006-01-02 15:04:05 MST"), updated)
	}

	if err := rows.Err(); err != nil {
		t.Fatal("Error iterating rows:", err)
	}
}

func TestAutoParseTimePostgres(t *testing.T) {
	db := GetConnectionsPostgresParseTime()
	defer db.Close()

	ctx := context.Background()

	if _, err := db.ExecContext(ctx, createOrdersPostgres); err != nil {
		t.Fatal("Error creating PostgreSQL orders table:", err)
	}

	// sql.NullTime{} has Valid=false, so the driver writes NULL
	script := "INSERT INTO orders (name, created_at, updated_at) VALUES ($1, $2, $3)"
	if _, err := db.ExecContext(ctx, script, "Keyboard", time.Now(), sql.NullTime{}); err != nil {
		t.Fatal("Error inserting NULL updated_at into PostgreSQL:", err)
	}

	if _, err := db.ExecContext(ctx, script, "Mouse", time.Now(), time.Now()); err != nil {
		t.Fatal("Error inserting into PostgreSQL:", err)
	}

	rows, err := db.QueryContext(ctx, "SELECT id, name, created_at, updated_at FROM orders ORDER BY id")
	if err != nil {
		t.Fatal("Error querying PostgreSQL:", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var createdAt time.Time
		var updatedAt sql.NullTime // NULL becomes Valid=false, not an error
		if err := rows.Scan(&id, &name, &createdAt, &updatedAt); err != nil {
			t.Fatal("Error scanning row:", err)
		}

		updated := "NULL"
		if updatedAt.Valid {
			updated = updatedAt.Time.Format("2006-01-02 15:04:05 MST")
		}
		fmt.Printf("id %d | name %s | created_at %s | updated_at %s\n",
			id, name, createdAt.Format("2006-01-02 15:04:05 MST"), updated)
	}

	if err := rows.Err(); err != nil {
		t.Fatal("Error iterating rows:", err)
	}
}
