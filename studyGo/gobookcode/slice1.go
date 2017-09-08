package main

import (
	"fmt"
)

func main() {
	//a := []int{1,2,3}
	//b := []int {4,5,6}
	//c := append(a, b...)
	//
	//fmt.Println(c)

	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}
	copy(slice1, slice2) // 只会复制slice1的前3个元素到slice2中

	fmt.Println(slice1)
}
