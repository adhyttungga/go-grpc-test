package utils

import (
	"fmt"
	"log"

	"github.com/golang-jwt/jwt/v5"
)

func ValidateToken(token string, userId *string) bool {
	key, err := jwt.ParseRSAPublicKeyFromPEM([]byte("add public key here"))
	if err != nil {
		log.Printf("error parse the key: %v", err)
		return false
	}

	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected method: %s", t.Header["alg"])
		}

		return key, nil
	})
	if err != nil {
		log.Printf("error parse the token: %v", err)
		return false
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		log.Printf("error validate token %v", err)
		return false
	}

	id, _ := claims["dat"].(string)
	userId = &id
	return true
}
