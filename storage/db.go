package storage

import (
	"database/sql"                   // SQL package for database interaction
	"log"                            // Logging errors
	_ "github.com/mattn/go-sqlite3"  // SQLite driver (imported for side effects only)
)

// DB is a global database connection variable
var DB *sql.DB

// InitDB initializes the SQLite database
func InitDB() {
	var err error

	// Open SQLite database connection
	DB, err = sql.Open("sqlite3", "./urls.db")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Create table if it doesn't exist
	createTable := `
	CREATE TABLE IF NOT EXISTS urls (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		short_id TEXT NOT NULL UNIQUE,
		original_url TEXT NOT NULL
	);`

	// Execute the table creation query
	_, err = DB.Exec(createTable)
	if err != nil {
		log.Fatal("Failed to create table:", err)
	}
}
