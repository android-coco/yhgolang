package main

import (
	"fmt"
	"time"
	"math/rand"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}
func worke(id int, c chan int) {
	for i := range c {
		//n, ok := <-c
		//if !ok {
		//	break
		//}
		time.Sleep(5 * time.Microsecond)
		fmt.Printf("Worker %d received %d\n",
			id, i)
	}

}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worke(id, c)
	return c
}
func main() {
	var c1, c2 chan int // c1 and c2 = nil
	c1, c2 = generator(), generator()

	var values []int
	tm := time.After(10 * time.Second)
	tick := time.Tick(time.Second)
	for {
		var activeWorker chan<- int
		var activeValue int
		if len(values) > 0 {
			activeWorker = createWorker(0)
			activeValue = values[0]
		}
		select {
		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
		case activeWorker <- activeValue:
			values = values[1:]
		case <-time.After(800*time.Millisecond):
			fmt.Println("timeout")
		case <- tick:
			fmt.Println("queue len =",len(values))
		case <-tm:
			fmt.Println("bye")
			return
		}
	}
}
