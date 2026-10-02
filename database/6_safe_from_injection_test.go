/*
users table used by these tests (created automatically by the tests):

mysql
CREATE TABLE IF NOT EXISTS users (
	id INT AUTO_INCREMENT PRIMARY KEY,
	username VARCHAR(50) NOT NULL UNIQUE,
	password VARCHAR(100) NOT NULL
);

psql
CREATE TABLE IF NOT EXISTS users (
	id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	username VARCHAR(50) NOT NULL UNIQUE,
	password VARCHAR(100) NOT NULL
);
*/

package database

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

// Account is one row of the users table. Passwords are plaintext only to keep
// the demo readable — real code stores hashes (bcrypt, argon2, ...).
type Account struct {
	ID       int
	Username string
}

const createUsersMySQL = `CREATE TABLE IF NOT EXISTS users (
	id INT AUTO_INCREMENT PRIMARY KEY,
	username VARCHAR(50) NOT NULL UNIQUE,
	password VARCHAR(100) NOT NULL
)`

const createUsersPostgres = `CREATE TABLE IF NOT EXISTS users (
	id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	username VARCHAR(50) NOT NULL UNIQUE,
	password VARCHAR(100) NOT NULL
)`

// username is UNIQUE, so re-seeding on every run inserts nothing twice.
const seedUsersMySQL = "INSERT IGNORE INTO users (username, password) VALUES ('alice', 'wonderland'), ('bob', 'builder')"
const seedUsersPostgres = "INSERT INTO users (username, password) VALUES ('alice', 'wonderland'), ('bob', 'builder') ON CONFLICT DO NOTHING"

// The SQL text is a constant; credentials only ever travel as arguments.
const loginMySQL = "SELECT id, username FROM users WHERE username = ? AND password = ?"
const loginPostgres = "SELECT id, username FROM users WHERE username = $1 AND password = $2"

// loginBypassPayload is the classic auth-bypass edge case. It closes the
// username quote, adds an always-true condition, and comments out the rest —
// including the password check:
//
//	WHERE username = '' OR 1=1 -- ' AND password = '...'
//
// The trailing space matters: MySQL's "--" comment needs whitespace after it.
const loginBypassPayload = "' OR 1=1 -- "

// searchAccount is the shared lookup logic. A login needs a single row, so
// rows.Next() is called exactly once inside an if — never a for loop.
// Arguments are bound by the driver: their content is data, never SQL.
func searchAccount(ctx context.Context, db *sql.DB, script string, args ...any) (Account, bool, error) {
	rows, err := db.QueryContext(ctx, script, args...)
	if err != nil {
		return Account{}, false, err
	}
	defer rows.Close()

	var account Account
	found := false
	if rows.Next() {
		if err := rows.Scan(&account.ID, &account.Username); err != nil {
			return Account{}, false, err
		}
		found = true
	}
	if err := rows.Err(); err != nil {
		return Account{}, false, err
	}
	return account, found, nil
}

// LoginMySQL authenticates one credential pair with bound parameters. An
// injection payload stays a literal username, so the login is refused.
func LoginMySQL(ctx context.Context, db *sql.DB, username, password string) (Account, bool, error) {
	return searchAccount(ctx, db, loginMySQL, username, password)
}

// LoginPostgres is the same check with PostgreSQL's $1/$2 placeholders.
func LoginPostgres(ctx context.Context, db *sql.DB, username, password string) (Account, bool, error) {
	return searchAccount(ctx, db, loginPostgres, username, password)
}

// LoginUnsafe pastes both credentials into the SQL text with concatenation.
// It exists only so the tests can show the auth bypass — never copy this
// pattern into real code.
func LoginUnsafe(ctx context.Context, db *sql.DB, username, password string) (Account, bool, error) {
	script := "SELECT id, username FROM users WHERE username = '" + username +
		"' AND password = '" + password + "'"
	return searchAccount(ctx, db, script)
}

func TestSafeFromInjectionMySQL(t *testing.T) {
	// connection from 2_polling_test.go
	db := GetConnectionsMySQL()
	defer db.Close()

	ctx := context.Background()

	if _, err := db.ExecContext(ctx, createUsersMySQL); err != nil {
		t.Fatal("Error creating MySQL users table:", err)
	}
	if _, err := db.ExecContext(ctx, seedUsersMySQL); err != nil {
		t.Fatal("Error seeding MySQL users:", err)
	}

	// 1. legitimate login
	_, ok, err := LoginMySQL(ctx, db, "alice", "wonderland")
	if err != nil {
		t.Fatal("Error logging in to MySQL:", err)
	}
	if !ok {
		t.Fatal("Expected alice to log in with the correct password")
	}

	// 2. wrong password is refused
	_, ok, err = LoginMySQL(ctx, db, "alice", "wrong-password")
	if err != nil {
		t.Fatal("Error logging in to MySQL:", err)
	}
	if ok {
		t.Fatal("Wrong password must not log in")
	}

	// 3. bypass payload through the placeholders stays a literal username
	_, ok, err = LoginMySQL(ctx, db, loginBypassPayload, "any-password")
	if err != nil {
		t.Fatal("Error running parameterized login on MySQL:", err)
	}
	if ok {
		t.Fatal("Parameterized login must not be bypassed by the payload")
	}

	// 4. the same payload concatenated into the SQL text comments out the
	//    password check and logs in (read-only demo on purpose)
	leaked, ok, err := LoginUnsafe(ctx, db, loginBypassPayload, "any-password")
	if err != nil {
		t.Fatal("Error running unsafe login on MySQL:", err)
	}
	if !ok {
		t.Fatal("Expected the concatenated login to be bypassed")
	}

	fmt.Printf("payload %q -> parameterized: blocked | concatenated: logged in as %q\n",
		loginBypassPayload, leaked.Username)
}

func TestSafeFromInjectionPostgres(t *testing.T) {
	// connection from 2_polling_test.go
	db := GetConnectionsPostgres()
	defer db.Close()

	ctx := context.Background()

	if _, err := db.ExecContext(ctx, createUsersPostgres); err != nil {
		t.Fatal("Error creating PostgreSQL users table:", err)
	}
	if _, err := db.ExecContext(ctx, seedUsersPostgres); err != nil {
		t.Fatal("Error seeding PostgreSQL users:", err)
	}

	// 1. legitimate login
	_, ok, err := LoginPostgres(ctx, db, "alice", "wonderland")
	if err != nil {
		t.Fatal("Error logging in to PostgreSQL:", err)
	}
	if !ok {
		t.Fatal("Expected alice to log in with the correct password")
	}

	// 2. wrong password is refused
	_, ok, err = LoginPostgres(ctx, db, "alice", "wrong-password")
	if err != nil {
		t.Fatal("Error logging in to PostgreSQL:", err)
	}
	if ok {
		t.Fatal("Wrong password must not log in")
	}

	// 3. bypass payload through the placeholders stays a literal username
	_, ok, err = LoginPostgres(ctx, db, loginBypassPayload, "any-password")
	if err != nil {
		t.Fatal("Error running parameterized login on PostgreSQL:", err)
	}
	if ok {
		t.Fatal("Parameterized login must not be bypassed by the payload")
	}

	// 4. the same payload concatenated into the SQL text comments out the
	//    password check and logs in (read-only demo on purpose)
	leaked, ok, err := LoginUnsafe(ctx, db, loginBypassPayload, "any-password")
	if err != nil {
		t.Fatal("Error running unsafe login on PostgreSQL:", err)
	}
	if !ok {
		t.Fatal("Expected the concatenated login to be bypassed")
	}

	fmt.Printf("payload %q -> parameterized: blocked | concatenated: logged in as %q\n",
		loginBypassPayload, leaked.Username)
}
