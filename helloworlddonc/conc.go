package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)
	for i := 0; i < 5000; i++ {
		go printHellWrold(i, ch)
	}

	for {
		msg := <-ch
		fmt.Println(msg)
	}


}

func printHellWrold(i int, ch chan string) {
	for{
		ch <- fmt.Sprintf("Hello world %d!\n", i)
	}
}
