// https://leetcode.com/problems/running-sum-of-1d-array/
package main

import (
	"fmt"
)

func runningSum(nums []int) []int {
	result := []int{}
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		result = append(result, sum)
	}
	return result
}

func main() {
	input := []int{3, 1, 2, 10, 1}
	fmt.Println([]int{3, 4, 6, 16, 17}, runningSum(input))
}
