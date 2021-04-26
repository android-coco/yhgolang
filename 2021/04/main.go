/*
 * @Author: yhlyl
 * @Date: 2021-04-26 17:30:39
 * @LastEditTime: 2021-04-26 17:50:36
 * @LastEditors: yhlyl
 * @Description:
 * @FilePath: /yhgolang/2021/04/main.go
 */
package main

import (
	"fmt"
)

var sink chan int

func main() {
	// test1()
	sink = make(chan int, 3)
	sink <- 1
	sink <- 2
	close(sink)
	sink <- 3
	fmt.Println(<-sink)
	fmt.Println(<-sink)
	fmt.Println(<-sink)
}

func test1() {
	sink = make(chan int)
	go Subscribe(sink)
	fmt.Println(Subscribe1(sink))
}

func Subscribe(sink chan<- int) {
	sink <- 2
}

func Subscribe1(sink <-chan int) int {
	return <-sink
}
