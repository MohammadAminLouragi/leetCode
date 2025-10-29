package queue

type MyCircularQueue struct {
	data []int
	front int
	rear int
	size int
	capacity int
    
}


func Constructor(k int) MyCircularQueue {
    return MyCircularQueue{
		data: make([]int, k),
		front: 0,
		rear: 0,
		size: 0,
		capacity: k,
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	this.data[this.rear] = value
    this.rear++
    if this.rear == this.capacity {
        this.rear = 0
    }
	this.size++
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.front++
    if this.front == this.capacity {
        this.front = 0
    }
	this.size--
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.data[this.front]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
        return -1
    }
    idx := this.rear - 1
    if idx < 0 {
        idx = this.capacity - 1
    }
    return this.data[idx]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.size == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.size == this.capacity
}