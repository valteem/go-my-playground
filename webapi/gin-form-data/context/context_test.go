package context

import (
	"net/http"
	"net/http/httptest"
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

	if !slices.Equal[[]string](getValues, expectValues) {
		t.Errorf("get\n%v\nexpect\n%v", getValues, expectValues)
	}

}
