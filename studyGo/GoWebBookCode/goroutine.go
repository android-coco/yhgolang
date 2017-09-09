package main

import (
	"fmt"
	"runtime"
	"time"
)

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum // send sum to c
}

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched() //表示让CPU把时间片让给别人，下次某个时候继续恢复执行该goroutine。
		fmt.Println(s)
	}
}

func fibonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func fibonacci1(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		default:
			// 当c阻塞的时候执行这里
		}
	}
}
func main() {
	go say("world") //开一个新的Goroutines执行
	say("hello")    //当前Goroutines执行

	a := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)

	//
	c1 := make(chan int, 10)
	go fibonacci(cap(c1), c1)
	for i := range c1 {
		fmt.Println(i) //1,1,2,3,5,8,13,21,34,55
	}

	//select
	c2 := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c2) //1,1,2,3,5,8,13,21,34,55
		}
		quit <- 0 //quit
	}()
	fibonacci1(c2, quit)

	//超时
	c3 := make(chan int)
	o := make(chan bool)
	go func() {
		for {
			select {
			case v := <-c3:
				fmt.Println(v)
			case <-time.After(5 * time.Second):
				fmt.Println("timeout")
				o <- true
				break
			}
		}
	}()
	<-o
}
