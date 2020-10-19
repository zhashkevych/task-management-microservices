package jwt

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseJWK(url string) error {
	keyData, err := parseJWKFile(url)
	if err != nil {
		return err
	}

	return setConfigEncryption(keyData)
}

func parseJWKFile(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func setConfigEncryption(jwkData []byte) error {
	var enc Encryption
	err := json.Unmarshal(jwkData, &enc)
	if err != nil {
		return err
	}

	cfg.Encryption = enc
	return nil
}
