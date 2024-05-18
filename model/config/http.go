package config

import (
	"fmt"
	"net/http"
	"time"
)

func NewHTTPServer(addr string) *http.Server {

	fmt.Printf("Server is starting on %s\n", addr)

	return &http.Server{
		Addr:           addr,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
