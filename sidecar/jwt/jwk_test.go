package jwt

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSetEncriptionKeyFromJWK(t *testing.T) {
	srv := httptest.NewServer(http.FileServer(http.Dir("./")))
	defer srv.Close()

	err := SetEncriptionKeyFromJWK(fmt.Sprintf("%s/test_jwk.json", srv.URL), "sim2")
	if err != nil {
		t.Fatal(err)
	}

	if cfg.Encryption.Key != "secret_key_jwk" {
		t.Fatalf("encription key was not set")
	}
}