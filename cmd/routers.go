package main

import (
	"database/sql"
	"social-network/internal/infra"

	"github.com/gorilla/mux"
)

func Router(db *sql.DB) *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/register", infra.Register(db)).Methods("POST")

	return router
}
