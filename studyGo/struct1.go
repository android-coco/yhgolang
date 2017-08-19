package main

import (
	"fmt"
)

//组合 不是 继承
type humam struct {
	Sex int
}

type teacher struct {
	humam //组合
	Name  string
	Age   int
}

type student struct {
	humam //组合
	Name  string
	Age   int
}

type test struct{}
type person struct {
	Name string
	Age  int
}

//匿名结构
type person1 struct {
	Name    string
	Age     int
	Contact struct {
		Phone, City string
	}
}

//匿名字段
type person3 struct {
	string
	int
}

func main() {
	a := test{}
	fmt.Println(a)
	//字段初始化
	b := person{
		Name: "joe",
		Age:  19,
	}
	//结构初始化一般用指针
	fmt.Println(b)
	A(b)
	fmt.Println(b)

	fmt.Println(b)
	B(&b)
	fmt.Println(b)
	//匿名结构
	c := &struct {
		Name string
		Age  int
	}{
		Name: "yh",
		Age:  19,
	}
	fmt.Println(c)
	//匿名结构初始化
	d := person1{Name: "yh", Age: 28}
	d.Contact.City = "深圳"
	d.Contact.Phone = "18822855252"
	fmt.Println(d)

	//匿名字段初始化
	e := person3{"joe", 19}
	fmt.Println(e)

	//结构之间赋值
	f := person{Name: "ok", Age: 23}
	var j person
	j = f
	fmt.Println(j)
	fmt.Println(j == f)
	//组合演示
	h := teacher{Name: "teacher", Age: 40, humam: humam{Sex: 0}}
	i := student{Name: "student", Age: 20, humam: humam{Sex: 1}}
	h.Name = "teacher2"
	h.Age = 41
	//h.humam.Sex = 100
	h.Sex = 100
	fmt.Println(h, i)
}

func A(per person) {
	per.Age = 13
	fmt.Println("A", per)
}

func B(per *person) {
	per.Age = 13
	fmt.Println("B", per)
}
