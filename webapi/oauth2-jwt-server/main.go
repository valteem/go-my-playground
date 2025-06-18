package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	//	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
	//	jwt "github.com/golang-jwt/jwt/v5"
	"net/http"
	//	"time"

	"webapi/oauth2-jwt-server/config"
	"webapi/oauth2-jwt-server/handlers"
)

var privateKey *rsa.PrivateKey

func init() {
	var err error
	privateKey, err = rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
}

func main() {
	config.ConnectDB()

	manager := manage.NewDefaultManager()
	manager.MapTokenStorage(store.NewMemoryTokenStore())
	manager.MapAccessGenerate(&handlers.JwtAccessTokenGen{})

	clientStore := store.NewClientStore()
	clientStore.Set("clientID", &models.Client{
		ID:     "clientID",
		Secret: "clientSecret",
		Domain: "http://localhost:8080",
	})
	manager.MapClientStorage(clientStore)

	srv := server.NewDefaultServer(manager)
	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)

	http.HandleFunc("/authorize", func(w http.ResponseWriter, r *http.Request) {
		handlers.AuthorizeHandler(srv, w, r)
	})

	http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		srv.HandleTokenRequest(w, r)
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		handlers.LoginHandler(w, r)
	})

	fmt.Println("OAuth2 server running at http://localhost:9090")
	http.ListenAndServe(":9090", nil)
}
