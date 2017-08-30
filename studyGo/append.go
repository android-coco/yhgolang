package main

import (
	_ "awesomeProject/org.yh/util"
	"fmt"
)

func main() {
	/**
			append主要用于给某个切片（slice）追加元素
			如果该切片存储空间（cap）足够，就直接追加，长度（len）变长；
			如果空间不足，就会重新开辟内存，并将之前的元素和新的元素一同拷贝进去
				第一个参数为切片，后面是该切片存储元素类型的可变参数
	    基础用法：
	*/
	slice := append([]int{1, 2, 3}, 4, 5, 6)
	fmt.Println(slice) //[1 2 3 4 5 6]
	//第二个参数也可以直接写另一个切片，将它里面所有元素拷贝追加到第一个切片后面。
	//要注意的是，这种用法函数的参数只能接收两个slice，并且末尾要加三个点
	slice1 := append([]int{1, 2, 3}, []int{4, 5, 6}...)
	fmt.Println(slice1) //[1 2 3 4 5 6]
	//还有种特殊用法，将字符串当作[]byte类型作为第二个参数传入
	//append函数返回值必须有变量接收，不然编译器会报错
	bytes := append([]byte("hello"), " world"...)
	fmt.Println(string(bytes))
}
