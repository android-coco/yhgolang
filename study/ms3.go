/*
* @Author: yhlyl
* @Date:   2018-08-30 19:39:45
* @Last Modified by:   yhlyl
* @Last Modified time: 2018-08-30 19:47:28
 */

package main

import (
	"fmt"
	"runtime"
)

type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

// func main() {
// 	t := Teacher{}
// 	t.ShowA()
// }

func main() {
	runtime.GOMAXPROCS(1)
	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)
	int_chan <- 1
	string_chan <- "hello"
	select {
	case value := <-int_chan:
		fmt.Println(value)
	case value := <-string_chan:
		panic(value)
	}
}
