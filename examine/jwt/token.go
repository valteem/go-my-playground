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

func NewTokenWithRoles(roles string, ttl time.Duration) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"eat":   time.Now().Add(ttl).Unix(),
		"iat":   time.Now(),
		"roles": roles,
	})

	tokenString, err := token.SignedString(privateKey)

	return tokenString, err

}

func GetRolesFromToken(tokenString string) (string, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidTokenSigningMethod
		}
		return privateKey, nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", ErrFailedToFetchClaimsFromToken
	}

	return claims["roles"].(string), nil

}
