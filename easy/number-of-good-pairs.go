// https://leetcode.com/problems/number-of-good-pairs/

func numIdenticalPairs(nums []int) int {
	result := 0
	for i, ni := range nums {
		for j, nj := range nums {
			if i < j && ni == nj {
				result += 1
			}
		}
	}
	return result
}
