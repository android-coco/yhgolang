package main

import "fmt"
import "net/http"

func Hello(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Hello, Welcome to go web programming...")
}

func Hello1(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Hello1, Welcome to go web programming...")
}
func main() {
	http.HandleFunc("/a/c", Hello1)
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":1111", nil)
}
