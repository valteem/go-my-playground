package reuse

import (
	"net"
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
