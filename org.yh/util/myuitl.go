package util

import (
	"fmt"
)

//首字母大写为共有函数
//2个数换位
func Swap(x *int, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}

//切片是可索引的，并且可以由 len() 方法获取长度。
//切片提供了计算容量的方法 cap() 可以测量切片最长可以达到多少。
func PrintSlice(x []int) {
	fmt.Printf("len = %d cap = %d slice = %v \n", len(x), cap(x), x)
}

//首字母小写为私有函数
func swap1(x *int, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}
