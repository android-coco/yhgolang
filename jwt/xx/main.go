package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

const Secret = ``

var header = `{
"alg": "ES256",
"kid": "P96PRNY574",
"typ": "JWT"
}`

var tow = `{
"iss": "c3130f43-07d2-4dbe-90d2-1eb81c399905",
"exp": 1528408800,
"aud": "appstoreconnect-v1"
}`
func main() {
	token := jwt.New(jwt.SigningMethodES256)
	claims := make(jwt.MapClaims)
	claims["exp"] = 1528408800
	claims["iss"] = "c3130f43-07d2-4dbe-90d2-1eb81c399905"
	claims["aud"] = "appstoreconnect-v1"
	token.Claims = claims
	token.Header = map[string]interface{}{
		"kid": "P96PRNY574",
		"typ": "JWT",
		"alg": "ES256",
	}
	//headerJson,err := json.Marshal(token.Header)
	//if err!=nil {
	//	fmt.Println(err)
	//}
	//headerBase64 := base64.StdEncoding.EncodeToString(headerJson)
	//
	//payLoad := token.Claims
	//payLoadJson,err := json.Marshal(payLoad)
	//if err!=nil {
	//	fmt.Println(err)
	//}
	//prk, _ := ecdsa.GenerateKey(elliptic.P521(),strings.NewReader(Secret))
	//payLoadBase64 := base64.StdEncoding.EncodeToString(payLoadJson)
	//
	//sign, err2 := jwt.SigningMethodES256.Sign(fmt.Sprintf("%s.%s",headerBase64,payLoadBase64), prk)
	//if err2 != nil {
	//	fmt.Println("dddd--",err2)
	//}
	//fmt.Println(sign)
	//token.Signature = sign
	pem, err2 := jwt.ParseECPrivateKeyFromPEM([]byte(Secret))
	fmt.Println(pem,err2)
	tokenString, err := token.SignedString(pem)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(tokenString)
}
