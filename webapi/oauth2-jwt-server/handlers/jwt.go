package handlers

import (
	"github.com/go-oauth2/oauth2/v4"
	jwtlib "github.com/golang-jwt/jwt/v5"
	"time"
)

var privateKey *rsa.PrivateKey

type JwtAccessTokenGen struct{}

func (j *JwtAccessTokenGen) Token(ctx oauth2.Context, data *oauth2.TokenInfo) (string, error) {
	now := time.Now()

	claims := jwtlib.MapClaims{
		"iss": "your-oauth2-server",
		"sub": data.GetUserID(),
		"aud": data.GetClientID(),
		"exp": now.Add(1 * time.Hour).Unix(),
		"nbf": now.Unix(),
		"iat": now.Unix(),
		"jti": data.GetAccess(),
	}

	token := jwtlib.NewWithClaims(jwtlib.SigningMethodRS256, claims)
	signedToken, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
