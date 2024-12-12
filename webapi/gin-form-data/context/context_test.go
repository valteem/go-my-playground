package context

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"slices"
	"strconv"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

const (
	urlKey     = "somekey"
	valueCount = 10
	valueStr   = "value"
	equalSign  = "="
	ampSign    = "&"
)

func initValues() (string, []string) {

	var sbURL, sbSeg, sbSingleValue strings.Builder
	values := make([]string, 0, valueCount)

	sbURL.WriteString("http://localhost/values?")

	for i := 0; i < valueCount; i++ {
		sbSeg.WriteString(urlKey)
		sbSeg.WriteString(equalSign)
		sbSeg.WriteString(valueStr)
		sbSingleValue.WriteString(valueStr)
		sbSeg.WriteString(strconv.Itoa(i))
		sbSingleValue.WriteString(strconv.Itoa(i))
		sbSeg.WriteString(ampSign)
		sbURL.WriteString(sbSeg.String())
		values = append(values, sbSingleValue.String())
		sbSeg.Reset()
		sbSingleValue.Reset()
	}

	url := strings.TrimRight(sbURL.String(), ampSign)

	return url, values

}

func handleValues(key string, values *[]string) gin.HandlerFunc {
	return func(c *gin.Context) {
		*values, _ = c.GetQueryArray(key)
	}
}

func TestKeyArray(t *testing.T) {

	url, expectValues := initValues()
	var getValues []string

	g := gin.Default()
	g.GET("/values", handleValues(urlKey, &getValues))

	req := httptest.NewRequest(http.MethodGet, url, nil)
	resp := httptest.NewRecorder()

	g.ServeHTTP(resp, req)

	if !slices.Equal(getValues, expectValues) {
		t.Errorf("get\n%v\nexpect\n%v", getValues, expectValues)
	}

}

func TestParam(t *testing.T) {

	tests := []struct {
		path     string
		url      string
		varnames []string
		values   []string
	}{
		{
			path:     "/user/:userid/location/:locationid",
			url:      "/user/some_user/location/some_location",
			varnames: []string{"userid", "locationid", "age"},
			values:   []string{"some_user", "some_location", ""},
		},
		{
			path:     "/item/:id/color/:color",
			url:      "/item/some_item/color/green",
			varnames: []string{"id", "color", "quantity"},
			values:   []string{"some_item", "green", ""},
		},
		// adding trailing slash to URL makes handler unreachable
		{
			path:     "/employee/:id",
			url:      "/employee/some_employee/",
			varnames: []string{"employeeid"},
			values:   []string{},
		},
		// mix of path and query variables works fine
		{
			path:     "/book/:bookid",
			url:      "/book/some_book?shelf=1",
			varnames: []string{"bookid"},
			values:   []string{"some_book"},
		},
	}

	for _, tc := range tests {

		var reqPath strings.Builder
		reqPath.WriteString("http://localhost")
		reqPath.WriteString(tc.url)

		req := httptest.NewRequest(http.MethodGet, reqPath.String(), nil)
		resp := httptest.NewRecorder()

		var values []string

		g := gin.Default()
		g.GET(tc.path, func(c *gin.Context) {
			for _, pathVar := range tc.varnames {
				values = append(values, c.Param(pathVar))
			}
		})

		g.ServeHTTP(resp, req)

		if !slices.Equal(values, tc.values) {
			t.Errorf("Param():\nget\n%v\nexpect\n%v", values, tc.values)
		}

	}

}

func TestQuery(t *testing.T) {

	tests := []struct {
		path     string
		url      string
		varnames []string
		values   []string
	}{
		{
			path:     "/query",
			url:      "/query?book=some_book&shelf=top_shelf&room=",
			varnames: []string{"book", "shelf", "room", "building"},
			values:   []string{"some_book", "top_shelf", "", ""},
		},
		{
			path:     "/query/library/:libraryid",
			url:      "/query/library/some_library?book=some_book&shelf=top_shelf&room=",
			varnames: []string{"book", "shelf", "room", "building"},
			values:   []string{"some_book", "top_shelf", "", ""},
		},
	}

	for _, tc := range tests {

		var reqPath strings.Builder
		reqPath.WriteString("http://localhost")
		reqPath.WriteString(tc.url)

		req := httptest.NewRequest(http.MethodGet, reqPath.String(), nil)
		resp := httptest.NewRecorder()

		var values []string

		g := gin.Default()
		g.GET(tc.path, func(c *gin.Context) {
			for _, queryVar := range tc.varnames {
				values = append(values, c.Query(queryVar))
			}
		})

		g.ServeHTTP(resp, req)

		if !slices.Equal(values, tc.values) {
			t.Errorf("Query(): URL %q\nget\n%v\nexpect\n%v", tc.url, values, tc.values)
		}

	}
}

func TestPostFormArray(t *testing.T) {

	tests := []struct {
		formData string
		varnames []string
		output   map[string][]string
	}{
		{
			formData: "var1=value1_1&var1=value1_2&var2=value2&var3=value3_1&var3=value3_2",
			varnames: []string{"var1", "var2", "var3"},
			output:   map[string][]string{"var1": []string{"value1_1", "value1_2"}, "var2": []string{"value2"}, "var3": []string{"value3_1", "value3_2"}},
		},
		{
			formData: "var1=value1&var2=",
			varnames: []string{"var1", "var2"},
			output:   map[string][]string{"var1": []string{"value1"}, "var2": []string{""}},
		},
	}

	for _, tc := range tests {

		req := httptest.NewRequest(http.MethodPost, "http://localhost/form", strings.NewReader(tc.formData))
		req.Header.Set("content-type", "application/x-www-form-urlencoded")
		resp := httptest.NewRecorder()

		output := make(map[string][]string)
		g := gin.Default()
		g.POST("/form", func(c *gin.Context) {
			for _, varname := range tc.varnames {
				output[varname] = c.PostFormArray(varname)
			}
		})

		g.ServeHTTP(resp, req)

		if !reflect.DeepEqual(output, tc.output) {
			t.Errorf("get\n%v\nexpect\n%v", output, tc.output)
		}

	}
}
