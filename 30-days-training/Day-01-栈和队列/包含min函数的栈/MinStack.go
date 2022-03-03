package main

import (
	"fmt"
	"math"
)

type MinStack struct {
	stack    []int
	minStack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{math.MaxInt64},
	}

}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	top := this.minStack[len(this.minStack)-1]
	this.minStack = append(this.minStack, min(x, top))

}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]

}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]

}

func (this *MinStack) Min() int {
	return this.minStack[len(this.minStack)-1]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x

}

func main() {
	obj := Constructor()
	obj.Push(3)
	obj.Push(1)
	obj.Push(2)
	fmt.Println(obj.Min())
	fmt.Println(obj.Top())
	obj.Pop()
	obj.Pop()
	fmt.Println(obj.Min())
	fmt.Println(obj.Top())

}
