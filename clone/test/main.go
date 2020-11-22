package main

import (
	"compress/gzip"
	"github.com/mozillazg/request"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func test1(url string) (bt []byte, err error) {
	testResp, err := http.Get(url)
	if err != nil {
		return
	}
	bt, err = ioutil.ReadAll(testResp.Body)
	return
}

func test2(url string) (bt []byte, err error) {
	c := new(http.Client)
	req := request.NewRequest(c)
	resp, err := req.Get(url)
	if err != nil {
		return
	}
	bt, err = ioutil.ReadAll(resp.Body)
	return
}

func test3(url string) (bt []byte, err error) {
	c := new(http.Client)
	req := request.NewRequest(c)
	resp, err := req.Get(url)
	if err != nil {
		return
	}
	var reader io.ReadCloser
	if resp.Header.Get("Content-Encoding") == "gzip" {
		reader, err = gzip.NewReader(resp.Body)
		if err != nil {
			return
		}
	} else {
		reader = resp.Body
	}

	bt, err = ioutil.ReadAll(reader)
	return
}

func main() {
	//urls := [...]string{"https://www.baidu.com", "http://www.zhuoyachina.com/"}
	//for i := range urls {
	//	url := urls[i]
	//	fmt.Println("\n----------\n")
	//	fmt.Println("url =", url)
	//
	//	b1, _ := test1(url)
	//	b2, _ := test2(url)
	//	fmt.Println(b1)
	//	fmt.Println(b2)
	//	fmt.Println(reflect.DeepEqual(b1, b2))
	//
	//	fmt.Println("\n----------\n")
	//
	//	b3, _ := test3(url)
	//	fmt.Println(string(b3))
	//	fmt.Println(reflect.DeepEqual(b1, b3))
	//}


	c := &http.Client{}
	req := request.NewRequest(c)
	resp, _ := req.Get("http://www.zhuoyachina.com/")
	defer resp.Body.Close() // **Don't forget close the response body**
	body, _ := ioutil.ReadAll(resp.Body)
	fr, _ := os.Create("request.html")
	fr.Write(body)

	res, _ := http.Get("http://www.zhuoyachina.com/")
	truebody, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	ft, _ := os.Create("get.html")
	ft.Write(truebody)
}
