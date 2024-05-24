package utils

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
		t.Fatalf("falha ao gerar JWT: %v", err)
	}

	valid, err := VerifyJwt(tokenString)
	if err != nil {
		t.Fatalf("falha ao validar JWT: %v", err)
	}

	if !valid {
		t.Fatalf("token JWT não é válido")
	}
}

func TestValidateInvalidJWT(t *testing.T) {
	invalidTokenString := "invalid.token"

	valid, err := VerifyJwt(invalidTokenString)
	if err == nil {
		t.Fatalf("esperava um erro, mas não houve nenhum")
	}

	if valid {
		t.Fatalf("esperava que o token fosse inválido")
	}
}
