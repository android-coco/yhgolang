/*
 * @Author: yhlyl
 * @Date: 2018-09-07 00:13:34
 * @LastEditTime: 2020-02-25 21:58:25
 * @LastEditors: yhlyl
 * @Description:
 * @FilePath: /yhgolang/helloworlddonc/conc.go
 * @https://github.com/android-coco
 */
package main

import (
	"fmt"
)

func main() {
	f1 := f()
	fmt.Println(f1())
	fmt.Println(f1())
	fmt.Println(f1())
	return
	ch := make(chan string)
	for i := 0; i < 5000; i++ {
		go printHellWrold(i, ch)
	}

	for {
		msg := <-ch
		fmt.Println(msg)
	}

}

func printHellWrold(i int, ch chan string) {
	for {
		ch <- fmt.Sprintf("Hello world %d!\n", i)
	}
}

func f() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
