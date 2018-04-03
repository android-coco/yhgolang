package main

import (
	"fmt"
	"runtime"
	"sync"
)

const N = 26

func main() {
	// const GOMAXPROCS = 1
	// runtime.GOMAXPROCS(GOMAXPROCS)

	// var wg sync.WaitGroup
	// wg.Add(2 * N)
	// for i := 0; i < N; i++ {
	// 	go func(i int) {
	// 		defer wg.Done()
	// 		// A

	// 		fmt.Printf("%c", 'a'+i)
	// 	}(i)
	// 	continue
	// 	go func(i int) {
	// 		defer wg.Done()
	// 		fmt.Printf("%c", 'A'+i)
	// 	}(i)
	// }
	// wg.Wait()

	runtime.GOMAXPROCS(1)
	waitGourp := sync.WaitGroup{}
	waitGourp.Add(20)
	for i := 0; i < 10; i++ {
		go Go(&waitGourp, i)
		go Go(&waitGourp, i)

	}
	waitGourp.Wait()
}
func Go(wg *sync.WaitGroup, index int) {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(index, sum)
	wg.Done()
}
