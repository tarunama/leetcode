// https://leetcode.com/problems/shuffle-the-array/
package main

func shuffle(nums []int, n int) []int {
	result := []int{}
	for i := 0; i < n; i++ {
		result = append(result, nums[i], nums[i+n])
	}
	return result
}
