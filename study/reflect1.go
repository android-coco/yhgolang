package main

import (
	"fmt"
	"reflect"
)

type X struct {
	Name string
	Age  int
}

func main() {
	var x1 = []X{}
	Info(x1)
}

func Info(o interface{}) {
	t := reflect.TypeOf(o)                   //反射使用 TypeOf 和 ValueOf 函数从接口中获取目标对象信息
	fmt.Println("Type:", t.Name(), t.Kind()) //调用t.Name方法来获取这个类型的名称
	if k := t.Kind(); k != reflect.Slice {   //通过kind方法判断传入的类型是否是我们需要反射的类型
		fmt.Println("xx")
		return
	}
	v := reflect.ValueOf(o) //打印出所包含的字段
	z := v.Interface().([]X)
	fmt.Println("Fields:", z, &z)

}
