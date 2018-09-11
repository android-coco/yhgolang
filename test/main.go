package main

import (
	"net/http"
	"net/http/httputil"
	"fmt"
	"strings"
)

func main() {
	//f := Get("http://www.baidu.com")
	//fmt.Println(f)
	fmt.Println(strings.NewReader("123").Len())
	fmt.Println(strings.NewReader("123").Size())

}
func Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	result, err := httputil.DumpResponse(resp, true)
	resp.Body.Close()
	if err != nil {
		panic(err)
	}
	return string(result)
}
