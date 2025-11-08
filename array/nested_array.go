package array

func ArrayNesting(nums []int) int {
	n := len(nums)
	visited := make([]bool, n)
	maxLen := 0

	for i := 0; i < n; i++ {
		if !visited[i] {
			length := 0
			j := i
			for !visited[j] {
				visited[j] = true
				j = nums[j]
				length++
			}
			if length > maxLen {
				maxLen = length
			}
		}
	}

	return maxLen
}
