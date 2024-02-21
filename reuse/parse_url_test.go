package reuse

import (
	"net/url"
	"strings"

	"testing"
)

func TestToURL(t *testing.T) {
	addr := "https://stackoverflow.com/questions/63197536/replacing-protocol-and-hostname-in-url-in-go"
	parsedURL := url.URL{
		Scheme: "https",
		Host:   "stackoverflow.com",
		Path:   "/questions/63197536/replacing-protocol-and-hostname-in-url-in-go",
	}
	u, _ := url.Parse(addr) // retuns pointer
	if *u != parsedURL {
		t.Errorf("URL representation should be %+v, is %+v", parsedURL, u)
	}

	s := strings.Trim(u.Path, "/") // remove leading and trailing slashes
	p := strings.Split(s, "/")
	if p[0] != "questions" {
		t.Errorf("wrong split result")
	}
}
