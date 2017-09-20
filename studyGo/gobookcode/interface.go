package main

import (
	"fmt"
	"reflect"
)

type A interface {
	Add() int
}

type B struct {
	ID int
}

func (b B) Add() int {
	return 1 + 1
}

type C struct {
	B
	Name string
}

func Transfer(slice []A) []A {
	var ifSlice []A = make([]A, len(slice))
	for idx, v := range slice {
		ifSlice[idx] = v
	}
	return ifSlice
}

func main() {
	a := []A{C{B: B{ID: 1}, Name: "youhao"}, C{B: B{ID: 2}, Name: "youhao"}, C{B: B{ID: 3}, Name: "youhao"}}
	c := []C{C{B: B{ID: 4}, Name: "youhao"}, C{B: B{ID: 5}, Name: "youhao"}, C{B: B{ID: 6}, Name: "youhao"}}

	fmt.Println(reflect.TypeOf(a[0]))
	fmt.Println(a)
	fmt.Println(c)

	fmt.Println([]A(c))

	fmt.Println(reflect.TypeOf(Transfer(a)[0]))
}
