package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.cn/problems/middle-of-the-linked-list/

func middleNode(head *ListNode) *ListNode {
	dm := &ListNode{-1, head}
	slow, fast := dm.Next, dm.Next
	for fast.Next != nil && fast != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}
