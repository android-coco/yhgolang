package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"fmt"
	"strings"
)
const Secret = `-----BEGIN PRIVATE KEY-----
MIGTAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBHkwdwIBAQQggYBqw9g3+9LtaYig
PofGXLDPmJ6D/2cOsGqrkVQYq7SgCgYIKoZIzj0DAQehRANCAATdABHmwtiEnXIW
49k8tDo+oCjShaR/2R0bskPS4GPKOXfHF9c1xjqEvcDIf42C/saZLpa/pzr3mfam
oU5UyuoG
-----END PRIVATE KEY-----`
func main() {
	prk, err := ecdsa.GenerateKey(elliptic.P521(),strings.NewReader(Secret))
	if err != nil {
		fmt.Println(err)
	}

	puk := prk.PublicKey
	fmt.Println(prk,puk)
}
