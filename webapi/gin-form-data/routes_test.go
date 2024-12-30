package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestRoutesPerson(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/person?name=somename&age=33", nil)
	res := httptest.NewRecorder()

	g := gin.Default()
	g.GET("/person", GetPerson)

	g.ServeHTTP(res, req)

	JSONactual := res.Body.String()
	JSONexpected := `{"age":33,"name":"somename"}`
	if JSONactual != JSONexpected {
		t.Errorf("JSON output: get %s, expect %s", JSONactual, JSONexpected)
	}

}

type someNestedStruct struct {
	NestedStr []string `json:"nestedstr"`
	NestedInt []int    `json:"nestedint"`
}

type someOuterStruct struct {
	Nested   []someNestedStruct `json:"nested"`
	OuterStr []string           `json:"outerstr"`
	OuterInt []int              `json:"outerint"`
}

func BenchmarkBindJSON(b *testing.B) {

	// https://stackoverflow.com/a/44339430
	// Find and Replace in Selection only
	input := `{"nested":
{
"nestedstr":
["nested0","nested1","nested2","nested3","nested4","nested5","nested6","nested7"],
"nestedint:"
[0,1,2,3,4,5,6,7]
},
"outerstr:"
["outer0","outer1","outer2","outer3","outer4","outer5","outer6","outer7"],
"outerint":
[0,1,2,3,4,5,6,7]
}`

	req, _ := http.NewRequest(http.MethodPost, "http://localhost:3001", strings.NewReader(input))
	c := gin.Context{}
	c.Request = req

	outer := someOuterStruct{}

	b.Run("gin", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			c.ShouldBindBodyWithJSON(outer)
		}
	})

	b.Run("std", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			json.NewDecoder(req.Body).Decode(&outer)
		}
	})

}

type readHeadersStr struct {
	apples []string `header:"Apples"`
}

type readHeadersInt struct {
	onions []int `header:"Onions"`
}

func BenchmarkBindHeadersStr(b *testing.B) {

	req, _ := http.NewRequest(http.MethodPost, "http://localhost:3001", nil)
	req.Header["Apples"] = []string{"apple0", "apple1", "apple2", "apple3", "apple4", "apple5", "apple6", "apple7"}
	c := gin.Context{}
	c.Request = req

	target := readHeadersStr{}

	b.Run("gin", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			c.ShouldBindHeader(target)
		}
	})

	b.Run("std", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			target.apples = req.Header["Apples"]
		}
	})
}

func BenchmarkBindHeadersInt(b *testing.B) {

	req, _ := http.NewRequest(http.MethodPost, "http://localhost:3001", nil)
	req.Header["Onions"] = []string{"0", "1", "2", "3", "4", "5", "6", "7"}
	c := gin.Context{}
	c.Request = req

	target := readHeadersInt{onions: make([]int, 8)} // pre-allocated slice

	b.Run("gin", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			c.ShouldBindHeader(target)
		}
	})

	b.Run("std", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for idx := 0; idx < len(req.Header["Onions"]); idx++ {
				target.onions[idx], _ = strconv.Atoi(req.Header["Onions"][idx])
			}
		}
	})
}
