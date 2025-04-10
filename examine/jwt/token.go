package main

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func init() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("failed to load env file")
	}

}

var (
	ErrInvalidTokenSigningMethod    = errors.New("invalid token signing method")
	ErrFailedToFetchClaimsFromToken = errors.New("failed to fetch claims from token")
	privateKey                      = []byte(os.Getenv("JWT_PRIVATE_KEY"))
)

func NewTokenWithCustomClaims(customClaims map[string]string, ttl time.Duration) (string, error) {

	claims := jwt.MapClaims{
		"eat": time.Now().Add(ttl).Unix(),
		"iat": time.Now(),
	}

	if len(customClaims) > 0 {
		for k, v := range customClaims {
			claims[k] = v
		}
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(privateKey)

	return tokenString, err

}

func GetCustomClaimsFromToken(tokenString string, customClaimNames []string) (map[string]string, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidTokenSigningMethod
		}
		return privateKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, ErrFailedToFetchClaimsFromToken
	}

	output := map[string]string{}
	for _, claimName := range customClaimNames {
		if v, ok := claims[claimName]; ok {
			output[claimName] = v.(string)
		}
	}

	return output, nil

}
