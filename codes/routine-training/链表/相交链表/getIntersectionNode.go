package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p1, p2 := headA, headB
	for p1 != p2 {
		// p1 走一步，如果走到 A 链表末尾，转到 B 链表
		if p1 == nil {
			p1 = headB
		} else {
			p1 = p1.Next
		}
		// p2 走一步，如果走到 B 链表末尾，转到 A 链表
		if p2 == nil {
			p2 = headA
		} else {
			p2 = p2.Next
		}
	}
	return p1
}
