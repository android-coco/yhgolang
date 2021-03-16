// hello.txt
//hello, go 1.16

// main.go
package main

import (
	_ "embed"
	"net/http"
)

//go:embed hello.txt
var s string

func main() {
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(s))
	}))
	http.ListenAndServe(":8181", nil)
}
