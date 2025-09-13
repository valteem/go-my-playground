package ensure

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"webapi/jwtmware/config"
)

const (
	authPrefix = "Bearer "
)

var (
	cfg                       *config.Config
	errAuthHeaderNotFound     = errors.New("authorization header not found")
	errInvalidAuthHeader      = errors.New("invalid authorization header")
	errTokenGenerationFailure = errors.New("token generation failure")
	errTokenExpired           = errors.New("token expired")
	errUserNotAuthorized      = errors.New("user not authorized")
)

type userCred struct {
	Username string
	Password string
}

func validateToken(r *http.Request) error {

	authHeader, ok := r.Header["Authorization"]
	if !ok {
		return errAuthHeaderNotFound
	}
	tokenString := strings.TrimPrefix(authHeader[0], authPrefix)
	token, err := jwt.Parse(tokenString,
		func(*jwt.Token) (any, error) {
			return []byte(cfg.Secret), nil
		},
		jwt.WithExpirationRequired(),
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
	)
	if err != nil {
		return err
	}

	claims := token.Claims.(jwt.MapClaims)
	if claims["username"].(string) != cfg.UserName {
		return errUserNotAuthorized
	}
	tokenExpirationTime := claims["exp"].(float64)
	if tokenExpirationTime < float64(time.Now().Unix()) {
		return errTokenExpired
	}

	return nil
}

type EnsureAuth struct {
	handler http.HandlerFunc
}

func NewEnsureAuth(handler http.HandlerFunc) *EnsureAuth {
	return &EnsureAuth{handler: handler}
}

func (er *EnsureAuth) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := validateToken(r); err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	er.handler(w, r)
}

type Server struct {
	http.Server
}

func (s *Server) handlePing(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

type authToken struct {
	AccessToken string `json:"accesstoken"`
}

func (s *Server) handleLogin(w http.ResponseWriter, r *http.Request) {

	var (
		err   error
		token string
		at    authToken
	)

	defer func() {
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}()

	credentialsBase64 := r.Header.Get("Authorization")
	decoded, err := base64.StdEncoding.DecodeString(credentialsBase64)
	if err != nil {
		return
	}
	credentialsStr := strings.Split(string(decoded), ":")
	if len(credentialsStr) != 2 {
		err = errInvalidAuthHeader
		return
	}

	if credentialsStr[0] != cfg.UserName || credentialsStr[1] != cfg.Password {
		err = errUserNotAuthorized
		return
	}

	if token, err = generateJWT(credentialsStr[0]); err != nil {
		err = errTokenGenerationFailure
		return
	}

	at = authToken{AccessToken: token}
	encoded, err := json.Marshal(at)
	if err != nil {
		return
	}

	w.Write(encoded)

}

func loadConfig(ctx context.Context) {
	var err error
	cfg, err = config.Load(ctx)
	if err != nil {
		panic(err)
	}
}

func Start() *Server {

	loadConfig(context.Background())

	mux := http.NewServeMux()

	s := &Server{
		Server: http.Server{
			Handler: mux,
		},
	}

	mux.Handle("/ping", NewEnsureAuth(s.handlePing))
	mux.HandleFunc("/login", s.handleLogin)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.Port))
	if err != nil {
		return nil
	}

	go func() {
		if err := s.Serve(listener); err != nil {
			panic(err)
		}
	}()

	return s

}

func generateJWT(username string) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["autorized"] = true
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(cfg.Expire)).Unix()

	return token.SignedString([]byte(cfg.Secret))
}
