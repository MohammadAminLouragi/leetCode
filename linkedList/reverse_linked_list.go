package linledList


type ListNodes struct {
	Val  int
	Next *ListNodes
}

func reverseBetween(head *ListNodes, left int, right int) *ListNodes {
    if head == nil || left == right {
        return head
    }

    dummy := &ListNodes{Next: head}
    prev := dummy

    // Step 1: Move prev to the node before "left"
    for i := 0; i < left-1; i++ {
        prev = prev.Next
    }

    // Reverse the sublist using head insertion
    curr := prev.Next
    for i := 0; i < right-left; i++ {
        temp := curr.Next
        curr.Next = temp.Next
        temp.Next = prev.Next
        prev.Next = temp
    }

    return dummy.Next
}