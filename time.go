package main

import (
	"fmt"
	"time"
)

func main() {
	time1()
	time2()
}

func time1() {
	time1 := "2015-03-20 08:50:29"
	time2 := "2015-03-21 08:50:29"
	//先把时间字符串格式化成相同的时间类型
	t1, err := time.Parse("2006-01-02 15:04:05", time1)
	t2, err := time.Parse("2006-01-02 15:04:05", time2)
	if err == nil && t1.Before(t2) {
		//处理逻辑
		fmt.Println("true")
	}
}

func time2() {
	time1 := "2015-03-20 08:50:29"
	//先把时间字符串格式化成相同的时间类型
	t1, err := time.Parse("2006-01-02 15:04:05", time1)
	t2 := time.Now()
	if err == nil && t1.Before(t2) {
		//处理逻辑
		fmt.Println("true")
	}
}
