package main

import (
	"bytes"
	"crypto/des"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
)

func main() {
	key := []byte("2fa6c1e9")
	str := "I love this beautiful world!"
	strEncrypted, err := DesEncrypt(str, key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Encrypted:", strEncrypted)
	strDecrypted, err := DesDecrypt(strEncrypted, key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Decrypted:", strDecrypted)
}

//Output:
//Encrypted: 5d2333b9fbbe5892379e6bcc25ffd1f3a51b6ffe4dc7af62beb28e1270d5daa1
//Decrypted: I love this beautiful world!

func DesEncrypt(text string, key []byte) (string, error) {
	src := []byte(text)
	block, err := des.NewCipher(key)
	if err != nil {
		return "", err
	}
	bs := block.BlockSize()
	src = ZeroPadding(src, bs)
	if len(src)%bs != 0 {
		return "", errors.New("need a multiple of the blocksize")
	}
	out := make([]byte, len(src))
	dst := out
	for len(src) > 0 {
		block.Encrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
	//hex.EncodeToString(out)
	return base64.StdEncoding.EncodeToString(out), nil
}

func DesDecrypt(decrypted string, key []byte) (string, error) {
	//hex.DecodeString(decrypted)
	src, err := base64.StdEncoding.DecodeString(decrypted)
	if err != nil {
		return "", err
	}
	block, err := des.NewCipher(key)
	if err != nil {
		return "", err
	}
	out := make([]byte, len(src))
	dst := out
	bs := block.BlockSize()
	if len(src)%bs != 0 {
		return "", errors.New("crypto/cipher: input not full blocks")
	}
	for len(src) > 0 {
		block.Decrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
	out = ZeroUnPadding(out)
	return string(out), nil
}

func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func ZeroUnPadding(origData []byte) []byte {
	return bytes.TrimFunc(origData,
		func(r rune) bool {
			return r == rune(0)
		})
}
