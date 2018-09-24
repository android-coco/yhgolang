package main

import (
	"fmt"
	"math/cmplx"
	"math"
)

var aa = 3
var ss = "kkk"
var (
	xz = 1
	xx = 2
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a, b, s = 3, 4, "abc"
	fmt.Println(a, b, s)
}

func varialbeShorter() {
	a, b, c, s := 3, 4, "abc", true
	fmt.Println(a, b, c, s)
}

func euler() {
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))

	fmt.Printf("%3f\n", cmplx.Exp(1i*math.Pi)+1)
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)
}

func triangle() {
	var a, b = 3, 4
	fmt.Println(calcTriangle(a,b))
}
func calcTriangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}

const filename = "abc.txt"

func consts() {
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

//枚举
func enums() {
	const (
		php        = 0
		java       = 1
		python     = 2
		golang     = 3
		javascript = 4
	)
	const (
		php1        = iota
		_
		python1
		golang1
		javascript1
	)

	//b,kb,mb,gb,tb,pb
	const (
		b  = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(php, java, python, golang, javascript)
	fmt.Println(php1, python1, golang1, javascript1)
	fmt.Println(b, kb, mb, gb, tb, pb)

}

func main() {
	fmt.Println("Hello world")
	variableZeroValue()
	variableInitialValue()
	varialbeShorter()
	fmt.Println(aa, ss, xz, xx)

	euler()
	triangle()

	consts()

	enums()
}
