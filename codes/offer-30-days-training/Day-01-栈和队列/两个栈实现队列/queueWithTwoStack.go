package main

import (
	"fmt"
)

type CQueue struct {
	inStack  []int
	outStack []int
}

func Constructor() CQueue {
	return CQueue{
		inStack:  []int{},
		outStack: []int{},
	}

}

func (this *CQueue) AppendTail(value int) {
	this.inStack = append(this.inStack, value)

}

func (this *CQueue) DeleteHead() int {
	if len(this.outStack) == 0 {
		for len(this.inStack) > 0 {
			this.outStack = append(this.outStack, this.inStack[len(this.inStack)-1])
			this.inStack = this.inStack[:len(this.inStack)-1]
		}
	}
	if len(this.outStack) != 0 {
		returnValue := this.outStack[len(this.outStack)-1]
		this.outStack = this.outStack[:len(this.outStack)-1]
		return returnValue

	}
	return -1
}

func main() {
	obj := Constructor()
	obj.AppendTail(1)
	obj.AppendTail(2)
	value := obj.DeleteHead()
	fmt.Println(value)
	obj.AppendTail(3)
	value = obj.DeleteHead()
	fmt.Println(value)
	value = obj.DeleteHead()
	fmt.Println(value)

}
