package main

import (
	"fmt"

	linledList "github.com/MohammadAminLouragi/leetCode/linkedList"
)

func main() {

	l1 := &linledList.CycleListNode{Val: 1}
	//l2 := &linledList.CycleListNode{Val: 2}
	// l3 := &linledList.CycleListNode{Val: 0}
	// l4 := &linledList.CycleListNode{Val: -4}
	// l5 := &linledList.CycleListNode{Val: 4}
	// l6 := &linledList.CycleListNode{Val: 5}

	// l1.Next = l2
	// l2.Next = l1
	//  l3.Next = l4
	// l4.Next = l2
	// l5.Next = l6
	// l6.Next = l3

	node := linledList.DetectCycle(l1)
	if node == nil {
		fmt.Println("There is no Cycle")
	}else{
		fmt.Println("Cycle start at Position with value = ",node.Val)
	}
	
}
