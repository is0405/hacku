package util

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashGenerateSha256(s string) string {
    b := []byte(s)
	sha256 := sha256.Sum256(b)

	return hex.EncodeToString(sha256[:])
}
