package main

import (
	"fmt"
	"time"
	"sync"
)

type atomicInt struct {
	value int
	lock  sync.Mutex
	coun sync.Cond
}

func (a *atomicInt) increment() {
	fmt.Println("safe increment")
	func(){
		a.lock.Lock()
		defer a.lock.Unlock()
		a.value++
	}()

}

func (a *atomicInt) get() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return int(a.value)
}

func main() {
	var a atomicInt
	a.increment()
	go func() {
		a.increment()
	}()
	time.Sleep(time.Millisecond)
	fmt.Println(a.get())
}
