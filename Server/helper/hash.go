package helper

import (
	"crypto/sha512"
	"encoding/base64"
)

// NewTokenHash will hash a token string with SHA512 and return the base64 encoded string of it
func NewTokenHash(token string) string {
	hash := sha512.New()
	hash.Write([]byte(token))

	return base64.StdEncoding.EncodeToString(hash.Sum(nil))
}
