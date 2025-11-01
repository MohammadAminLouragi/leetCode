package daily

import "fmt"

func SmallestNumber(n int) int {
	t := n + 1

	p := int(1)
	fmt.Println(p)

	for  p < t {
		p <<= 1
	
		fmt.Println(p)
	}

	return p - 1
}