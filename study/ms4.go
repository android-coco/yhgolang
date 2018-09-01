/*
* @Author: yhlyl
* @Date:   2018-08-30 19:52:20
* @Last Modified by:   yhlyl
* @Last Modified time: 2018-08-30 20:08:54
 */
package main

import (
	"fmt"
	//"sync"
)

// type UserAges struct {
// 	ages map[string]int
// 	sync.Mutex
// }

// func (ua *UserAges) Add(name string, age int) {
// 	ua.Lock()
// 	defer ua.Unlock()
// 	ua.ages[name] = age
// }

// func (ua *UserAges) Get(name string) int {
// 	if age, ok := ua.ages[name]; ok {
// 		return age
// 	}
// 	return -1
// }

func Iter() <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		//set.RLock()

		// for elem := range set.s {
		// 	ch <- elem
		// }
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
		//set.RUnlock()

	}()
	return ch
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main1() {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}

func main() {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)
}
