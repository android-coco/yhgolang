package main

import (
	"context"
	"fmt"
	"strings"
	"time"
)

const (
	x = iota
	_
	y
	z = "zz"
	k
	j
	p = iota
)

type person struct {
	name string
}

func x11() {
	ch := make(chan string, 6)
	go func() {
		for {
			v, ok := <-ch
			if !ok {
				fmt.Println("结束")
				return
			}
			fmt.Println(v)
		}
	}()

	ch <- "煎鱼还没进锅里..."
	ch <- "煎鱼进脑子里了！"
	close(ch)
	time.Sleep(time.Second)
}

func x22() {
	ch := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				ch <- struct{}{}
				return
			default:
				fmt.Println("煎鱼还没到锅里...")
			}

			time.Sleep(500 * time.Millisecond)
		}
	}(ctx)

	go func() {
		time.Sleep(3 * time.Second)
		cancel()
	}()

	<-ch
	fmt.Println("结束")
}

func xx(a []int) {
	a[0] = 22
}

func clin() {
	s1 := []int{1, 2, 3}
	s2 := s1[1:]
	s2[1] = 4
	fmt.Println(s1)
	s2 = append(s2, 5, 6, 7)
	fmt.Println(s1)
}

type Math struct {
	x, y int
}

var m = map[string]Math{
	"foo": Math{2, 3},
}

func f(n int) (r int) {
	defer func() {
		r += n
		recover()
	}()
	var f func()
	defer f()
	f = func() {
		r += 2
	}
	return n + 1
}
func main() {

	fmt.Println(f(3))
	return
	clin()
	j := []int{1, 2, 3}
	fmt.Println(j)
	xx(j)
	fmt.Println(j)
	return
	x11()
	x22()
	return
	xx := strings.Join([]string{"2", "3"}, "1")
	fmt.Println(xx)

	fmt.Println(x, y, z, k, j, p)

	var m map[person]int
	p := person{"mike"}
	fmt.Println(m[p])

	i := []int{5, 6, 7}
	hello(i)
	fmt.Println(i[0])
	edisKey := fmt.Sprintf("khl:smsloginerr:%v", "req.UserPhone")
	fmt.Println(edisKey)

	i1 := 5
	defer hello1(i1)
	i1 = i1 + 10

	c := Work{3}
	var a A = &c
	var b B = &c
	fmt.Println(a.ShowA())
	fmt.Println(b.ShowB())

	s := [3]int{1, 2, 3}
	//a := s[:0]
	//b := s[:2]
	fmt.Println(cap(s))
	c1 := s[1:2:cap(s)]
	fmt.Println(cap(c1))
}

func hello(num []int) {
	num[0] = 18
	//i := 65
	fmt.Println(string(65))
}

func hello1(i int) {
	fmt.Println(i)
}

func incr(p *int) int {
	*p++
	return *p
}

type A interface {
	ShowA() int
}
type B interface {
	ShowB() int
}
type Work struct {
	i int
}

func (w *Work) ShowA() int {
	return w.i + 10
}
func (w *Work) ShowB() int {
	return w.i + 20
}
