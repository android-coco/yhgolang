// hello.txt
//hello, go 1.16

// main.go
package main

import (
	_ "embed"
	"fmt"
	"math/rand"
	"time"
)

//--go:embed hello.txt
var s string

func main() {
	//http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//	w.Write([]byte(s))
	//}))
	//http.ListenAndServe(":8181", nil)
	//2021-04-07
	//fmt.Println(CreateCaptcha())
	dt := time.Date(2023, 1, 10, 0, 0, 1, 100, time.Local)
	fmt.Println(dt.After(time.Now()))     // true
	fmt.Println(time.Now().Before(dt))    // false
}
func CreateCaptcha() string {
	return fmt.Sprintf("%08v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(100000000))
}