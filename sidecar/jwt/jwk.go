package jwt

import (
	"encoding/base64"
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

// SetEncriptionKeyFromJWK parses remote JWK File from url
// And sets Key and KeyID values for private package-level config instance
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
	keys, err := parseKeys(data)
	if err != nil {
		return err
	}

	return setEncryptionConfigFromKeys(keys, keyId)
}

func parseKeys(jwkData []byte) (keys, error) {
	var keys keys
	err := json.Unmarshal(jwkData, &keys)
	return keys, err
}

func setEncryptionConfigFromKeys(k keys, keyId string) error {
	for _, k := range k.Keys {
		if k.KeyID == keyId {
			// github.com/dgrijalva/jwt-go lib does not encode secret to base64
			// but KrakenD assumes that secret is encoded
			// so we need to do that manually
			cfg.Encryption.Key = base64.URLEncoding.EncodeToString([]byte(k.SecretKey))
			cfg.Encryption.KeyId = keyId

			return nil
		}
	}

	return errors.New("encription key was not set")
}