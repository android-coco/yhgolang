package main

import (
	"fmt"
	"strings"
)

//仔细分析这段代码
func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice1(s) //6 6 [2,3,5,6,11,13]

	// Slice the slice to give it zero length.
	s = s[:0] //0 6 []
	printSlice1(s)

	// Extend its length.
	s = s[:4] //4 6  [2,3,5,7]
	printSlice1(s)

	// Drop its first two values.
	s = s[2:] //2 4 [5,7]
	printSlice1(s)

	fmt.Println(swap1(1,2))

	fmt.Println(len(strings.Fields("aa aa a")))
}

func swap1(i, j int64) (a, b int64) {
	i ^= j
	j ^= i
	i ^= j
	return i, j
}

//切片是可索引的，并且可以由 len() 方法获取长度。
//切片提供了计算容量的方法 cap() 可以测量切片最长可以达到多少。
func printSlice1(x []int) {
	fmt.Printf("len = %d cap = %d slice = %v \n", len(x), cap(x), x)
}
