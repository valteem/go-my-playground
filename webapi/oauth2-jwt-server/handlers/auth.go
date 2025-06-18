package handlers

import (
	"fmt"
	"net/http"
	//	"time"

	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
)

func AuthorizeHandler(srv *server.Server, w http.ResponseWriter, r *http.Request) {
	user := r.FormValue("user")
	if user == "" {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	ctx := r.Context()
	newCtx := oauth2.WithUser(ctx, &models.BasicUser{
		ID: user,
	})
	r = r.WithContext(newCtx)

	srv.AuthorizeTokenAccess(server.GenerateAccessToKen)(w, r)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")

		if username == "test" && password == "test" {
			http.Redirect(w, r, fmt.Sprintf("/authorize?user=%s", username), http.StatusFound)
			return
		}
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	w.Write([]byte(`
        <form method="POST">
            <input name="username" placeholder="username"/>
            <input name="password" placeholder="password" type="password"/>
            <button type="submit">Login</button>
        </form>
    `))
}
