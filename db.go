package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// InitDB initializes the SQLite database and creates the users table if it doesn't exist.
// this will be used in main.go
func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./tasks.db")
	if err != nil {
		log.Fatal(err)
	}

	// the above only creates a database, but we will also need to create the table
	// SQL statement to create the users table
	createTable := `CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		description TEXT,
		complete BOOLEAN NOT NULL
		);`

	// Execute the SQL statement above
	if _, err = DB.Exec(createTable); err != nil {
		log.Fatal(err)
	}
}
