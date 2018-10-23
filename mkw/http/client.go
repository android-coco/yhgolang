package main

import (
	"net/http"
	"net/http/httputil"
	"fmt"
)

const url = "https://www.imooc.com"

func main() {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}

	request.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_1 like Mac OS X) AppleWebKit/603.1.30 (KHTML, like Gecko) Version/10.0 Mobile/14E304 Safari/602.1")

	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("Redirect:",req)
			return nil
		},
	}

	//resp, err := http.DefaultClient.Do(request)
	resp, err := client.Do(request)

	//resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bytes, e := httputil.DumpResponse(resp, true)
	if e != nil {
		panic(err)
	}
	fmt.Printf("%s\n", bytes)
}
