package post

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"social-network/pkg/post"
	"strings"
)

type postInfosRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func CreatePostRoute(db *sql.DB, verifyJwt func(token string) (bool, error), extractUsername func(token string) (string, error)) func(http.ResponseWriter, *http.Request) {
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

		var req postInfosRequest

		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if err := post.CreatePost(db, username, req.Title, req.Description); err != nil {
			log.Printf("Failed to create a post: %v", err)
			http.Error(w, "Failed to create a post", http.StatusInternalServerError)
			return
		}

		log.Printf("Post created successfully")

		if _, err := w.Write([]byte("Post created successfully")); err != nil {
			log.Printf("Error to send a message: %v", err)
		}
	}
}
