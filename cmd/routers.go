package main

import (
	"database/sql"
	"social-network/internal/infra/auth"
	"social-network/internal/infra/post"

	"github.com/gorilla/mux"
)

func Router(db *sql.DB) *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/register", auth.Register(db)).Methods("POST")
	router.HandleFunc("/login", auth.Login(db)).Methods("POST")
	router.HandleFunc("/posts", post.CreatePostRoute(db)).Methods("POST")

	return router
}
