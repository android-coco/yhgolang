/*
 * @Author: yhlyl
 * @Date: 2020-01-29 22:13:09
 * @LastEditTime : 2020-01-30 17:38:59
 * @LastEditors  : yhlyl
 * @Description:
 * @FilePath: /yhgolang/select/select1/mian.go
 * @https://github.com/android-coco
 */
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		c2 := make(chan string)
		go func(c2 chan string) {
			defer wg.Done()
			time.Sleep(1 * time.Millisecond)
			c2 <- "11"
		}(c2)
		T(c2)
		fmt.Println("====",i)
	}
	wg.Wait()
	return
	ch := make(chan int)
	quit := make(chan int)
	go fibonacci(ch, quit)
	quit <- 0
	time.Sleep(2 * time.Second)
	//非阻塞的收发
	ch = make(chan int)
	select {
	case i := <-ch:
		println(i)
	default:
		println("default")
	}
	//随机执行
	ch = make(chan int)
	go func() {
		for range time.Tick(1 * time.Second) {
			ch <- 0
		}
	}()

	for {
		select {
		case <-ch:
			println("case1")

		case <-ch:
			println("case2")
		}
	}
}

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func test1() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		c2 <- "video data"
	}()

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "audio data"

		time.Sleep(time.Second * 1)
		c2 <- "video data"
	}()

	for i := 0; i < 3; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("Received:", msg1)
		case msg2 := <-c2:
			fmt.Println("Received:", msg2)
		}
	}

}
func T(c2  chan string)  {
	timeTimer := time.NewTimer(200 * time.Millisecond)
	select {
	case data := <-c2:
		fmt.Println("=====",data)
	case <-timeTimer.C:
		fmt.Println("超时")
	}
}