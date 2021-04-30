package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	countG int
	//互斥锁
	countGuard sync.Mutex
)

func read(mapA map[string]string) {
	for {
		countGuard.Lock()
		var _ string = mapA["name"]
		countG += 1
		countGuard.Unlock()
	}
}

func write(mapA map[string]string) {
	for {
		countGuard.Lock()
		mapA["name"] = "johny"
		countG += 1
		time.Sleep(time.Millisecond * 3)
		countGuard.Unlock()
	}
}

func Guard() {
	var num = 3
	var mapA = map[string]string{"neaddma": ""}
	for i := 0; i < num; i++ {
		go read(mapA)
	}

	for i := 0; i < num; i++ {
		go write(mapA)
	}

	time.Sleep(time.Second * 3)
	fmt.Printf("最终读写次数：%d\n", countG)
}

var (
	count  int
	rwLock sync.RWMutex
)

func readRw(mapA map[string]string) {
	for {
		rwLock.RLock() //这里定义了一个读锁
		var _ string = mapA["name"]
		count += 1
		rwLock.RUnlock()
	}
}

func writeRW(mapA map[string]string) {
	for {
		rwLock.Lock() //这里定义了一个写锁
		mapA["name"] = "johny"
		count += 1
		time.Sleep(time.Millisecond * 3)
		rwLock.Unlock()
	}
}

func Rw() {
	var num = 3
	var mapA = map[string]string{"nema": ""}

	for i := 0; i < num; i++ {
		go readRw(mapA)
	}

	for i := 0; i < num; i++ {
		go writeRW(mapA)
	}

	time.Sleep(time.Second * 3)
	fmt.Printf("最终读写次数：%d\n", count)
}

func RwLockDemo() {
	for i := 0; i < 2; i++ {
		go func() {
			for i := 1000000; i > 0; i-- {
				rwLock.Lock()
				count++
				rwLock.Unlock()
			}
			fmt.Println(count)
		}()
	}

	fmt.Scanf("\n") //等待子线程全部结束
}

func main() {
	Guard()
	Rw()

}
