package main

import (
	"fmt"
	"sync"
	"time"
)

func updateSlice(s []int) {
	s[0] = 100
}

var mu sync.Mutex
var chain string

func A() {
	mu.Lock()
	defer mu.Unlock()
	chain = chain + " -->A"
	B()
}
func B() {
	mu.Lock()
	defer mu.Unlock()
	chain = chain + " -->B"
	C()
}

func C() {
	mu.Lock()
	defer mu.Unlock()
	chain = chain + " -->C"
}

//求交集
func intersect(slice1, slice2 []int) []int {
	m1 := make(map[int]int)
	m2 := make([]int, 0)
	tmp := make(map[int]int)
	for _, v := range slice1 {
		m1[v]++
	}

	for _, v := range slice2 {
		if _, ok := tmp[v]; ok {
			continue
		}
		times, _ := m1[v]
		tmp[v] = 0
		if times == 1 {
			m2 = append(m2, v)
		}
	}
	return m2
}

type query func(string) string

func exec(name string,vs ...query)string  {
	ch := make(chan string)
	fn := func(i int) {
		ch <- vs[i](name)
	}
	for i,_:=range vs {
		go fn(i)
	}
	return  <- ch
}

func proc()  {
	panic("ok")
}
func main() {

	go func() {
		for {
			t := time.Tick(1 * time.Second)
			select {
			case <-t:
				defer proc()
				if r := recover(); r != nil {
					fmt.Println("recover...:", r)
				}
				fmt.Println("11")
			}
		}
	}()
	select {

	}
	ret:= exec("111",func(n string)string{
		return n + "func1"
	},func(n string)string{
		return n + "func2"
	},func(n string)string{
		return n + "func3"
	},func(n string)string{
		return n + "func4"
	})
	fmt.Print(ret)

	mums1 := []int{4, 9, 5}
	mums2 := []int{9, 4, 8, 4}
	fmt.Print(intersect(mums1, mums2))
	return
	chain = "main"
	A()
	fmt.Print(chain)
	return
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	fmt.Println("arr[2:6] =", arr[2:6])
	fmt.Println("arr[:6] =", arr[:6])
	s1 := arr[2:]
	fmt.Println("s1 =", s1)
	s2 := arr[:]
	fmt.Println("s2 =", s2)

	fmt.Println("After updateSlice(s1)")
	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(arr)

	fmt.Println("After updateSlice(s2)")
	updateSlice(s2)
	fmt.Println(s2)
	fmt.Println(arr)

	fmt.Println("Reslice")
	fmt.Println(s2)
	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[2:]
	fmt.Println(s2)

	fmt.Println("Extending slice")
	arr[0], arr[2] = 0, 2
	fmt.Println("arr =", arr)
	s1 = arr[2:6]
	s2 = s1[3:5] // [s1[3], s1[4]]
	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n",
		s1, len(s1), cap(s1))
	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n",
		s2, len(s2), cap(s2))

	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	fmt.Println("s3, s4, s5 =", s3, s4, s5)
	// s4 and s5 no longer view arr.
	fmt.Println("arr =", arr)

	// Uncomment to run sliceOps demo.
	// If we see undefined: sliceOps
	// please try go run slices.go sliceops.go
	fmt.Println("Uncomment to see sliceOps demo")
	// sliceOps()
}
