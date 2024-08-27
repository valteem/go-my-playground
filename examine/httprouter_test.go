package examine

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func HandleOutputParam(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// for _, v := range p {
	// 	w.Write([]byte(v.Key + ":" + v.Value + ";"))
	// }
	j, _ := json.Marshal(p)
	w.Write([]byte(j))
}

func TestParamsBasic(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, "http://a.b:80/param1/v1/param2/v2", nil)
	resp := httptest.NewRecorder()
	//	writer := &mockResponseWriter{}

	h := httprouter.New()
	h.GET("/param1/:param1_value/param2/:param2_value", HandleOutputParam)

	h.ServeHTTP(resp, req)

	r := resp.Result()
	if r.StatusCode != http.StatusOK {
		t.Errorf("expect %d status code, get %d", http.StatusOK, r.StatusCode)
	}

	buf, err := io.ReadAll(r.Body)
	if err != nil {
		t.Errorf("error reading response body: %v", err)
	}
	var outputParams httprouter.Params
	json.Unmarshal(buf, &outputParams)
	expectedParams := httprouter.Params{{Key: "param1_value", Value: "v1"}, {Key: "param2_value", Value: "v2"}}
	if !reflect.DeepEqual(outputParams, expectedParams) {
		t.Errorf("unexpected URL parsing result:\nget\n%vexpect\n%v", outputParams, expectedParams)
	}

}
