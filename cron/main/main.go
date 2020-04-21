package main

import (
	"fmt"
	"sync"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var c = make(chan int, 10)
var x sync.Map
func main() {
	x.Store("1",1)
	x.Store("2",1)
	x.Store("33",1)
	x.Store("4",1)
	x.Range(func(key, value interface{}) bool {
		fmt.Println(key,value)
		return false
	})
	return
	go func() {
		for i := 0; i < 10; i++ {
			c <- 1
		}

		//for{
		//	time.Sleep(1 * time.Second)
		//	c <- 1
		//	fmt.Println("1111",len(c))
		//}

	}()
	go func() {
		for {
			//for v := range c {
			//	fmt.Println("====", len(c), v)
			//}
			select {
			case v := <-c:
				fmt.Println("====", len(c), v)
			}
			fmt.Println("===1111",10 % 4)
		}

	}()
	time.Sleep(100 * time.Hour)
	//fmt.Println(<- c)
	//fmt.Println(len(c))
	//var x uintptr
	//fmt.Println(x)
	//fmt.Sprintf("1%s2", "fdsafasd")
	//https://xjwl-qp.oss-accelerate.aliyuncs.com/img/head_1.jpg
	//var wireteString = ""
	//for i := 1; i <= 2180; i++ {
	//	wireteString += strconv.Itoa(i) + "\n"
	//}
	//
	//var d1 = []byte(wireteString)
	//err2 := ioutil.WriteFile("./output2.txt", d1, 0666) //写入文件(字节数组)
	//check(err2)
}
