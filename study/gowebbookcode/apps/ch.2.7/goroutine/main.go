// Example code for Chapter 2.7 from "Build Web Application with Golang"
// Purpose: Shows how to launch a simple gorountine
package main

import (
	"fmt"
	"runtime"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		//让出当前goroutine的执行权限，调度器安排其他等待的任务运行，并在下次某个时候从该位置恢复执行
		runtime.Gosched()
		fmt.Println(s)
	}
}

func main() {
	go say("world") // create a new goroutine
	say("hello")    // current goroutine
}
