package main

import (
	"fmt"
	"regexp"
)

func main() {
	a := [3]int{1, 2, 3}
	fmt.Printf("a %T\n", a)
	fmt.Printf("a %#v\n", a)
	fmt.Printf("a %#v\n", []byte("1m,22"))

	var b []int
	b = append(b, 1)
	fmt.Printf("b %#v\n", b)
	b = append(b, 1, 2, 3)
	fmt.Printf("b %#v\n", b)
	b = append(b, 1)
	fmt.Printf("b %#v\n", b)
	b = append(b, []int{3, 4, 5}...)
	fmt.Printf("b %#v\n", b)

	var c = []int{1, 2, 3}
	c = append([]int{0}, c...)
	fmt.Printf("c %#v\n", c)

	var d = []int{1, 2, 3}
	//在第i个位置插入x
	i := 2
	x := 4
	d = append(d[:i], append([]int{x}, d[i:]...)...)
	fmt.Printf("d %#v\n", d)

	var e []int
	e = append(e, 0)
	f := copy(e[1:], e[:])
	fmt.Printf("f %#v\n", f)
	e[0] = 3
	fmt.Printf("e %#v\n", e)

	e = []int{7, 8, 9, 10}
	j := copy(d, e)
	fmt.Printf("j %#v\n", j)
	fmt.Printf("d %#v\ne %#v\n", d, e)

	bytes := []byte("WO SHI NI MMM")
	h := TirmSpace(bytes)
	fmt.Printf("h %#v\n", string(h))

	i1 := Filter([]byte("WO,SHI,NI,MMMM"), func(x byte) bool {
		if x == ',' {
			return true
		} else {
			return false
		}
	})
	fmt.Printf("ii %#v\n", string(i1))

	regexp.MustCompile("")
}

func TirmSpace(s []byte) []byte {
	b := s[:0]
	for _, x := range s {
		if x != ' ' {
			b = append(b, x)
		}
	}
	return b
}

type StirngFilter func(x byte) bool

func Filter(s []byte, fn StirngFilter) []byte {
	b := s[:0]
	for _, x := range s {
		if !fn(x) {
			b = append(b, x)
		}
	}
	return b
}
