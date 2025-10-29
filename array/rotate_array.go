package array

func Rotate(nums []int, k int)  {
	j:=len(nums)-1
	if k == 0 {
		return
	}
    for i := 0; i < len(nums); i++ {
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
			j--
		}
	}

	j = k - 1
	if k > len(nums) {
		j = (k % len(nums)) - 1
	}
	for i := 0; i < k; i++ {
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
			j--
		}
	}

	j = len(nums) - 1
	if k > len(nums) {
		k = k % len(nums)
	}
	for i := k; i < len(nums); i++ {
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
			j--
		}
	}
}