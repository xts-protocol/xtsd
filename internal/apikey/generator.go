package apikey

import (
	"crypto/rand"
	"encoding/hex"
)

func Generate() string {
	buf := make([]byte, 16)

	_, err := rand.Read(buf)
	if err != nil {
		return ""
	}

	return "xts_" + hex.EncodeToString(buf)
}
