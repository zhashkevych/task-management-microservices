package jwt

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSetEncriptionKeyFromJWK(t *testing.T) {
	srv := httptest.NewServer(http.FileServer(http.Dir("./")))
	defer srv.Close()

	err := SetEncriptionKeyFromJWK(fmt.Sprintf("%s/test_jwk.json", srv.URL), "1")
	if err != nil {
		t.Fatal(err)
	}

	if cfg.Encryption.Key != base64.URLEncoding.EncodeToString([]byte("secret_key_jwk")) {
		t.Fatalf("encription key was not set")
	}

	if cfg.Encryption.KeyId != "1" {
		t.Fatalf("encription key id was not set")
	}
}