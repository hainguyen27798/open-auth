package utils

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func CreateSession(size int) string {
	b := make([]byte, size)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("error:", err)
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}
