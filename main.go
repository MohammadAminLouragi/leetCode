package main

import (
	"fmt"

	linledList "github.com/MohammadAminLouragi/leetCode/linkedList"
)

func main() {
	l1 := &linledList.P3ListNode{Val: 2}
	l1.Next = &linledList.P3ListNode{Val: 4}
	l1.Next.Next = &linledList.P3ListNode{Val: 3}

	l2 := &linledList.P3ListNode{Val: 5}
	l2.Next = &linledList.P3ListNode{Val: 6}
	l2.Next.Next = &linledList.P3ListNode{Val: 4}

	// h := l1
	// for h !=nil{
	// 	fmt.Print(h.Val , " -> ")
	// 	h = h.Next
	// }

	// fmt.Println()
	// h = l2
	// for h !=nil{
	// 	fmt.Print(h.Val , " -> ")
	// 	h = h.Next
	// }
	result := linledList.AddTwoNumbers(l1, l2)

	fmt.Println()
	for result !=nil{
		fmt.Print(result.Val , " -> ")
		result = result.Next
	}

}
