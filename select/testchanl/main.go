package main

import (
	"fmt"
	"time"
)

func main() {
	var x chan int32
	x = make(chan int32,10000)
	for i := 0; i < 10000; i++ {
		go func() {
			for {
				select {
				case a := <-x:
					fmt.Println("xxxx:", a)
				}
			}
		}()
	}

	for i := 0; i < 10000; i++ {
		go func() {
			for {
				select {
				case a := <-x:
					fmt.Println("yyyy:", a)
				}
			}
		}()
	}

	for {
		time.Sleep(1 * time.Second)
		x <- 1
		x <- 2
	}
}
