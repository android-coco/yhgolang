package main

import "fmt"

var lastOccurred = make([]int, 0xffff) //65535
func lengthOfNonRepeatingSubStr(s string) int {
	//每个元素出现的位置key->元素  value->位置
	//lastOccurred := make(map[rune]int)

	for i := range lastOccurred {
		lastOccurred[i] = -1
	}

	// 开始位置
	start := 0
	// 最长多少
	maxLength := 0

	for i, ch := range []rune(s) {
		//if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
		//	start = lastI + 1
		//}
		if lastI := lastOccurred[ch]; lastI != -1 && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}

	return maxLength
}
func main() {
	fmt.Println(
		lengthOfNonRepeatingSubStr("abcabcbb"))
	fmt.Println(
		lengthOfNonRepeatingSubStr("bbbbb"))
	fmt.Println(
		lengthOfNonRepeatingSubStr("pwwkew"))
	fmt.Println(
		lengthOfNonRepeatingSubStr(""))
	fmt.Println(
		lengthOfNonRepeatingSubStr("b"))
	fmt.Println(
		lengthOfNonRepeatingSubStr("abcdef"))
	fmt.Println(
		lengthOfNonRepeatingSubStr("这里是慕课网"))
	fmt.Println(
		lengthOfNonRepeatingSubStr("一二三二一"))
	fmt.Println(
		lengthOfNonRepeatingSubStr(
			"黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"))
}
