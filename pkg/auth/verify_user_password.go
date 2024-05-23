package auth

import (
	"database/sql"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func VerifyPassword(db *sql.DB, username, password string) error {
	var hashedPassword string
	err := db.QueryRow("SELECT password FROM users WHERE username = $1", username).Scan(&hashedPassword)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	if err != nil {
		return errors.New("invalid password")
	}

	return nil
}
