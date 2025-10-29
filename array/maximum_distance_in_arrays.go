package array

func MaxDistance(arrays [][]int) int {
	minValue := arrays[0][0]
    maxValue := arrays[0][len(arrays[0])-1]
    result := 0

    for i := 1; i < len(arrays); i++ {
        arr := arrays[i]
        result = max(result, max(abs(arr[len(arr)-1]-minValue), abs(maxValue-arr[0])))
        minValue = min(minValue, arr[0])
        maxValue = max(maxValue, arr[len(arr)-1])
    }
    return result
}

func max(a, b int) int {
    if a > b { return a }
    return b
}

func min(a, b int) int {
    if a < b { return a }
    return b
}

func abs(x int) int {
    if x < 0 { return -x }
    return x
}