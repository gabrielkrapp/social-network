package config

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestNewHTTPServer(t *testing.T) {
	addr := ":8080"
	server := NewHTTPServer(addr)

	if server.Addr != addr {
		t.Errorf("expected address %v, got %v", addr, server.Addr)
	}

	if server.ReadTimeout != 10*time.Second {
		t.Errorf("expected read timeout %v, got %v", 10*time.Second, server.ReadTimeout)
	}

	if server.WriteTimeout != 10*time.Second {
		t.Errorf("expected write timeout %v, got %v", 10*time.Second, server.WriteTimeout)
	}

	if server.MaxHeaderBytes != 1<<20 {
		t.Errorf("expected max header bytes %v, got %v", 1<<20, server.MaxHeaderBytes)
	}
}

func TestNewHTTPServer_Start(t *testing.T) {
	addr := ":8080"
	server := NewHTTPServer(addr)

	errChan := make(chan error, 1)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errChan <- fmt.Errorf("could not start server: %v", err)
		}
		close(errChan)
	}()

	time.Sleep(100 * time.Millisecond)

	resp, err := http.Get("http://localhost:8080")
	if err != nil {
		t.Fatalf("could not connect to server: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("expected status not found, got %v", resp.StatusCode)
	}

	if err := server.Close(); err != nil {
		t.Fatalf("could not close server: %v", err)
	}
}
