package main

import (
	"fmt"
)

func main() {
	var fs = [4]func(){} //定义的切片  元素为函数类型
	for i := 0; i < 4; i++ {
		defer fmt.Println("defer i = ", i)                           //3,2,1,0
		defer func() { fmt.Println("defer_closure1 i=", i) }()       //4,4,4,4
		defer func(i int) { fmt.Println("defer_closure2 i=", i) }(i) //3,2,1,0
		fs[i] = func() {
			fmt.Println("closure i = ", i) //4,4,4,4
		}
	}

	for _, f := range fs {
		f()
	}
}
