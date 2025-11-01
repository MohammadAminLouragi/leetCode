package main

import (
	"fmt"

	"github.com/MohammadAminLouragi/leetCode/stack"
)

// type ListNode struct {
//     Val  int
//     Next *ListNode
// }
// func numComponents(head *ListNode, nums []int) int {
//     set := make(map[int]bool)
//     for _, n := range nums {
//         set[n] = true
//     }

//     count := 0
//     for head != nil {
//         if set[head.Val] && (head.Next == nil || !set[head.Next.Val]) {
//             count++
//         }
//         head = head.Next
//     }

//     return count
// }

// // Helper to build linked list
// func makeList(vals []int) *ListNode {
//     if len(vals) == 0 {
//         return nil
//     }
//     head := &ListNode{Val: vals[0]}
//     curr := head
//     for _, v := range vals[1:] {
//         curr.Next = &ListNode{Val: v}
//         curr = curr.Next
//     }
//     return head
// }

func main() {

	obj := stack.CustomStackIncreamental(3)
	obj.Push(1)
	obj.Push(2)
	fmt.Println(obj.Pop()," Pop")
	obj.Push(2)
	obj.Push(3)
	obj.Push(4)
	obj.Increment(5, 100)
	obj.Increment(2, 100)
	fmt.Println(obj.Pop()," Pop")
	fmt.Println(obj.Pop()," Pop")
	fmt.Println(obj.Pop()," Pop")
	fmt.Println(obj.Pop()," Pop")
	

	// head := makeList([]int{0, 1, 2, 3})
	// nums := []int{0, 1, 3}
	// fmt.Println(numComponents(head, nums)) // Output: 2
}

// func findNumbers(nums []int) int {
//     count := 0
//     for _, val := range nums{
//         dig:= 0
//         for val >=1 {
//             val = val / 10
//             dig++
//         }
//         if dig%2==0{
//             count++
//         }
//     }

//     return count
// }
