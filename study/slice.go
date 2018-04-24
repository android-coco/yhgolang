package main

//切片是对数组的抽象。
/* 内置类型切片("动态数组"),与数组相比
切片的长度是不固定的，可以追加元素，
在追加时可能使切片的容量增大。*/

import (
	"fmt"
)

func main() {
	var numbers = make([]int, 3, 5)

	printSlice(numbers)

	fmt.Printf("=====空(nil)切片=====\n")

	var numbers1 []int

	printSlice(numbers1)

	str1 := []string{}

	fmt.Println(len(str1), str1)

	if numbers1 == nil {
		fmt.Printf("切片是空的\n")
	}

	fmt.Printf("=====切片截取=====\n")

	/* 创建切片 */
	numbers2 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}

	printSlice(numbers2)

	fmt.Printf("=====打印原始切片=====\n")
	/* 打印原始切片 */
	fmt.Println("numbers2 ==", numbers2)

	fmt.Printf("=====打印子切片从索引1(包含) 到索引4(不包含)=====\n")
	/* 打印子切片从索引1(包含) 到索引4(不包含)*/
	fmt.Println("numbers2[1:4] ==", numbers2[1:4])

	fmt.Printf("=====打印子numbers2[:3]切片从索引0(包含) 到索引3(不包含)=====\n")
	/* 默认下限为 0*/
	fmt.Println("numbers2[:3] ==", numbers2[:3])

	fmt.Printf("=====打印子numbers2[4:]切片从索引4(包含) 到索引最后(不包含)=====\n")
	/* 默认上限为 len(s)*/
	fmt.Println("numbers2[4:] ==", numbers2[4:])

	numbers3 := make([]int, 0, 5)
	printSlice(numbers3)

	/* 打印子切片从索引  0(包含) 到索引 2(不包含) */
	numbers4 := numbers2[:2]
	printSlice(numbers4)

	/* 打印子切片从索引 2(包含) 到索引 5(不包含) */
	numbers5 := numbers2[2:5]
	printSlice(numbers5)

	fmt.Printf("=====append() 和 copy() 函数=====\n")
	var numbers6 []int
	printSlice(numbers6)

	/* 允许追加空切片 */
	fmt.Printf("=====append()空切片=========\n")
	numbers6 = append(numbers6, 0)
	printSlice(numbers6)

	/* 向切片添加一个元素 */
	numbers6 = append(numbers6, 1)
	printSlice(numbers6)

	/* 向切片添加多个元素 */
	numbers6 = append(numbers6, 2, 3, 4)
	printSlice(numbers6)

	/* 创建切片 numbers7 是之前切片的两倍容量*/
	numbers7 := make([]int, len(numbers6), (cap(numbers6))*2)

	/* 拷贝 numbers6 的内容到 numbers7 */
	copy(numbers7, numbers6)
	printSlice(numbers7)


	fmt.Printf("=====XXXX=========\n")
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)//6 6

	// Slice the slice to give it zero length.
	s = s[1:4]//0 6
	printSlice(s)

	// Extend its length.
	s = s[:5] //4 6
	printSlice(s)

	// Drop its first two values.
	s = s[3:]//2 4
	printSlice(s)

}

//切片是可索引的，并且可以由 len() 方法获取长度。
//切片提供了计算容量的方法 cap() 可以测量切片最长可以达到多少。
func printSlice(x []int) {
	fmt.Printf("len = %d cap = %d slice = %v \n", len(x), cap(x), x)
}
