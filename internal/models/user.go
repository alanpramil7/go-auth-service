package models

import (
	"database/sql"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// Function to create user in datbase
func RegisterUser(db *sql.DB, creds Credentials) error {
	// Check if user exists
	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE username=$1)", creds.UserName).Scan(&exists)
	if err != nil {
		return err
	}

	if exists {
		return errors.New("User already exist")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", creds.UserName, string(hashedPassword))

	return err
}

// Function to login User
func AutheneticateUser(db *sql.DB, creds Credentials)(string, error){
  var storedPassword string

  err := db.QueryRow("SELECT password FROM users WHERE username = $1", creds.UserName).Scan(&storedPassword)
  if err != nil {
    if err == sql.ErrNoRows{
      return "", errors.New("No user found")
    }
    return "", err
  }

  err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(creds.Password))
  if err != nil {
    return "", errors.New("Invalid Password")
  }

  token, err := GenerateToken(creds.UserName)
  if err != nil {
    return "", err
  }

  return token, nil
}

