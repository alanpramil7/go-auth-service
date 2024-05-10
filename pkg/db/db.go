package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func Initialize(connectionString string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("Cannot Connect to db: %v", err)
	}

	if err = createTable(db); err != nil {
		return nil, err
	}

	fmt.Println("Database connected sucessfully")
	return db, nil
}

func createTable(db *sql.DB) error {
	query := `
          CREATE TABLE IF NOT EXISTS users (
              id SERIAL PRIMARY KEY,
              username VARCHAR(255) UNIQUE NOT NULL,
              password VARCHAR(255) NOT NULL
          );
      `

	_, err := db.Exec(query)
	return err
}
