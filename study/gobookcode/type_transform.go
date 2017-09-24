/*
* @Author: yhlyl
* @Date:   2017-09-24 12:59:53
* @Last Modified by:   yhlyl
* @Last Modified time: 2017-09-24 21:37:36
 */
package main

import (
	"fmt"
	"reflect"
)

type YH int
type HanderI func(a, b int)

func main() {
	var x = 7
	fmt.Println(YH(x))
	// value, ok := x.(YH)
	// if !ok {
	// 	fmt.Println("It's not ok for type YH")
	// 	return
	// }
	//fmt.Println("The value is ", value)
	var c = a
	fmt.Println(reflect.TypeOf(c))
	fmt.Println(c)
	transform(a, 10, 10)
	//常用
	v1 := reflect.ValueOf(c)
	y := v1.Interface().(HanderI) // y 有类型float64
	fmt.Println(y)
}

func a(a, b int) {
	fmt.Println(a + b)
}

func transform(x interface{}, a1, b1 int) {
	switch v := x.(type) { //v表示b1 接口转换成Bag对象的值
	case func(a, b int):
		x.(func(a, b int))(a1, b1)
	default:
		fmt.Println("错误", v)
	}

}
