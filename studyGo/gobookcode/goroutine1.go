package main

import (
	"fmt"
	"runtime"
	"time"
)

//不过设计上我们要遵循：不要通过共享来通信，而要通过通信来共享。
func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched() //表示让CPU把时间片让给别人
		fmt.Println(s)
	}
}

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total //把结果发送出去
}

func main() {
	runtime.GOMAXPROCS(3)
	go say("hello")
	say("world")

	a := []int{7, 2, 8, -9, 4, 0}
	//无缓冲channel是在多个goroutine之间同步很棒的工具。
	c := make(chan int)
	go sum(a[:len(a)/2], c) //17
	go sum(a[len(a)/2:], c) //-5
	x, y := <-c, <-c
	fmt.Println(x, y, x+y) //输出顺序不定
	//创建了可以存储4个元素的bool 型channel。在这个channel 中，前4个元素可以无阻塞的写入。
	//当写入第5个元素时，代码将会阻塞，直到其他goroutine从channel 中读取一些元素，腾出空间。
	c_buff := make(chan bool, 2)

	c_buff <- true
	c_buff <- false

	fmt.Println(<-c_buff)
	fmt.Println(<-c_buff)

	//for i := range c能够不断的读取channel里面的数据，直到该channel被显式的关闭
	c1 := make(chan int, 10)
	go fibonacci(cap(c1), c1)
	for i := range c1 {
		fmt.Println(i)
	}

	//select
	c2 := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c2)
		}
		quit <- 0
	}()
	fibonacci1(c2, quit)

	//用select设置超时
	c3 := make(chan int)
	o := make(chan bool)
	go func() {
		for {
			select {
			case v, _ := <-c3:
				fmt.Println(v)
			case <-time.After(5 * time.Second):
				fmt.Println("timeout")
				o <- true //通知超时了
				return    //这里可以用break
			}
			fmt.Println("go func()for")
		}
	}()
	<-o //取出是否超时
}
func fibonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	//记住应该在生产者的地方关闭channel，而不是消费的地方去关闭它，这样容易引起panic
	//另外记住一点的就是channel不像文件之类的，不需要经常去关闭，
	//只有当你确实没有任何发送数据了，或者你想显式的结束range循环之类的
	close(c)
}

func fibonacci1(c, quit chan int) {
	x, y := 1, 1
	//select默认是阻塞的，只有当监听的channel中有发送或接收可以进行时才会运行
	//当多个channel都准备好的时候，select是随机的选择一个执行的
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return //这里不能用break(break 只跳出select这个循环)
		default:
			//fmt.Println("default")
			//在select里面还有default语法，
			//select其实就是类似switch的功能，
			//default就是当监听的channel都没有准备好的时候，默认执行的（select不再阻塞等待channel）。
		}
		fmt.Println("fibonacci1 for")
	}
}
