package main

import (
	"fmt"
	"awesomeProject/mkw/retriever/mock"
	"awesomeProject/mkw/retriever/real"
	"time"
)

const url = "http://www.imooc.com"

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get(url)
}

type Poster interface {
	Post(url string, form map[string]string) string
}

func post(poster Poster) string {
	return poster.Post(url,
		map[string]string{
			"name":   "ccmouse",
			"course": "golang",
		})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another faked imooc.com",
	})
	return s.Get(url)
}

func main() {
	mockR := &mock.Retriever{Contents: "this is a fake imooc.com"}
	inspect(mockR)
	relR := &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimOut:    time.Minute,
	}
	inspect(relR)
	// Type assertion
	if realRetriever, ok := Retriever(relR).(*mock.Retriever); ok {
		fmt.Println(realRetriever.Contents)
	} else {
		fmt.Println("not a mock retriever")
	}
	fmt.Println("Tyr a session")
	//fmt.Println(download(mockR))
	fmt.Println(session(mockR))
	fmt.Println(post(mockR))

}

func inspect(r Retriever) {
	fmt.Println("Inspectiong", r)
	fmt.Printf(" > %T %v\n", r, r)
	fmt.Println(" > Type switch:")
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
	fmt.Println()
}
