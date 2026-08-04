package database

import (
	"database/sql"
	"fmt"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func GetConnectionsMySQL() *sql.DB {
	db, err := sql.Open("mysql", "go:lang@tcp(localhost:3306)/test")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)                 	// minimum pool open
	db.SetMaxOpenConns(100)                	// maxium pool open
	db.SetConnMaxIdleTime(3 * time.Minute) 	// maximum idle connection
	db.SetConnMaxLifetime(time.Hour)    	// maximum connection lifetime

	return db
}

func GetConnectionsPostgres() *sql.DB {
	db, err := sql.Open("postgres", "postgres://go:lang@localhost:5432/test?sslmode=disable")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)                 	// minimum pool open
	db.SetMaxOpenConns(100)                	// maxium pool open
	db.SetConnMaxIdleTime(3 * time.Minute) 	// maximum idle connection
	db.SetConnMaxLifetime(time.Hour)    	// maximum connection lifetime

	return db
}

func TestConnectionsMySQL(t *testing.T) {
	db := GetConnectionsMySQL()
	defer db.Close()

	for i := 1; i <= 100; i++ {
		db.Ping()
		fmt.Println("Success ")
	}
}

func TestConnectionsPostgres(t *testing.T) {
	db := GetConnectionsPostgres()
	defer db.Close()

	for i := 1; i <= 100; i++ {
		db.Ping()
		fmt.Println("Success ")
	}
}