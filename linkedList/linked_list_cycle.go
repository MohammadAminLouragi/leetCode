package linledList

type CycleListNode struct {
	Val  int
	Next *CycleListNode
}

func DetectCycle(head *CycleListNode) *CycleListNode {

	if head == nil {
        return nil
    }

    slow := head
    fast := head

    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next

        if slow == fast {
            ptr := head
            for ptr != slow {
                ptr = ptr.Next
                slow = slow.Next
            }
            return ptr
        }
    }

    return nil 
}
