package main

import (
	"fmt"
)

var c chan string

func Pinpang() {
	i := 0
	for {
		fmt.Println(<-c)
		c <- fmt.Sprintf("From Pinpang: Hi,#%d", i)
		i++
	}
}

func main() {
	c := make(chan string)
	go Pinpang()

	for i := 0; i < 10; i++ {
		c <- fmt.Sprintf("From Main: Hello,#%d", i)
		fmt.Println(<-c)
	}
}
