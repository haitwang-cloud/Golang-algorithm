package main

import "container/list"


type CQueue struct {
	stackA, stackB *list.List
}

func Constructor() CQueue {
	return CQueue{
		stackA: list.New(),
		stackB: list.New(),
	}

}

func (this *CQueue) AppendTail(value int) {
	this.stackA.PushBack(value)

}

func (this *CQueue) DeleteHead() int {
	if this.stackB.Len() == 0 {
		for this.stackA.Len() > 0 {
			this.stackB.PushBack(this.stackA.Remove(this.stackA.Back()))
		}
	}
	if this.stackB.Len() != 0 {
		returnValue := this.stackB.Back()
		this.stackB.Remove(returnValue)
		return returnValue.Value.(int)
	}
	return -1
}
