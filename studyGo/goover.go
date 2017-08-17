package main

import (
	std "fmt"
	"math"
	"strconv"
)

type (
	byte uint8
	rune int32
	文本   string
)

const (
	a1 int = 1
	b1     = 'A'
	c1
	d1
)

const (
	k, u = 1, "2"
	k1, u1
)

//常量只能引用常量和内置函数
const x, y, z = 1, "2", true
const (
	z1 = 'A'
	z2
	z3 = iota
	z4
)

//计算机存储单位
const (
	_          = iota
	KB float64 = 1 << (iota * 10)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

//chan管道
//channel并发<-

/*
 6: 0110
11: 1011
----------
&   0010 = 2
|   1111 = 15
^   1101 = 13
&^  0100 = 4

&^第二个位数上为1时强制把第一个数该位上的数改为0
*/

func main() {
	var b 文本
	b = "中文类型名"
	std.Println(math.MaxInt8)
	std.Println(b)

	var a float32 = 100.1
	std.Println(a)
	c := int(a)
	std.Println(c)

	var d int = 65
	e := string(d)
	f := strconv.Itoa(d)
	g, _ := strconv.Atoi(f)
	std.Println(e) //A  字符
	std.Println(f) //65 字符串
	std.Println(g) //65 int
	std.Println(a1, b1, x, y, z, c1, d1)
	std.Println(k, u, k1, u1)
	std.Println("==================")
	std.Println(z1, z2, z3, z4)
	std.Println("==================")
	std.Println(6 & 11)
	std.Println("==================")
	std.Println(6 | 11)
	std.Println("==================")
	std.Println(6 ^ 11)
	std.Println("==================")
	std.Println(6 &^ 11)
	std.Println("=====结合常量的iota与<<运算符实现计算机的存储单位===========")
	std.Println(KB, MB, GB, TB, PB, EB, ZB, YB)
}
