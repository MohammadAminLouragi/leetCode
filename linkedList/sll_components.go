package linledList


type ListNode struct {
	data int
	next *ListNode
}


func numComponents(head *ListNode, nums []int) int {
	h := head
	counter := 0
	for h != nil {
		first := h.data
		second := h.next.data
		find := 0
		for _, val := range nums {
			if first == val {
				find++
			}
			if second == val {
				find++
			}
		}
		if find == 2 {
			h.next = h.next.next
			counter++
		} else {
			h = h.next
		}
	}
	return counter
}
