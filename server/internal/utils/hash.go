package utils

import (
	"crypto/sha256"
	"encoding/base64"
)

func Hash256(text string) string {
	byteArray := sha256.Sum256([]byte(text))
	hashString := base64.StdEncoding.EncodeToString(byteArray[:])
	return hashString
}
