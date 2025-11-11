package array

func FindMaxLength(nums []int) int {
	prefixIndex := map[int]int{0: -1} // prefix sum → first index
	count := 0
	maxLen := 0

	for i, num := range nums {
		if num == 1 {
			count++
		} else {
			count--
		}

		if idx, exists := prefixIndex[count]; exists {
			length := i - idx
			if length > maxLen {
				maxLen = length
			}
		} else {
			prefixIndex[count] = i
		}
	}

	return maxLen
}
