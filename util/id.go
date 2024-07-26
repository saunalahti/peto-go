package util

import (
	"crypto/sha1"
	"encoding/hex"
)

func GenerateID(location string, datetime string) string {
	hasher := sha1.New()
	hasher.Write([]byte(location + datetime))
	return hex.EncodeToString(hasher.Sum(nil))
}
