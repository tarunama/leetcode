// https://leetcode.com/problems/running-sum-of-1d-array/
package main

func runningSum(nums []int) []int {
	result := []int{}
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		result = append(result, sum)
	}
	return result
}
