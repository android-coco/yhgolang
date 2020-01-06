package main

import (
	"fmt"
	"time"
)

func main() {
	//host, port, err := net.SplitHostPort("127.0.0.1:8880")
	//fmt.Println(host, port, err)
	//err := jorm.InitMySQL("root:123456@tcp(127.0.0.1:3306)/wallet?charset=utf8&parseTime=true&loc=Local") //
	//if err != nil {
	//	fmt.Println("err:", err)
	//}
	//result, err := jorm.CallProcedure("mypro", 2, 1).InParams(1,1).Query()
	//if err != nil {
	//	fmt.Println("err:", err)
	//}
	//for _, v := range result {
	//	fmt.Println(v)
	//}

	t := time.NewTimer(5 * time.Second)
	go func() {
		select {
		case <-t.C:
			fmt.Println("===")
		}
	}()

	t.Reset(0)
	time.Sleep(4 * time.Second)
}
