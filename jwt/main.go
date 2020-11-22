package main

import (
	"fmt"
	"github.com/fwhezfwhez/jwt"
)

const Secret = `-----BEGIN PRIVATE KEY-----
MIGTAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBHkwdwIBAQQggYBqw9g3+9LtaYig
PofGXLDPmJ6D/2cOsGqrkVQYq7SgCgYIKoZIzj0DAQehRANCAATdABHmwtiEnXIW
49k8tDo+oCjShaR/2R0bskPS4GPKOXfHF9c1xjqEvcDIf42C/saZLpa/pzr3mfam
oU5UyuoG
-----END PRIVATE KEY-----`
func main() {
	//{
	//"alg": "ES256",
	//"kid": "2X9R4HXF34",
	//"typ": "JWT"
	//}
	//{
	//	"iss": "57246542-96fe-1a63-e053-0824d011072a",
	//	"exp": 1528408800,
	//	"aud": "appstoreconnect-v1"
	//}
	//获取token管理对象
	token := jwt.GetToken()
	//添加令牌关键信息
	token.AddHeader("kid","P96PRNY574").AddHeader("typ", "JWT").AddHeader("alg", "ES256")
	//添加令牌期限
	//exp:=time.Now().Add(1*time.Hour)
	token.AddPayLoad("exp", "1528408800")
	token.AddPayLoad("iss", "c3130f43-07d2-4dbe-90d2-1eb81c399905")
	token.AddPayLoad("aud", "appstoreconnect-v1")
	//获取令牌，并添加进reponse的header里
	jwts, jwets1, erre := token.JwtGenerator(Secret)
	if erre != nil {
		fmt.Println("token生成出错")
		return
	}
	fmt.Println("生成的jwt是1:", jwts)
	fmt.Println("生成的jwt是2:", jwets1)
}
