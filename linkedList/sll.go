package linledList

type MyLinkedList struct {
	size int
	head *Node
	tail *Node
}

type Node struct {
	value int
	next  *Node
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		size: 0,
		head: nil,
		tail: nil,
	}
}

func (this *MyLinkedList) Get(index int) int {
	h := this.head
	if h == nil || index < 0 || index >= this.size {
		return -1
	}
	for i := 0; i < index; i++ {
		h = h.next
	}
	return h.value
}

func (this *MyLinkedList) AddAtHead(val int) {
	h := this.head
	newNode := &Node{
		value: val,
		next:  h,
	}
	this.head = newNode
	if this.size == 0 {
		this.tail = newNode
	}
	this.size++

}

func (this *MyLinkedList) AddAtTail(val int) {
	t := this.tail
	newNode := &Node{
		value: val,
		next:  nil,
	}
	if t != nil {
		t.next = newNode
	}
	this.tail = newNode
	if this.size == 0 {
		this.head = newNode
	}
	this.size++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	h := this.head
	newNode := &Node{
		value: val,
		next:  nil,
	}
	if index == 0 {
		this.AddAtHead(val)
		return
	}
	if index == this.size {
		this.AddAtTail(val)
		return
	}

	if index > 0 && index < this.size {
		prev := h
		for i := 0; i < index-1; i++ {
			prev = prev.next
		}
		newNode.next = prev.next
		prev.next = newNode
		this.size++
	}

}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	h := this.head
	if h == nil || index < 0 || index >= this.size {
		return
	}
	if index == 0 {
		this.head = h.next
		if this.size == 1 {
			this.tail = nil
		}
		this.size--
		return
	}
	prev := h
	for i := 0; i < index-1; i++ {
		prev = prev.next
	}
	prev.next = prev.next.next
	if index == this.size-1 {
		this.tail = prev
	}
	this.size--
}

func (this *MyLinkedList) PrintList() {
	h := this.head
	for h != nil {
		print(h.value, " ")
		h = h.next
	}
	println()
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
