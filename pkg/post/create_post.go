package post

import (
	"database/sql"
	"errors"
	"social-network/pkg/database"
)

func CreatePost(db *sql.DB, username, title, description string) error {

	exists, err := database.VerifyIfUserExist(db, username)
	if err != nil {
		return err
	}

	if !exists {
		return errors.New("user does not exist")
	}

	_, err = db.Exec("INSERT INTO posts (username, title, description) VALUES ($1, $2, $3)", username, title, description)
	if err != nil {
		return err
	}

	return err
}
