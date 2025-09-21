package utils

import (
	"fmt"
	"log"
	"time"

	"github.com/adhyttungga/go-grpc-test/pkg/config"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userId int64) (string, error) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(config.Config.PrivateKey))
	if err != nil {
		log.Printf("error parse the key: %v", err)
		return "", err
	}

	ct := time.Now().UTC()
	claims := make(jwt.MapClaims)
	claims["dat"] = fmt.Sprintf("%d", userId)
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
