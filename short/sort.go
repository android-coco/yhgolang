package main

import (
	"sort"
	"fmt"
)

func main() {
	a := []int{3, 6, 1, 2, 9, 10, 9, 8}
	sort.Ints(a)
	for _, v := range a {
		fmt.Println(v)
	}
}
