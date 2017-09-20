package main

import (
	"fit"
	"fmt"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request, _ yhhttp.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps yhhttp.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func main() {
	router := yhhttp.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)

	log.Fatal(http.ListenAndServe(":8181", router))
}
