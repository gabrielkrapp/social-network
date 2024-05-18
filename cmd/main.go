package main

import (
	"log"
	"social-network/config"
)

func main() {

	server := config.NewHTTPServer(":8080")

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
