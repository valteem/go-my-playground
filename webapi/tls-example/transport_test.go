package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"testing"
)

const (
	port        = ":8443"
	responseMsg = "some simple response"
)

func serve(addr string) {

	mux := http.NewServeMux()
	mux.Handle("/info", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(responseMsg))
		w.Header().Add("Content-Type", "text/plain")
	}))

	http.ListenAndServeTLS(addr, "cert.pem", "key.pem", mux)

}

func newClient() *http.Client {

	certRaw, _ := os.ReadFile("cert.pem")

	block, _ := pem.Decode(certRaw)

	cert, _ := x509.ParseCertificate(block.Bytes)

	certPool := x509.NewCertPool()
	certPool.AddCert(cert)

	config := &tls.Config{RootCAs: certPool}

	transport := &http.Transport{TLSClientConfig: config}

	client := &http.Client{Transport: transport}

	return client

}

func TestTransportTLS(t *testing.T) {

	go serve(port)

	time.Sleep(100 * time.Millisecond) // allow server some time to start properly

	client := newClient()

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://localhost%s/info", port), nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("failed to fetch response: %v", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}

	if body := string(bodyBytes); body != responseMsg {
		t.Errorf("response body: get %q, expect %q", body, responseMsg)
	}

}
