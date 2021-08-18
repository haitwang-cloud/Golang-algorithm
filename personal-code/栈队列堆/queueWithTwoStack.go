package main

import "container/list"

//type CQueue struct {
//	stack1, stack2 *list.List
//}
//
//func Constructor() CQueue {
//	return CQueue{
//		stack1: list.New(),
//		stack2: list.New(),
//	}
//}
//
//func (this *CQueue) AppendTail(value int)  {
//	this.stack1.PushBack(value)
//}
//
//func (this *CQueue) DeleteHead() int {
//	if this.stack2.Len() == 0 {
//		for this.stack1.Len() > 0 {
//			this.stack2.PushBack(this.stack1.Remove(this.stack1.Back()))
//		}
//	}
//	if this.stack2.Len() != 0 {
//		e := this.stack2.Back()
//		this.stack2.Remove(e)
//		return e.Value.(int)
//	}
//	return -1
//}

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

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
