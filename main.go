package main

import linledList "github.com/MohammadAminLouragi/leetCode/linkedList"

//"github.com/MohammadAminLouragi/leetCode/array"

// "github.com/MohammadAminLouragi/leetCode/queue"

func main() {

	// arrays := [][]int{{-2}, {-3, -2, 1}}
	// result := array.MaxDistance(arrays)
	// println(result)

	// q := Constructor(3)
	// fmt.Println(q.EnQueue(1))
	// fmt.Println(q.EnQueue(2))
	// fmt.Println(q.EnQueue(3))
	// fmt.Println(q.DeQueue())
	// fmt.Println(q.DeQueue())
	// fmt.Println(q.EnQueue(3))
	// fmt.Println(q.EnQueue(5))
	// fmt.Println(q.DeQueue())
	// fmt.Println(q.DeQueue())
	// fmt.Println(q.DeQueue())
	// fmt.Println(q.DeQueue())

	obj := linledList.Constructor()
	obj.AddAtIndex(2,1)
	obj.AddAtIndex(3,4)
	obj.AddAtTail(1)
	obj.PrintList()
	param_1 := obj.Get(0)
	println(param_1)
	obj.DeleteAtIndex(0)
	param_2 := obj.Get(0)
	println(param_2)
	
}
