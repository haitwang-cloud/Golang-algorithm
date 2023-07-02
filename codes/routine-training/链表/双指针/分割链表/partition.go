package main

// https://leetcode.cn/problems/partition-list/

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	dm1, dm2 := &ListNode{-1, nil}, &ListNode{-1, nil}
	p1, p2 := dm1, dm2
	cur := head
	for cur != nil {
		if cur.Val < x {
			p1.Next = cur
			p1 = p1.Next
		} else {
			p2.Next = cur
			p2 = p2.Next
		}
		cur = cur.Next
	}
	p1.Next = dm2.Next
	p2.Next = nil
	return dm1.Next
}
