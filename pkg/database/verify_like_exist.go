package database

import "database/sql"

func VerifyIfLikeExist(db *sql.DB, username string, postID int) (bool, error) {
	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM likes WHERE username = $1 AND post_id = $2)", username, postID).Scan(&exists)
	return exists, err
}
