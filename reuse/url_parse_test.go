package reuse_test

import (
	"net"
	"net/http"
	"net/url"
	"reflect"
	"slices"
	"strings"
	"testing"
)

func TestURLParse(t *testing.T) {

	tests := []struct {
		input    string
		path     string
		rawQuery string
		values   map[string][]string
	}{
		{
			input:    `http://example.com:443/collection_name/pp1/resource_name/pp2/subresource_name/pp3?query_param1=qp1&query_param2=qp2&query_param3=qp3`,
			path:     `/collection_name/pp1/resource_name/pp2/subresource_name/pp3`,
			rawQuery: `query_param1=qp1&query_param2=qp2&query_param3=qp3`,
			values:   url.Values{"query_param1": []string{"qp1"}, "query_param2": []string{"qp2"}, "query_param3": []string{"qp3"}},
		},
	}

	for _, tc := range tests {
		req, err := http.NewRequest(http.MethodGet, tc.input, nil)
		if err != nil {
			t.Fatalf("failed to create new request with method %s and URL %q", http.MethodGet, tc.input)
		}
		u := req.URL
		path := u.Path
		if path != tc.path {
			t.Errorf("URL path: get %q, expect %q", path, tc.path)
		}
		rawQuery := u.RawQuery
		if rawQuery != tc.rawQuery {
			t.Errorf("URL query: get %q, expect %q", rawQuery, tc.rawQuery)
		}
		values, err := url.ParseQuery(rawQuery)
		if err != nil {
			t.Errorf("error parsing  raw query: %v", err)
		}
		if !reflect.DeepEqual(values, url.Values(tc.values)) {
			t.Errorf("URL query values:\n%v\n%v", values, tc.values)
		}
		for v := range values {
			if !slices.Equal(values[v], tc.values[v]) {
				t.Errorf("wrong query parameter %s value: get %v, expect %v", v, values[v], tc.values[v])
			}
		}
	}

}

func TestParseRequestURI(t *testing.T) {

	// parsing malformed URL query string
	addr := `http://a.b:80/pp1/pv1?qp1=qv1&qp2=qv2/pp2/pv2`
	rawQueryExpected := `qp1=qv1&qp2=qv2`
	u, err := url.ParseRequestURI(addr)
	if err != nil {
		t.Errorf("error parsing address %s: %v", addr, err)
	}
	// fails, everything after question mark (including slashes) is included in raw query string
	if u.RawQuery != rawQueryExpected {
		t.Errorf("query string is expected to be\n%s\nget\n%s", rawQueryExpected, u.RawQuery)
	}
}

func TestToURL(t *testing.T) {

	addr := "https://stackoverflow.com/questions/63197536/replacing-protocol-and-hostname-in-url-in-go"
	parsedURL := url.URL{
		Scheme: "https",
		Host:   "stackoverflow.com",
		Path:   "/questions/63197536/replacing-protocol-and-hostname-in-url-in-go",
	}
	pathSegments := []string{"questions", "63197536", "replacing-protocol-and-hostname-in-url-in-go"}

	u, _ := url.Parse(addr) // retuns pointer
	if *u != parsedURL {
		t.Errorf("URL representation should be %+v, is %+v", parsedURL, u)
	}

	s := strings.Trim(u.Path, "/") // remove leading and trailing slashes
	p := strings.Split(s, "/")
	if !slices.Equal(p, pathSegments) {
		t.Errorf("wrong split result:\nget\n%v\nexpect%v\n", p, pathSegments)
	}

}

func TestURLQuery(t *testing.T) {

	addr := "https://someuser:somepassword@localhost:44567/?storage=Main&sku=0001&qty=1"
	expected := map[string]string{
		"storage": "Main",
		"sku":     "0001",
		"qty":     "1",
	}

	u, _ := url.Parse(addr)
	q := u.Query()
	for k, v := range expected {
		if q.Get(k) != v { // Get() returns only first value associated with the given key
			t.Errorf("wrong URL query value for key %s: get %s, expect %s", k, q.Get(k), v)
		}
	}

	if u.User.Username() != "someuser" {
		t.Errorf("wrong username %s", u.User.Username())
	}

	if p, _ := u.User.Password(); p != "somepassword" {
		t.Errorf("wrong password %s", p)
	}

}

func TestHostSplit(t *testing.T) {

	addr := "https://localhost:44536"
	u, _ := url.Parse(addr)
	h, p, _ := net.SplitHostPort(u.Host)
	if h != "localhost" {
		t.Errorf("wrong host name: %+v", h)
	}
	if p != "44536" {
		t.Errorf("wrong port number: %+v", p)
	}

	addr = "https://127.0.0.1"
	u, _ = url.Parse(addr)
	h, p, e := net.SplitHostPort(u.Host) // returns empty 'h' and 'p' if port is not specified, throws error
	if e == nil {
		t.Errorf("error should be thrown")
	}
	if h != "" {
		t.Errorf("wrong host name: %+v", h)
	}
	if p != "" {
		t.Errorf("wrong port number: %+v", p)
	}

}

func TestMultipleQueryParamValues(t *testing.T) {

	tests := []struct {
		input  string
		output url.Values
	}{
		{
			input:  "http://example.org:8080/query_params?id=101&id=102&qty=50",
			output: url.Values{"id": {"101", "102"}, "qty": {"50"}},
		},
		{
			input:  "http://example.org:8080/query_params?id=101,102&qty=50",
			output: url.Values{"id": {"101,102"}, "qty": {"50"}},
		},
	}

	for _, tc := range tests {
		u, _ := url.Parse(tc.input)
		var values url.Values // map[string][]string
		values, _ = url.ParseQuery(u.RawQuery)
		if !reflect.DeepEqual(values, tc.output) {
			t.Fatalf("error parsing input\n%s:\nget\n%v\nexpect\n%v", tc.input, values, tc.output)
		}
	}

}
