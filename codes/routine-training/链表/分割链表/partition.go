package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	dummyHead1, dummyHead2 := &ListNode{-1, nil}, &ListNode{-1, nil}
	p1, p2 := dummyHead1, dummyHead2
	cur := head
	for cur != nil {
		if cur.Val < x {
			// 小于x的节点
			p1.Next = cur
			p1 = p1.Next
		} else {
			// 大于等于x的节点
			p2.Next = cur
			p2 = p2.Next
		}
		cur = cur.Next
	}
	// 将两个链表连接起来
	p1.Next = dummyHead2.Next
	// 将p2的Next置为nil
	p2.Next = nil
	return dummyHead1.Next
}
