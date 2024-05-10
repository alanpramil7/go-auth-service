package models

import (
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(db *sql.DB, creds Credentials) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", creds.UserName, string(hashedPassword))

	return err
}
