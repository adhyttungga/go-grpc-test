package utils

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userId int64) (string, error) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte("Add private key here"))
	if err != nil {
		log.Printf("error parse the key: %v", err)
		return "", err
	}

	ct := time.Now().UTC()
	claims := make(jwt.MapClaims)
	claims["dat"] = userId
	claims["exp"] = ct.Add(7 * time.Hour).Unix()
	claims["iat"] = ct.Unix()
	claims["nbf"] = ct.Unix()

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)
	if err != nil {
		log.Printf("error generate token: %v", err)
		return "", err
	}

	return token, nil
}
