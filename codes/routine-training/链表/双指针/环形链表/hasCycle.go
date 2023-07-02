package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.cn/problems/linked-list-cycle/
func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}
