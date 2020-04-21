package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var randomColor []int32
var mu sync.Mutex // guards balance


func Deposit() {
	mu.Lock()
	randomColor = RandomColor()
	mu.Unlock()
}

func Balance() []int32 {
	mu.Lock()
	b := randomColor
	mu.Unlock()
	return b
}
func main() {

	for {
		go func() {
			Deposit()
			fmt.Println("协程1:", randomColor)
		}()
		go func() {
			fmt.Println("协程2:", Balance())
		}()
		go func() {
			fmt.Println("协程3:", Balance())
		}()
		time.Sleep(2 * time.Second)
	}

}

//红色、绿色、黄色
// 0 16 32
var forestColor = []int32{
	0x00, 0x10, 0x20,
	0x00, 0x10, 0x20,
	0x00, 0x10, 0x20,
	0x00, 0x10, 0x20,
	0x00, 0x10, 0x20,
	0x00, 0x10, 0x20,
	0x00, 0x10, 0x20,
	0x00, 0x10, 0x20,
}

// 随机颜色
func RandomColor() []int32 {

	var fColors = make([]int32, 0)

	for i := len(forestColor) - 1; i > 0; i-- {
		num := RandInterval(0, int32(i))
		forestColor[i], forestColor[num] = forestColor[num], forestColor[i]
	}

	fColors = append(fColors, forestColor...)

	return fColors
}

func RandInterval(b1, b2 int32) int32 {
	if b1 == b2 {
		return b1
	}

	min, max := int64(b1), int64(b2)
	if min > max {
		min, max = max, min
	}
	return int32(rand.Int63n(max-min+1) + min)
}
