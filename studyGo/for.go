package main

import "fmt"

func main() {
	var b int = 5
	var a int
	numbers := [6]int{1, 2, 3, 5}

	for a := 0; a < 10; a++ {
		fmt.Printf("a 的值为：%d\n", a)
	}

	for a < b {
		//a++;
		fmt.Printf("a 的值为：%d\n", a)
		a++
	}

	for i, x := range numbers {
		fmt.Printf("第 %d 位 x1 的值 = %d\n", i, x)
	}
	println("================================输出 2 到 100 间的素数=====================")
	var C, c int  //声明变量
	C = 1         /*这里不写入FOR循环是因为For语句执行之初会将C的值变为1，当我们goto A时for语句会重新执行（不是重新一轮循环）*/
	for C < 100 { //打一个标记
		C++ //C=1不能写入for这里就不能写入
		for c = 2; c < C; c++ {
			if C%c == 0 {
				goto A //若发现因子则不是素数
			}
		}
		fmt.Println(C, "是素数")
	}
A:
	println("=====================================================")
	/* 定义局部变量 */
	var i, j int
	for i = 2; i < 100; i++ {
		for j = 2; j <= (i / j); j++ {
			if i%j == 0 {
				break // 如果发现因子，则不是素数
			}
		}
		if j > (i / j) {
			fmt.Printf("%d  是素数\n", i)
		}
	}
B:
	for i := 0; i < 3; i++ {
		for {
			break B
		}
	}
C:
	for i := 0; i < 3; i++ {
		for {
			continue C
			fmt.Println("dasfadfasd") //这句话不会被执行
		}
	}
	fmt.Println("Over")
	//冒泡排序
	myArray := [5]int{10, 11, 4, 9, 2}
	fmt.Println(myArray) //排序前
	fmt.Println(myArray[1:5])
	lenNumber := len(myArray)
	for i := 0; i < lenNumber; i++ {
		for j := i + 1; j < lenNumber; j++ {
			if myArray[i] < myArray[j] {
				temp := myArray[j]
				myArray[j] = myArray[i]
				myArray[i] = temp
			}
		}
	}
	fmt.Println(myArray) //排序后
	p := &myArray
	fmt.Println(p)
	fmt.Println(p[0:])
	yh1, yh2 := 1, 2
	yhArray := [3]*int{&yh1, &yh2}
	//var yhArray1 []*int
	var yhArray1 = []*int{&yh1, &yh2}
	fmt.Println(yhArray)
	fmt.Println(yhArray1)
}
