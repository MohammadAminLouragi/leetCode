package array

import "math"

func MaxDistance(arrays [][]int) int {
	maxDistance := 0
	subArray := []int{}
	maxArray := []int{}
	minArray := []int{}

	for _, array := range arrays {
		subArray = append(subArray, array[0], array[len(array)-1])
	}
	if len(subArray)%2 != 0 {
		for i := 0; i < len(subArray); i += 2 {
			if subArray[i] < subArray[i+1] {
				minArray = append(minArray, subArray[i])
				maxArray = append(maxArray, subArray[i+1])
			} else {
				minArray = append(minArray, subArray[i+1])
				maxArray = append(maxArray, subArray[i])
			}

		}
	} else {
		for i := 1; i < len(subArray); i += 2 {
			if subArray[i] < subArray[i+1] {
				minArray = append(minArray, subArray[i])
				maxArray = append(maxArray, subArray[i+1])
			} else {
				minArray = append(minArray, subArray[i+1])
				maxArray = append(maxArray, subArray[i])
			}
		}
		minArray = append(minArray, subArray[0])
		maxArray = append(maxArray, subArray[0])
	}

	min := minArray[0]
	for i := 0; i < len(minArray); i++ {
		if min > minArray[i] {
			min = minArray[i]
		}
	}

	max := maxArray[0]
	for i := 0; i < len(maxArray); i++ {
		if max < maxArray[i] {
			max = maxArray[i]
		}
	}

	maxDistance = int(math.Abs(float64(max - min)))

	return maxDistance
}
