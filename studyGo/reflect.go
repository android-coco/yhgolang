package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) Hello() {
	fmt.Println("Hello World")
}

func main() {
	u := User{1, "yh", 28}
	Info(u)
}

func Info(o interface{}) {
	//获取该接口是什么类型
	t := reflect.TypeOf(o)
	fmt.Println("Type:", t.Name())

	//传进来的值为值拷贝,如果传指针类型就报错
	if k := t.Kind(); k != reflect.Struct {
		fmt.Println("错误")
	}

	v := reflect.ValueOf(o) //获取所有的属性
	fmt.Println("Fields:")
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)               //属性
		val := v.Field(i).Interface() //值
		fmt.Printf("%s:%v = %v\n", f.Name, f.Type, val)
	}
	fmt.Println("Method:")
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%s:%v\n", m.Name, m.Type)
	}
}
