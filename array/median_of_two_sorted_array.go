package array

import "math"

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// Ensure nums1 is the smaller array
	if len(nums1) > len(nums2) {
		return FindMedianSortedArrays(nums2, nums1)
	}

	m, n := len(nums1), len(nums2)
	left, right := 0, m

	for left <= right {
		// Partition indices
		partitionA := (left + right) / 2
		partitionB := (m+n+1)/2 - partitionA

		// Get border values (handle out-of-bounds)
		maxLeftA := math.Inf(-1)
		if partitionA > 0 {
			maxLeftA = float64(nums1[partitionA-1])
		}
		minRightA := math.Inf(1)
		if partitionA < m {
			minRightA = float64(nums1[partitionA])
		}

		maxLeftB := math.Inf(-1)
		if partitionB > 0 {
			maxLeftB = float64(nums2[partitionB-1])
		}
		minRightB := math.Inf(1)
		if partitionB < n {
			minRightB = float64(nums2[partitionB])
		}

		// Check if correct partition
		if maxLeftA <= minRightB && maxLeftB <= minRightA {
			// Found correct partition
			if (m+n)%2 == 0 {
				return (math.Max(maxLeftA, maxLeftB) + math.Min(minRightA, minRightB)) / 2.0
			} else {
				return math.Max(maxLeftA, maxLeftB)
			}
		} else if maxLeftA > minRightB {
			right = partitionA - 1
		} else {
			left = partitionA + 1
		}
	}

	// Should never reach here if inputs are valid
	return 0.0
}