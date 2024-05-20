package main

import (
	"log"
	"social-network/config"
	"social-network/config/db"
)

func main() {

	server := config.NewHTTPServer(":8080")

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	database, err := db.DatabaseConnect()

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	defer database.Close()

	router := Router(database)
	server.Handler = router

}
