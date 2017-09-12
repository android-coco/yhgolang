package main

import (
	"time"
	"fmt"
)

func main() {
	t := time.Now()
	fmt.Println(t)
	fmt.Println(t.Format(time.ANSIC))

	s := []string{"a","b","c"}
	//闭包是引用传递
	for _,v := range s{
		go func(v string) {
			fmt.Println(v)//输出c,c,c  如果要输出a,b,c需要用参数传递
		}(v)
	}
	select {
	}
}
