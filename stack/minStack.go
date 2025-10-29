package stack

type MinStack struct {
	values []int
	mins   []int
}

func Constructor() MinStack {
	return MinStack{
		values: []int{},
		mins:   []int{},
	}
}

func (this *MinStack) Push(val int) {
	this.values = append(this.values, val)
	if len(this.mins) == 0 || val <= this.mins[len(this.mins)-1] {
		this.mins = append(this.mins, val)
	}
}

func (this *MinStack) Pop() {
	val := this.values[len(this.values)-1]
	this.values = this.values[:len(this.values)-1]
	if val == this.mins[len(this.mins)-1] {
		this.mins = this.mins[:len(this.mins)-1]
	}
}

func (this *MinStack) Top() int {
	val := this.values[len(this.values)-1]
	return val
}

func (this *MinStack) GetMin() int {
	val := this.mins[len(this.mins)-1]
	return val
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
