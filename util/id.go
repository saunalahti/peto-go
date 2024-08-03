package util

import (
	"crypto/sha1"
	"encoding/hex"
)

func GenerateID(location string, datetime string) string {
	hasher := sha1.New()
	hasher.Write([]byte(location + datetime))
	fullHash := hasher.Sum(nil)
	truncatedHash := fullHash[:8] // Taking first 8 bytes of SHA-1 hash
	return hex.EncodeToString(truncatedHash)
}
