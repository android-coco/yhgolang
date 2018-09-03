package main

import (
	"fmt"
	"awesomeProject/mkw/retriever/mock"
	"awesomeProject/mkw/retriever/myreal"
	"time"
)

const url  = "http://www.imooc.com"
type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

func post(poster Poster) {
	poster.Post(url,
		map[string]string{
			"name": "ccmouse",
			"course": "golang",
	})
}
type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster)  string{
	s.Post(url,map[string]string{
		"contents":"another faked imooc.com",
	})
	return s.Get(url)
}

func download(r Retriever) string {
	return r.Get(url)
}

func main() {
	var r Retriever
	mockR := mock.Retriever{Contents: "this is a fake imooc.com"}
	inspect(&mockR)
	r = &myreal.Retriever{
		UserAgent: "Mozilla/5.0",
		TimOut:    time.Minute,
	}
	inspect(r)
	// Type assertion
	if realRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(realRetriever.Contents)
	} else {
		fmt.Println("not a mock retriever")
	}
	fmt.Println("Tyr a session")
	//fmt.Println(download(r))
	fmt.Println(session(&mockR))

}

func inspect(r Retriever) {
	fmt.Println("Inspectiong",r)
	fmt.Printf(" > %T %v\n", r, r)
	fmt.Print(" > Type switch:")
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *myreal.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
	fmt.Println()
}
