package database

import (
	"context"
	"fmt"
	"testing"
)

func TestQueryMySQL(t *testing.T) {
    // connection from 2_polling_test.go
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
}

func TestQueryPostgres(t *testing.T) {    
	// connection from 2_polling_test.go
    db := GetConnectionsPostgres()
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
}