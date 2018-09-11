package main

import (
	"time"
	"fmt"
)

func worke(id int, c chan int) {
	for i := range c {
		//n, ok := <-c
		//if !ok {
		//	break
		//}
		fmt.Printf("Worker %d received %d\n",
			id, i)
	}

}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worke(id, c)
	return c
}
func chanDemo() {
	//var c chan int // c == nil
	//c := make(chan int)
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}
	//发数据
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	//发数据
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
	time.Sleep(time.Millisecond)
	//c <- 1
	//c <- 2

}

func bufferedChannel() {
	c := make(chan int, 3)
	go worke(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	time.Sleep(time.Millisecond)
}

func channelClose() {
	c := make(chan int, 3)
	go worke(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)
	time.Sleep(time.Millisecond)
}
func main() {
	fmt.Println("Channel as first-class citizen")
	chanDemo()
	fmt.Println("Buffered channel")
	bufferedChannel()
	fmt.Println("Channel close and range")
	channelClose()
}
