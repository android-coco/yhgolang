package main

import (
	"fmt"
	//"time"
	"runtime"
)

var myChan chan bool

//Channel 共享内存
func main() {
	var chan_test chan interface{}    //双项通道
	var read_test <-chan interface{}  //单项通道 只支持 读
	var write_test chan<- interface{} //单项通道 只支持 写
	//read_test <- "xiao"               //会报错。因为这个通道 只是单项通道只支持读 不支持写
	//<-write_test                      //会报错。因为这个通道 只是单项通道只支持写 不支持读
	fmt.Println(chan_test)
	fmt.Println(read_test)
	fmt.Println(write_test)
	//time.Sleep(2 * time.Second) //休眠2秒
	Test1() //无缓存 堵塞
	Test2() //无缓存 堵塞
	Test3() //有缓存 异步

	runtime.GOMAXPROCS(runtime.NumCPU()) //使用电脑的多核
	//输出次数不定
	c := make(chan bool)
	for i := 0; i < 10; i++ {
		go Test4(c, i)
	}
	<-c

	//输出10次
	c1 := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		go Test5(c1, i)
	}
	for i := 0; i < 10; i++ {
		<-c1
	}
}

func Test1() {
	c := make(chan bool) //双项通道 无缓存
	go func() {
		fmt.Println("GO GO GO!!!!")
		c <- true //存
	}()
	<-c //取
}

func Test2() {
	c := make(chan bool) //双项通道无缓存
	go func() {
		fmt.Println("GO GO GO!!!!")
		c <- true //存
		//关闭
		close(c)
	}()
	for v := range c {
		fmt.Println(v)
	}
}

func Test3() {

	c := make(chan bool, 1) //双项通道 有缓存
	go func() {
		fmt.Println("GO GO GO!!!!")
		c <- true //存
	}()
	<-c //取
}

func Test4(c chan bool, index int) {
	a := 1
	for i := 0; i < 100000000; i++ {
		a += i
	}

	fmt.Println(index, a)

	if index == 9 {
		c <- true
	}
}

func Test5(c chan bool, index int) {
	a := 1
	for i := 0; i < 100000000; i++ {
		a += i
	}

	fmt.Println(index, a)

	c <- true
}

func Go() {
	fmt.Println("GO GO GO!!!!")
}
