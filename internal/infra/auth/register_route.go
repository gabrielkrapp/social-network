package auth

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"social-network/pkg/auth"

	"golang.org/x/crypto/bcrypt"
)

type userInfosRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Register(db *sql.DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var req userInfosRequest

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Error while hashing the password", http.StatusInternalServerError)
			return
		}

		if err := auth.RegisterUser(db, req.Username, string(hashedPassword)); err != nil {
			log.Printf("Failed to save user: %v", err)
			http.Error(w, "Failed to save the user", http.StatusInternalServerError)
			return
		}

		log.Printf("User %s registered successfully", req.Username)

		if _, err := w.Write([]byte("User registered successfully")); err != nil {
			log.Printf("Error to send a message: %v", err)
		}
	}
}
