// Example code for Chapter 2.7 from "Build Web Application with Golang"
// Purpose: Shows how to close and interate through a channel
package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	//cap 分配的长度   len 实际长度
	c := make(chan int, 10)
	a := make([]int, 8, 10)
	a = []int{1,2,3,4}
	fmt.Println(cap(a),len(a))
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}

}
