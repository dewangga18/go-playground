/*
mysql
CREATE TABLE customer (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

psql
CREATE TABLE customer (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);
*/

package database

import (
	"context"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func TestExecMySQL(t *testing.T) {
    // connection from 2_polling_test.go
    db := GetConnectionsMySQL()
	defer db.Close()

    ctx := context.Background()
    script := "INSERT INTO customer (name) VALUES (?)"

    _, err := db.ExecContext(ctx, script, "John Doe")
    if err != nil {
        t.Fatal("Error inserting into MySQL:", err)
    }

    fmt.Println("Inserted into MySQL successfully")
}

func TestExecPostgres(t *testing.T) {
    // connection from 2_polling_test.go
    db := GetConnectionsPostgres()
    defer db.Close()

    ctx := context.Background()
    script := "INSERT INTO customer (name) VALUES ($1)"
    
    _, err := db.ExecContext(ctx, script, "Jane Doe")
    if err != nil {
        t.Fatal("Error inserting into PostgreSQL:", err)
    }

    fmt.Println("Inserted into PostgreSQL successfully")
}