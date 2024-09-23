package reuse_test

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"
)

var (
	testFormField1Value, testFormField2Value string
)

func TestFormData(t *testing.T) {

	mux := http.NewServeMux()
	mux.Handle("/form", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		testFormField1Value = r.FormValue("field1")
		testFormField2Value = r.FormValue("field2")
	}))

	formData := url.Values{
		"field1": {"value1"},
		"field2": {"value2"},
	}

	formDataEncoded := formData.Encode()
	reader := strings.NewReader(formDataEncoded)
	req := httptest.NewRequest(http.MethodPost, "/form", reader)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp := httptest.NewRecorder()

	mux.ServeHTTP(resp, req)

	if testFormField1Value != "value1" || testFormField2Value != "value2" {
		t.Errorf("form data:\nexpect \"value1\" for \"field1\", get %q\nexpect \"value1\" for \"field1\", get %q", testFormField1Value, testFormField2Value)
	}

}

func TestFormEncoded(t *testing.T) {

	input := "field1=value1_1&field1=value1_2&field2=value2"
	req := httptest.NewRequest(http.MethodPost, "/form", strings.NewReader(input))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	v1, v2 := req.FormValue("field1"), req.FormValue("field2")
	if v1 != "value1_1" { // second value of field1 is ignored
		t.Errorf("%s", v1)
	}
	if v2 != "value2" {
		t.Errorf("%s", v2)
	}
}

func TestFormDataEncode(t *testing.T) {

	formData := url.Values{
		"field1": {"value1_1", "value1_2"},
		"field2": {"value2"},
	}

	req := httptest.NewRequest(http.MethodPost, "/form", strings.NewReader(formData.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// populates http.request.Form and http.request.PostForm fields
	if err := req.ParseForm(); err != nil {
		t.Fatalf("failed to parse form data: %v", err)
	}

	// req.Form, req.Postform are of type url.Values (map[string][]string)
	// can consume multiple field values
	if !reflect.DeepEqual(formData, req.PostForm) {
		t.Errorf("parsing form data:\nget\n%v\nexpect\n%v", req.PostForm, formData)
	}

	if !reflect.DeepEqual(formData, req.Form) {
		t.Errorf("parsing form data:\nget\n%v\nexpect\n%v", req.Form, formData)
	}
}

func TestFormDataContentType(t *testing.T) {

	input := map[string]any{"client_id": "id0123", "token": "0123456789abc"}

	var buf bytes.Buffer
	mp := multipart.NewWriter(&buf)
	for k, v := range input {
		s, ok := v.(string)
		if !ok {
			t.Fatalf("failed convert field value to string:%v", v)
		}
		mp.WriteField(k, s)
	}
	mp.Close()

	req := httptest.NewRequest(http.MethodPost, "/form", &buf)
	req.Header.Set("Content-Type", mp.FormDataContentType()) // "multipart/form-data" + boundary
	resp := httptest.NewRecorder()

	mux := http.NewServeMux()
	mux.Handle("/form", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(1 << 8)
		w.Write([]byte(r.MultipartForm.Value["client_id"][0] + ":" + r.MultipartForm.Value["token"][0]))
	}))
	mux.ServeHTTP(resp, req)

	outputActual, outputExpected := resp.Body.String(), "id0123:0123456789abc"
	if outputActual != outputExpected {
		t.Errorf("get\n%q\nexpect\n%q", outputActual, outputExpected)
	}
}
