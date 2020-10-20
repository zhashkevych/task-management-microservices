package jwt

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type keys struct {
	Keys []key `json:"keys"`
}

type key struct {
	KeyID     string `json:"kid"`
	SecretKey string `json:"k"`
}

func SetEncriptionKeyFromJWK(url, keyId string) error {
	keyData, err := parseJWKFile(url)
	if err != nil {
		return err
	}

	return setConfigEncryption(keyData, keyId)
}

func parseJWKFile(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func setConfigEncryption(data []byte, keyId string) error {
	var keys keys
	err := json.Unmarshal(data, &keys)
	if err != nil {
		return err
	}

	for _, k := range keys.Keys {
		if k.KeyID == keyId {
			cfg.Encryption.Key = k.SecretKey
			cfg.Encryption.KeyId = keyId

			return nil
		}
	}

	return errors.New("encription key was not set")
}
