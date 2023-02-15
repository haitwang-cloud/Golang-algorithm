package main

import (
	"fmt"
)

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElementsRecursion(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	head.Next = removeElements(head.Next, val)
	if head.Val == val {
		return head.Next
	} else {
		return head
	}

}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	dummyHead := &ListNode{}
	dummyHead.Next = head
	cur := dummyHead
	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return dummyHead.Next
}

func main() {
	originList := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{2, nil}}}}
	testResultNew := removeElements(originList, 2)
	fmt.Println(testResultNew)
}
