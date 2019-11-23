package main

import (
	"fmt"
	"github.com/iGoogle-ink/jorm"
)

func main() {
	err := jorm.InitMySQL("root:123456@tcp(127.0.0.1:3306)/wallet?charset=utf8&parseTime=true&loc=Local") //
	if err != nil {
		fmt.Println("err:", err)
	}
	result, err := jorm.CallProcedure("mypro", 2, 1).InParams(1,1).Query()
	if err != nil {
		fmt.Println("err:", err)
	}
	for _, v := range result {
		fmt.Println(v)
	}
}
