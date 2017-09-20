package main

//方法没有重载
import (
	"fmt"
)

//为int 添加方法
type TZ int

type A struct {
	Name string
}

type B struct {
	Name string
}

func main() {
	a := A{}
	a.Print()
	fmt.Println(a.Name)

	b := B{}
	b.Print()
	fmt.Println(b.Name)

	var c TZ
	c.Print()
	(*TZ).Print(&c)
}

func (a *A) Print() {
	a.Name = "AA"
	fmt.Println("A")
}

//编译错误
// func (a A) Print(b int) int {
// 	fmt.Println("A")
// }

func (b B) Print() {
	b.Name = "BB"
	fmt.Println("B")
}
func (tz *TZ) Print() {
	fmt.Println("TZ")
}
