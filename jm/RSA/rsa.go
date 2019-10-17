package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
)

// 私钥生成
//openssl genrsa -out rsa_private_key.pem 1024
var privateKey = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIICXAIBAAKBgQCR7v3GU1nTNcTAfoUryN+f7Od8dqbGuKE7bb78jRSsiaOwveYx
wo13cwwD3hLXv7a4JS+or6eXjBQlfhBSueTtSnK7URTrmg9MGPdtmm4UDXtj4eHG
FsoOSOxHhYOMhoc0U91+eEd7Tycuf5vHb2ieYrERrl6+Z+t3AzjNRqRWYQIDAQAB
AoGAIoWubA3IuJHGLyFfAIoe+Lay1js9XdJMdgISxazcQKq42fU6cPgMvj6tj3an
73jvUSWe3iFbnJqrI2lslB1cvbgTOtPWxffC4QrAlq/41NOihASimcUYn58xmzJu
DovvoB+uR2wmZ1Y83Oaid63n0OcfLa8X2r7ZVJWrRiZHQl8CQQDOYQckdb7xB3Pa
Yi/8e8yaQyMQNSjDXZMP85BKJtRM3p8h3JySe8K3And4l3HQAlf0zg4nvfUAugVZ
FaiPss3vAkEAtQVx6K4J7yL2GMJYZirN61u6Qcqgnem3TlF8+9nu/RzrgCZlmgot
7/YZIKvOabfAHk/qkBsFvorgUoyxlL9wrwJAY5HWg1W3qMxCrfM/WZ5VCXwot5Ie
N5u27zRAwjXXqbqiphCtDdNeDzPGdk0C4SuwSfD8TVpNkWsuV1umtqW6ywJATBlo
5Jddp8F70bbJ8NFn1dyu9X+YfCpHnE4Xi8z3ckLZIfuCVPqYiztbHuf7E2hjBJs2
EbS6enrpiOqKsOy2EwJBALYhuJRztHxd9yR1H+n5SuGaYwU9rqzUlWsZUdn2zvPl
V4P3mWVjUiaNOeuwTKu8CrG54N7JhSR9JhYDtI58aQc=
-----END RSA PRIVATE KEY-----
`)

// 公钥: 根据私钥生成
//openssl rsa -in rsa_private_key.pem -pubout -out rsa_public_key.pem
var publicKey = []byte(`
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCR7v3GU1nTNcTAfoUryN+f7Od8
dqbGuKE7bb78jRSsiaOwveYxwo13cwwD3hLXv7a4JS+or6eXjBQlfhBSueTtSnK7
URTrmg9MGPdtmm4UDXtj4eHGFsoOSOxHhYOMhoc0U91+eEd7Tycuf5vHb2ieYrER
rl6+Z+t3AzjNRqRWYQIDAQAB
-----END PUBLIC KEY-----
`)

func main() {
	data, _ := RsaEncrypt([]byte("ceshi01"))
	fmt.Println(base64.StdEncoding.EncodeToString(data))
	bytes, _ := base64.StdEncoding.DecodeString("NXSas2JvcN+e8QhTMaaOE6DPq4riNtXLK7Jf12RkolVI4MdrzkJbbnEQ/t3v1ZqfqF5ehQrgsKhAG7XS9WAxytsrT0X9IAqcSnsRzq9x2PatQC394AaXCEHlje6Tll+1pWacanFa+ROU2yj4iIpRMVHL4oPXnk41+rmUElWYAIA=")
	origData, _ := RsaDecrypt(bytes)
	fmt.Println(string(origData))
}

// 加密
func RsaEncrypt(origData []byte) ([]byte, error) {
	//解密pem格式的公钥
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//加密
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

// 解密
func RsaDecrypt(ciphertext []byte) ([]byte, error) {
	//解密
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error")
	}
	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 解密
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}
