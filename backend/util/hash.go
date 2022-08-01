package util

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
	"math/rand"
)

func HashGenerateSha256(s string) string {
    b := []byte(s)
	sha256 := sha256.Sum256(b)

	return hex.EncodeToString(sha256[:])
}

func CodeGenerate() string {
	rand.Seed(time.Now().UnixNano())
	i := 0
	cha := "0123456789"
	paw := ""

	for i < 6 {
		a := rand.Intn(10)
		paw += string(cha[a])
		i++
	}
	
	return paw
}
