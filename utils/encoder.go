package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

func EncodeBase64(secretKey string, signedFieldNames string) string {
	// Create a new HMAC by defining the hash type and the key (as byte array)
	key := []byte(secretKey)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(signedFieldNames))

	hmacHash := h.Sum(nil)

	return base64.StdEncoding.EncodeToString(hmacHash)

}

func DecodeBase64(encodedString string) string {
	data, err := base64.StdEncoding.DecodeString(encodedString)
	if err != nil {
		return ""
	}
	return string(data)
}

func VerifySignature(secretKey string, signedFieldNames string, signature string) bool {
	// Create a new HMAC by defining the hash type and the key (as byte array)
	key := []byte(secretKey)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(signedFieldNames))

	hmacHash := h.Sum(nil)

	// Encode the resulting hash to base64
	encodedHash := base64.StdEncoding.EncodeToString(hmacHash)

	// Compare the generated hash with the signature
	return encodedHash == signature
}
