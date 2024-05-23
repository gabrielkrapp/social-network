package auth

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"social-network/pkg/auth"
	"social-network/pkg/database"
	"social-network/utils"
)

func Login(db *sql.DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var req userInfosRequest

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		exists, err := database.VerifyIfUserExist(db, req.Username)
		if err != nil {
			http.Error(w, "Error verifying if user exist", http.StatusInternalServerError)
			return
		}

		if !exists {
			http.Error(w, "User don’t exist", http.StatusBadRequest)
			return
		}

		err = auth.VerifyPassword(db, req.Username, req.Password)
		if err != nil {
			http.Error(w, "Invalid password", http.StatusUnauthorized)
			return
		}

		tokenString, err := utils.GenerateJWT(req.Username)
		if err != nil {
			http.Error(w, "Error generating token", http.StatusInternalServerError)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:  "token",
			Value: tokenString,
			Path:  "/",
		})

		log.Printf("User %s logged in successfully", req.Username)
		w.Write([]byte("User logged in successfully"))
	}
}
