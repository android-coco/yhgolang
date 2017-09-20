package main

import "fmt"

func Add1(a, b int) (c int) {
	c = a + b
	return
}

func MyPrintf(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is an int value.")
		case string:
			fmt.Println(arg, "is an string value.")
		case int64:
			fmt.Println(arg, "is an int64 value.")
		case struct{}:
			fmt.Println(arg, "is an struct{} value.")
		case func():
			fmt.Println(arg, "is an func() value.")
		case chan int:
			fmt.Println(arg, "is an chan int value.")
		case map[string]string:
			fmt.Println(arg, "is an map[string]string  value.")
		case []int:
			fmt.Println(arg, "is an []int  value.")
		default:
			fmt.Println(arg, "is an unknown type.")
		}
	}
}

func main() {
	var v1 int = 1
	var v2 int64 = 234
	var v3 string = "hello"
	var v4 float32 = 1.234
	var v5 struct{}
	var v6 func()
	var v7 []int
	var v8 map[string]string
	var v9 chan int
	var v10 float64 = 1.2
	MyPrintf(v1, v2, v3, v4, v5, v6, v7, v8, v9, v10)
	/*
	1     is an int value.
	234   is an int64 value.
	hello is an string value.
	1.234 is an unknown type.
	{}    is an struct{} value.
	<nil> is an func() value.
	[]    is an []int  value.
	map[] is an map[string]string  value.
	<nil> is an chan int value.
	1.2 is an unknown type.
	*/
	fmt.Println(Add1(1,2))
}
