package reuse_test

import (
	"regexp"
	"strconv"
	"sync"
	"testing"
)

// Adopted from https://github.com/benhoyt/go-routing/blob/master/reswitch/route.go
func match(path, pattern string, vars ...any) bool {
	regex := mustCompileCached(pattern)
	matches := regex.FindStringSubmatch(path)
	if len(matches) <= 0 {
		return false
	}
	for i, match := range matches[1:] {
		switch p := vars[i].(type) {
		case *string:
			*p = match
		case *int:
			n, err := strconv.Atoi(match)
			if err != nil {
				return false
			}
			*p = n
		default:
			panic("vars must be *string or *int")
		}
	}
	return true
}

var (
	regexen = make(map[string]*regexp.Regexp)
	relock  sync.Mutex
)

func mustCompileCached(pattern string) *regexp.Regexp {
	relock.Lock()
	defer relock.Unlock()

	regex := regexen[pattern]
	if regex == nil {
		regex = regexp.MustCompile("^" + pattern + "$")
		regexen[pattern] = regex
	}
	return regex
}

func TestReSwitch(t *testing.T) {
	tests := []struct {
		desc    string
		path    string
		pattern string
		match   bool
	}{
		{
			desc:    "one param",
			path:    "/api/stuff",
			pattern: "/api/[^/]+",
			match:   true,
		},
		{
			desc:    "two params instead of just one",
			path:    "/api/some/stuff",
			pattern: "/api/[^/]+",
			match:   false,
		},
		{
			desc:    "two params",
			path:    "/api/some/stuff",
			pattern: "/api/([^/]+)/([^/]+)",
			match:   true,
		},
		{
			desc:    "one param instead of two",
			path:    "/api/stuff",
			pattern: "/api/([^/]+)/([^/]+)",
			match:   false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			var param1, param2 string
			if m := match(tc.path, tc.pattern, &param1, &param2); m != tc.match {
				t.Errorf("match(): expect %t, get %t", m, tc.match)
			}
		})
	}
}
