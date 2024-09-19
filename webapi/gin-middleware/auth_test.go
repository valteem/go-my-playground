package main

import (
	"encoding/base64"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func authorizationHeader(user, password string) string {
	base := user + ":" + password
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(base))
}

// Securing website with a few admin user records
// https://github.com/gin-gonic/gin/issues/2226
// https://github.com/gin-gonic/gin/pull/2609
func TestBasicAuth(t *testing.T) {

	accounts := gin.Accounts{"some_user": "some_user_password", "another_user": "another_user_password"}
	testUserName := "some_user"

	mux := gin.New()
	mux.Use(gin.BasicAuth(accounts))
	mux.GET("/login", func(c *gin.Context) {
		c.String(http.StatusOK, c.MustGet(gin.AuthUserKey).(string))
	})

	req := httptest.NewRequest(http.MethodGet, "/login", nil)
	req.Header.Add("Authorization", authorizationHeader(testUserName, accounts[testUserName]))
	resp := httptest.NewRecorder()

	mux.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("return code: get %d, expect %d", resp.Code, http.StatusOK)
	}

	if respUserName := resp.Body.String(); respUserName != testUserName {
		t.Errorf("username: get %s, expect %s", respUserName, testUserName)
	}

}
