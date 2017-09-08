package main

import (
	"fmt"
)

func main() {
	var j int = 5

	b := func() (int){
		fmt.Println("afadfad")
		return 1
	}

	a := func() (func()) {
		var i int = 10

		return func() {
			fmt.Printf("i, j: %d, %d\n", i, j)
		}
	}() //这里写括号 代码执行这个返回的函数 ,或者用a()()

	b()

	a()
	//a()()

	j *= 2

	a()
}
