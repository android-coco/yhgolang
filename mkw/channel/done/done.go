package main

import (
	"fmt"
	"sync"
	"time"
)

func doWorke(id int, w worker) {
	for n := range w.in {
		fmt.Printf("Worker %d received %c\n",
			id, n)
		w.done()
	}

}

type worker struct {
	in chan int
	//done chan bool
	done func()
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		//done: make(chan bool),
		done: func() {
			wg.Done()
		},
	}
	go doWorke(id, w)
	return w
}
func chanDemo() {
	//var c chan int // c == nil
	//c := make(chan int)
	var workers [10]worker
	//等待多人 完成任务
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}
	wg.Add(20)
	//发数据
	for i, worker := range workers {
		worker.in <- 'a' + i
	}
	// wait for all of them
	//for _, worker := range workers {
	//	<-worker.done
	//}
	//发数据
	for i, worker := range workers {
		worker.in <- 'A' + i
	}

	wg.Wait()
	// wait for all of them
	//for _, worker := range workers {
	//	<-worker.done
	//}
}
func main() {
	chanDemo()
	fmt.Println(time.AfterFunc(time.Second, func() {
		fmt.Println("1111")
	}))
	time.Sleep(time.Second * 2)
}
