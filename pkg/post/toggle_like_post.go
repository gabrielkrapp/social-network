package post

import (
	"database/sql"
	"errors"
	"social-network/pkg/database"
)

func ToggleLike(db *sql.DB, username string, postID string) error {
	userExists, err := database.VerifyIfUserExist(db, username)
	if err != nil {
		return err
	}

	if !userExists {
		return errors.New("user does not exist")
	}

	postExists, err := database.VerifyIfPostExist(db, postID)
	if err != nil {
		return err
	}

	if !postExists {
		return errors.New("post does not exist")
	}

	likeExists, err := database.VerifyIfLikeExist(db, username, postID)
	if err != nil {
		return err
	}

	if likeExists {
		_, err = db.Exec("DELETE FROM likes WHERE username = $1 AND post_id = $2", username, postID)
		if err != nil {
			return err
		}
	} else {
		_, err = db.Exec("INSERT INTO likes (username, post_id) VALUES ($1, $2)", username, postID)
		if err != nil {
			return err
		}
	}

	return nil
}
