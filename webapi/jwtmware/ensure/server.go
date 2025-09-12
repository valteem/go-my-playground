// cybertec-postgresql/pgwatch

package ensure

import (
	"context"
	"encoding/json"
	"errors"
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
	cfg                   *config.Config
	errAuthHeaderNotFound = errors.New("authorization header not found")
	errUserNotAuthorized  = errors.New("user not authorized")
	errTokenExpired       = errors.New("token expired")
)

type userCred struct {
	Username string
	Password string
}

func validateToken(r *http.Request) error {

	u := userCred{}

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		return err
	}

	authHeader, ok := r.Header["Authorization"]
	if !ok {
		return errAuthHeaderNotFound
	}
	tokenString := strings.TrimPrefix(authHeader[0], authPrefix)
	token, err := jwt.Parse(tokenString,
		func(*jwt.Token) (any, error) {
			return cfg.Secret, nil
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
	if claims["expire"].(time.Time).Unix() < time.Now().Unix() {
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
	if err := validateToken(r); r != nil {
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

type authTocken struct {
	AccessToken string `json:"accesstoken"`
}

func (s *Server) handleLogin(w http.ResponseWriter, r *http.Request) {

	var (
		err   error
		ut    userCred
		token string
		at    authTocken
	)

	defer func() {
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}()

	if err = json.NewDecoder(r.Body).Decode(&ut); err != nil {
		return
	}

	if ut.Username != cfg.UserName || ut.Password != cfg.Password {
		return
	}

	if token, err = generateJWT(ut.Username); err != nil {
		return
	}

	at = authTocken{AccessToken: token}
	if err = json.NewEncoder(w).Encode(at); err != nil {
		return
	}

}

func Start() *Server {

	var err error
	cfg, err = config.Load(context.Background())
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()

	s := &Server{
		Server: http.Server{
			Handler: mux,
		},
	}

	mux.Handle("/ping", NewEnsureAuth(s.handlePing))
	mux.HandleFunc("/login", s.handleLogin)

	return s

}

func generateJWT(username string) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["autorized"] = true
	claims["username"] = username
	claims["expire"] = time.Now().Add(time.Hour * time.Duration(cfg.Expire)).Unix()

	return token.SignedString(cfg.Secret)
}
