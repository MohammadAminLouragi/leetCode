package stack

import "fmt"

type CustomStack struct {
	values []int
	maxSize int
}

func CustomStackIncreamental(maxSize int) CustomStack {
	return CustomStack{
		values: make([]int, 0),
		maxSize: maxSize,
	}
}

func (this *CustomStack) Push(x int) {
	if len(this.values) < this.maxSize {
		this.values = append(this.values, x)
	}
	this.Println()
}

func (this *CustomStack) Println() {
	for _, n := range this.values {
		fmt.Print(n," --> ")
	}
	fmt.Println()
}

func (this *CustomStack) Pop() int {
	if len(this.values) == 0 {
		return -1
	}
	val := this.values[len(this.values)-1]
	this.values = this.values[:len(this.values)-1]
	return val
}

func (this *CustomStack) Increment(k int, val int) {
	if k >= len(this.values) {
		k = len(this.values)
	}

	j := 0
	for _, n := range this.values[:k] {
		this.values[j] = n + val
		j++
	}
	this.Println()

}