package main

import (
	"fmt"
	"time"
)

var time1 time.Time

func main() {

	//返回现在时间
	tNow := time.Now()
	//时间转化为string，layout必须为 "2006-01-02 15:04:05"
	timeNow := tNow.Format("2006-01-02 15:04:05")
	fmt.Println("tNow(time format):", tNow, time.ANSIC)
	fmt.Println("tNow(string format):", timeNow)

	//string转化为时间，layout必须为 "2006-01-02 15:04:05"
	t, _ := time.Parse("2006-01-02 15:04:05", "2014-06-15 08:37:18")
	time1 = t
	fmt.Println("t(time format)", time1)

	//某个时间点 前后判断
	trueOrFalse := t.After(tNow)
	if trueOrFalse == true {
		fmt.Println("t（2014-06-15 08:37:18）在tNow之后!")
	} else {
		fmt.Println("t（2014-06-15 08:37:18）在tNow之前!")
	}
	fmt.Println()

	// s := []string{"a", "b", "c"}
	// //闭包是引用传递
	// for _, v := range s {
	// 	go func(v string) {
	// 		fmt.Println(v) //输出c,c,c  如果要输出a,b,c需要用参数传递
	// 	}(v)
	// }
	// select {}
}
