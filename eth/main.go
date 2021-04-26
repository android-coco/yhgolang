package main

import (
	"fmt"
	"golang.org/x/crypto/sha3"
	"strings"
)

func TextHash(data []byte) []byte {
	hash, _ := TextAndHash(data)
	return hash
}
func TextAndHash(data []byte) ([]byte, string) {
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), string(data))
	hasher := sha3.NewLegacyKeccak256()
	hasher.Write([]byte(msg))
	return hasher.Sum(nil), msg
}

func main() {
	fmt.Println(strings.Compare("Jason","As"))
}
