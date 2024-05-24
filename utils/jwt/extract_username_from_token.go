package jwt

import "github.com/dgrijalva/jwt-go"

func ExtractUsernameFromToken(token string) (string, error) {
	claims := &Claims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return "", err
	}

	return claims.Username, nil
}
