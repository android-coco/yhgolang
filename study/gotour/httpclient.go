package main

import (
	"net/http"
	"os"
	"io"
	"fmt"
)

func main() {
	clinet := &http.Client{}
	url := "http://www.baidu.com"
	request,err := http.NewRequest("GET",url,nil)

	if err != nil {
		panic(err)
	}

	response,_:= clinet.Do(request)

	//将结果定位到标准输出 也可以直接打印出来 或者定位到其他地方进行相应的处理
	stdout := os.Stdout
	_, err = io.Copy(stdout, response.Body)

	//返回的状态码
	status := response.StatusCode

	fmt.Println(status)
}
