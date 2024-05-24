package jwt

import (
	"os"
	"testing"
)

func init() {
	os.Setenv("JWT_SECRET_KEY", "test_secret_key")
}

func TestGenerateAndValidateJWT(t *testing.T) {
	username := "testuser"

	tokenString, err := GenerateJWT(username)
	if err != nil {
		t.Fatalf("Fail to generate JWT: %v", err)
	}

	valid, err := VerifyJwt(tokenString)
	if err != nil {
		t.Fatalf("Fail to valid JWT: %v", err)
	}

	if !valid {
		t.Fatalf("token JWT is not valid")
	}
}

func TestValidateInvalidJWT(t *testing.T) {
	invalidTokenString := "invalid.token"

	valid, err := VerifyJwt(invalidTokenString)
	if err == nil {
		t.Fatalf("expected a eror")
	}

	if valid {
		t.Fatalf("expected the token was invalid")
	}
}
