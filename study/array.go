package main

import (
	"fmt"
)

//长度为6  不够补0
var array1 = [6]int{1, 2, 3}

//不限定长度
var balance = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
var balance1 = []float64{1000.0, 2.0, 3.4, 7.0, 50.0}

func main() {
	for i, a := range array1 {
		fmt.Printf("第 %d 位 x1 的值 = %d\n", i, a)
	}
	for i, a := range balance {
		fmt.Printf("第 %d 位 x1 的值 = %f\n", i, a)
	}

	for i, a := range balance1 {
		fmt.Printf("第 %d 位 x1 的值 = %f\n", i, a)
	}
}
