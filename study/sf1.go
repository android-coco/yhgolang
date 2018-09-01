/*
* @Author: yhlyl
* @Date:   2018-09-01 20:39:18
* @Last Modified by:   yhlyl
* @Last Modified time: 2018-09-01 21:04:15
 */
package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	j := 0
	lens := len(nums)
	for i := 0; i < lens; i++ {
		if nums[i] != nums[j] {
			j++
			nums[j] = nums[i]
		}
	}
	j++
	return j
}

func main() {
	z := []int{1, 1, 1, 2}
	x := removeDuplicates(z)
	fmt.Println(x)
	for i := 0; i < x; i++ {
		fmt.Println(z[i])
	}
}
