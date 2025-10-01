package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func generateBytes(n int64) (string, error) {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

func main() {
	b, err := generateBytes(16)
	if err != nil {
		panic(err)
	}
	fmt.Println(b)
}
