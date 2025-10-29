package main

import (
	"fmt"

	"github.com/MohammadAminLouragi/leetCode/array"
)

func main() {

	nums := []int{1,2,3}
	k := 4
	array.Rotate(nums, k)

	fmt.Println(nums)
}
