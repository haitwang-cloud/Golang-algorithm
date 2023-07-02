package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.cn/problems/middle-of-the-linked-list/

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}
