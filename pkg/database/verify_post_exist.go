package database

import (
	"database/sql"
)

func VerifyIfPostExist(db *sql.DB, postID string) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM posts WHERE post_id = $1)"
	err := db.QueryRow(query, postID).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}
