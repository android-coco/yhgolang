package main

import (
	out "fmt"
)

//切片扩容后内存地址改变
func main() {
	a := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'j'}
	a1 := a[1:4]
	a2 := a1[1:3]
	out.Println(string(a1), len(a1), cap(a1))
	out.Println(string(a2), len(a2), cap(a2))

	b := make([]int, 3, 6)
	out.Printf("%v%p\n", b, b)
	b = append(b, 1, 2, 3)
	out.Printf("%v%p\n", b, b)

	d := []int{1, 2, 3, 4, 5}
	d1 := d[:4] //[1 2 3 4]
	d2 := d[3:] //[4 5]
	out.Printf("%v%v\n", d1, d2)
	d1[3] = 9
	//[1 2 3 9] [9 5]
	d2[0] = 10
	//[1 2 3 10] [10 5]
	out.Printf("%v%v\n", d1, d2)

	e1 := []int{1, 2, 3, 4}
	e2 := []int{5, 6}
	copy(e1, e2)
	out.Printf("%v\n", e1) //[5 6 3 4]

	e1 = []int{1, 2, 3, 4}
	e2 = []int{5, 6}
	copy(e2, e1)
	//copy(e2[1:2], e1[3:4])
	out.Printf("%v\n", e2) //[1 2]
}
