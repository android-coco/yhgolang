package myreal

import (
	"time"
	"net/http"
	"net/http/httputil"
)

type Retriever struct {
	UserAgent string
	TimOut time.Duration
}

func (r *Retriever) Get(url string) string {
	resp,err := http.Get(url)
	if err != nil {
		panic(err)
	}
	result,err :=httputil.DumpResponse(resp,true)
	resp.Body.Close()
	if err != nil {
		panic(err)
	}
	return string(result)
}
