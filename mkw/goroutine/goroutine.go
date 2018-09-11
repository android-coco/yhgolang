package main

import (
	"fmt"
	"time"
	"runtime"
)
// 协程
//轻量级 "线程"
//非抢占式多任务处理，由协程主动交出控制权
//编译器/解释器/虚拟机层面的多任务
//多个协程可能在一个或多个线程上运行
func main() {
	tryGoroutine1()
	//tryGoroutine2()
}

func tryGoroutine2() {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func(i int) { // race condition !
			for {
				a[i] ++
				//如果不让出给别人运行就死循环
				runtime.Gosched()
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}

func tryGoroutine1() {
	for i := 0; i < 1000; i++ {
		go func(i int) {
			for {
				fmt.Printf("Hello from goroutine %d \n", i)
			}
		}(i)
	}
	time.Sleep(time.Minute)
}
