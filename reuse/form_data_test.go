package reuse_test

import (
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
