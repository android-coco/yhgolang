package main

import (
	"fmt"
	"net/http"
)
var url = "https://open.weixin.qq.com/connect/oauth2/authorize?appid=wxdb81c5ac3ddd1a68&redirect_uri=https://test-cbc.ycandyz.com/test?xxx=9999&response_type=code&scope=snsapi_base&state=123#wechat_redirect"
func SayHello(w http.ResponseWriter, req *http.Request) {


	values := req.URL.Query()
	code := values.Get("code")
	xxx := values.Get("xxx")
	fmt.Println(code)
	fmt.Println(xxx)
	//http.Redirect(w,req, url, http.StatusFound)
	//location.href = 'weixin://dl/business/?t=akd7SqtNfpi'
	w.Write([]byte("<html>\n<b>test</b>\n<script>\n    window.onload=function() {\n        window.location = 'weixin://dl/business/?t=akd7SqtNfpi';\n    }\n</script>\n</html>"))
}
func main() {
	http.HandleFunc("/", SayHello)
	http.ListenAndServe(":8888", nil)
}
