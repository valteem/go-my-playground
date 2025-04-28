package services

import (
	"context"
	"errors"
	"fmt"
	"time"
	"webapi/product-catalog/hashing"
	"webapi/product-catalog/model"
	"webapi/product-catalog/repository"
	"webapi/product-catalog/repository/repoerr"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrParseToken = errors.New("cannot parse token")
)

type TokenClaims struct {
	jwt.Claims
	UserId int
}

type UserService struct {
	userRepo       repository.User
	passwordHasher hashing.Hasher
	signKey        string
	tokenTTL       time.Duration
}

func NewUserService(ur repository.User, ph hashing.Hasher, sk string, t time.Duration) *UserService {
	return &UserService{
		userRepo:       ur,
		passwordHasher: ph,
		signKey:        sk,
		tokenTTL:       t,
	}
}

func (us *UserService) CreateUser(ctx context.Context, user model.User) (int, error) {

	id, err := us.userRepo.CreateUser(ctx, user)
	if err != nil {
		if errors.Is(err, repoerr.ErrAlreadyExists) {
			return 0, fmt.Errorf("user name %s already exists", user.Name)
		}
		return 0, fmt.Errorf("cannot create user with name %s: %v", user.Name, err)
	}

	return id, nil

}

func (us *UserService) GenerateToken(ctx context.Context, inputUser model.User) (string, error) {

	user, err := us.userRepo.GetUserByNameAndPassword(ctx, inputUser.Name, us.passwordHasher.Hash(inputUser.Password))
	if err != nil {
		if errors.Is(err, repoerr.ErrNotFound) {
			return "", fmt.Errorf("user %s not found", inputUser.Name)
		}
		return "", fmt.Errorf("failed to fetch data for user %s", inputUser.Name)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(us.tokenTTL)},
			IssuedAt:  &jwt.NumericDate{Time: time.Now()},
		},
		user.Id,
	})

	tokenSigned, err := token.SignedString([]byte(us.signKey))
	if err != nil {
		return "", fmt.Errorf("cannot sign token for user %s", inputUser.Name)
	}

	return tokenSigned, nil

}

func (us *UserService) ParseToken(token string) (int, error) {

	parsedToken, err := jwt.ParseWithClaims(token, TokenClaims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return 0, fmt.Errorf("invalid token signing method: %s", token.Method.Alg())
		}
		return []byte(us.signKey), nil
	})

	if err != nil {
		return 0, ErrParseToken
	}

	claims, ok := parsedToken.Claims.(*TokenClaims)
	if !ok {
		return 0, ErrParseToken
	}

	return claims.UserId, nil

}

func (us *UserService) GetUserById(ctx context.Context, id int) (*model.User, error) {
	user, err := us.userRepo.GetUserById(ctx, id)
	return user, err
}

func (us *UserService) GetUserByName(ctx context.Context, name string) (*model.User, error) {
	user, err := us.userRepo.GetUserByName(ctx, name)
	return user, err

}

func (us *UserService) GetUserByNameAndPassword(ctx context.Context, name, password string) (*model.User, error) {
	user, err := us.userRepo.GetUserByNameAndPassword(ctx, name, password)
	return user, err
}
