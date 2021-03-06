// https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/submissions/

func kidsWithCandies(candies []int, extraCandies int) []bool {
	result := []bool{}
	maxC := 0
	for i := 0; i < len(candies); i++ {
		if candies[i] > maxC {
			maxC = candies[i]
		}
	}

	for k := 0; k < len(candies); k++ {
		result = append(result, candies[k]+extraCandies >= maxC)
	}
	return result
}
