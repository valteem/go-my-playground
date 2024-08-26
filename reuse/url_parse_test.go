package reuse_test

import (
	"net/http"
	"net/http/httptest"
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
			values:   map[string][]string{"query_param1": []string{"qp1"}, "query_param2": []string{"qp2"}, "query_param3": []string{"qp3"}},
		},
	}

	for _, tc := range tests {
		req := httptest.NewRequest(http.MethodGet, tc.input, nil)
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
