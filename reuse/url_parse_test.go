package reuse_test

import (
	"net/http"
	"net/url"
	"reflect"
	"slices"
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
