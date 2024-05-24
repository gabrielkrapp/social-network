package jwt

import (
	"os"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func init() {
	os.Setenv("JWT_SECRET_KEY", "test_secret_key")
}

func TestGenerateJWT(t *testing.T) {
	username := "testuser"
	tokenString, err := GenerateJWT(username)
	if err != nil {
		t.Fatalf("failed to generate JWT: %v", err)
	}

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		t.Fatalf("failed to parse JWT: %v", err)
	}

	if !token.Valid {
		t.Fatalf("token is not valid")
	}

	if claims.Username != username {
		t.Errorf("expected username %v, got %v", username, claims.Username)
	}

	if !claims.VerifyExpiresAt(time.Now().Add(23*time.Hour).Unix(), true) {
		t.Errorf("token expiry time is incorrect, got %v", claims.ExpiresAt)
	}
}
