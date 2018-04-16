// Example code for Chapter 2.7 from "Build Web Application with Golang"
// Purpose: Shows how to create and use a timeout
package main

import (
	"fmt"
	"time"
)

func add(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
}
func main() {
	//1,每个case都必须是一个通信
	//2,所有channel表达式都会被求值
	//3,所有被发送的表达式都会被求值
	//4,如果任意某个通信可以进行，它就执行；其他被忽略。
	//5,如果有多个case都可以运行，Select会随机公平地选出一个执行。其他不会执行。
	//6,否则：如果有default子句，则执行该语句。
	//7,如果没有default字句，select将阻塞，直到某个通信可以运行；Go不会重新对channel或值进行求值。

	//test2()
	//test3()
	test4()
}

//根据条件5：如果有多个case都可以运行，Select会随机公平地选出一个执行。其他不会执行。
//但是运行上述代码，当ch通道中存在数据时，time.After总是得不到运行，因此到时超时未生效（就像是两个case都成立时，select 都”公平”地选择了 case <- ch，导致超时逻辑未生效）
func test2() {
	ch := make(chan int, 10)
	go add(ch)
	for {
		select {
		case <-time.After(2 * time.Second):
			fmt.Println("timeout")
			return
		case i := <-ch:
			fmt.Println(i) // if ch not empty, time.After will nerver exec
			fmt.Println("sleep one seconds ...")
			time.Sleep(1 * time.Second)
			fmt.Println("sleep one seconds end...")
		}
	}
}

// 随机性失败
//当case <- ch 和 case <- ticker.C 同时成立时，Select会随机公平地选出一个执行，有可能选择到前者，导致超时随机行失败
func test3() {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	ch := make(chan int, 10)
	go add(ch)
	for {
		select {
		case <-ticker.C:
			fmt.Println("timeout")
			return
		case i := <-ch:
			fmt.Println(i) // if ch not empty, time.After will nerver exec
			fmt.Println("sleep one seconds ...")
			time.Sleep(1 * time.Second)
			fmt.Println("sleep one seconds end...")
		}
	}
}

// 最终解决方式
//将【超时】和【收包】放在各自单独的select里面，【超时】一定可以执行到
func test4() {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	ch := make(chan int, 10)
	go add(ch)
	for {
		select {
		case i := <-ch:
			fmt.Println(i) // if ch not empty, time.After will nerver exec
			fmt.Println("sleep one seconds ...")
			time.Sleep(1 * time.Second)
			fmt.Println("sleep one seconds end...")
		}
		select {
		case <-ticker.C:
			fmt.Println("timeout")
			return
		}
	}
}
