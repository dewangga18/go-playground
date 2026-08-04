package database

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func TestConnectionMySQL(t *testing.T) {
	// DSN format: user:pass@tcp(host:port)/db
	db, err := sql.Open("mysql", "go:lang@tcp(localhost:3306)/test")
	if err != nil {
		t.Fatal("Error opening MySQL connection:", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		t.Fatal("Error pinging MySQL:", err)
	}
	t.Log("MySQL connection successful")
}

func TestConnectionPostgres(t *testing.T) {
	// lib/pq accepts a URL DSN: postgres://user:pass@host:port/dbname
	db, err := sql.Open("postgres", "postgres://go:lang@localhost:5432/test?sslmode=disable")
	if err != nil {
		t.Fatal("Error opening PostgreSQL connection:", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		t.Fatal("Error pinging PostgreSQL:", err)
	}
	t.Log("PostgreSQL connection successful")
}