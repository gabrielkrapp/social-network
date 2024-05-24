package jwt

import "github.com/dgrijalva/jwt-go"

func VerifyJwt(token string) (bool, error) {
	claims := &Claims{}
	parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return false, err
		}
		return false, err
	}

	if !parsedToken.Valid {
		return false, nil
	}

	return true, nil
}
