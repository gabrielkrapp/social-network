package auth

import (
	"database/sql"
	"errors"
	"social-network/pkg/database"
)

func RegisterUser(db *sql.DB, username, hashedPassword string) error {

	exists, err := database.VerifyIfUserExist(db, username)
	if err != nil {
		return err
	}

	if exists {
		return errors.New("user already exists")
	}

	_, err = db.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", username, hashedPassword)
	if err != nil {
		return err
	}

	return err
}
