package stack


func ValidateStackSequences(pushed []int, popped []int) bool {

	stack := make([]int, 0)

	j:= 0 

	for _, val := range pushed {
		stack = append(stack, val)
		for len(stack) > 0 && stack[len(stack)-1] == popped[j] {
			stack = stack[:len(stack)-1]
			j++
		}
	}

	return len(stack) == 0
}
