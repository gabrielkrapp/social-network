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

func TestExtractUsernameFromJWT(t *testing.T) {
	username := "testuser"

	tokenString, err := GenerateJWT(username)
	if err != nil {
		t.Fatalf("Fail to generate JWT: %v", err)
	}

	usernameFromToken, err := ExtractUsernameFromToken(tokenString)
	if err != nil {
		t.Fatalf("Fail to extract user from JWT: %v", err)
	}

	if usernameFromToken != username {
		t.Fatalf("Username extracted from token is incorrect, got %v", usernameFromToken)
	}
}

func TestExtractUsernameFromInvalidJWT(t *testing.T) {
	invalidTokenString := "invalid.token.string"

	_, err := ExtractUsernameFromToken(invalidTokenString)
	if err == nil {
		t.Fatalf("Expected error when extracting from invalid token, but got nil")
	}
}

func TestExtractUsernameFromExpiredJWT(t *testing.T) {
	username := "testuser"

	expirationTime := time.Now().Add(-1 * time.Hour)
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		t.Fatalf("Fail to generate expired JWT: %v", err)
	}

	_, err = ExtractUsernameFromToken(tokenString)
	if err == nil {
		t.Fatalf("Expected error when extracting from expired token, but got nil")
	}
}

func TestExtractUsernameFromJWTWithInvalidSignature(t *testing.T) {
	username := "testuser"

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	invalidKey := []byte("invalid_secret_key")
	tokenString, err := token.SignedString(invalidKey)
	if err != nil {
		t.Fatalf("Fail to generate JWT with invalid signature: %v", err)
	}

	_, err = ExtractUsernameFromToken(tokenString)
	if err == nil {
		t.Fatalf("Expected error when extracting from token with invalid signature, but got nil")
	}
}
