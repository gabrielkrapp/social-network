package post

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"social-network/pkg/post"
	"strings"
)

type likePostInfoRequest struct {
	PostId string `json:"postId"`
}

func ToggleLikePostRoute(db *sql.DB, verifyJwt func(token string) (bool, error), extractUsername func(token string) (string, error)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			http.Error(w, "Invalid token format", http.StatusUnauthorized)
			return
		}

		valid, err := verifyJwt(tokenString)
		if err != nil || !valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		username, err := extractUsername(tokenString)
		if err != nil {
			http.Error(w, "Failed to extract username from token", http.StatusUnauthorized)
			return
		}

		var req likePostInfoRequest

		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if err := post.ToggleLike(db, username, req.PostId); err != nil {
			log.Printf("Failed to like a post: %v", err)
			http.Error(w, "Failed to like a post", http.StatusInternalServerError)
			return
		}

		log.Printf("Post liked successfully")

		if _, err := w.Write([]byte("Post Liked successfully")); err != nil {
			log.Printf("Error to send a message: %v", err)
		}
	}
}
