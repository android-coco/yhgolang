package main

import (
	"fmt"
	"math"
)

func main() {
	/* 定义局部变量 */
	var a int = 100
	var b int = 200
	var ret int

	/* 调用函数并返回最大值 */
	ret = max(a, b)

	fmt.Printf("最大值是 : %d\n", ret)

	a1, b2 := swap("Mahesh", "Kumar")
	fmt.Println(a1, b2)
	fmt.Println("======声明函数变量========")
	/* 声明函数变量 */
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}
	/* 使用函数 */
	fmt.Println(getSquareRoot(9))
	fmt.Println("======闭包========")
	/* nextNumber 为一个函数，函数 i 为 0 */
	nextNumber := getSequence()
	/* 调用 nextNumber 函数，i 变量自增 1 并返回 */
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	/* 创建新的函数 nextNumber1，并查看结果 */
	nextNumber1 := getSequence()
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())
	d := A
	d(a, b)
	//匿名函数
	e := func() {
		fmt.Println("匿名函数")
	}
	e()

	f := closure(10)
	fmt.Println(f(1))
	fmt.Println(f(2))
	//a,c,b 执行顺序相反
	fmt.Println("a")
	defer fmt.Println("b")
	defer fmt.Println("c")

	for i := 0; i < 3; i++ {
		//2,1,0
		defer fmt.Println(i)
	}
	for i := 0; i < 3; i++ {
		defer func() {
			//3,3,3
			fmt.Println(i)
		}()
	}
	J()
	K()
	L()
}
func J() {
	fmt.Println("Func J")
}
func K() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("RECOVER IN K")
		}
	}()
	panic("Panic in K")
}
func L() {
	fmt.Println("Func C")
}

/* 函数作为返回值 */
func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func max(num1 int, num2 int) int {
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

//返回多个值(交换)
func swap(x string, y string) (string, string) {
	return y, x
}

func A(a ...int) {
	fmt.Println(a)
}
func B() {

}

//返回值为函数  闭包
func closure(x int) func(int) int {
	fmt.Printf("%p\n", &x)
	return func(y int) int {
		fmt.Printf("%p\n", &x)
		return x + y
	}
}
