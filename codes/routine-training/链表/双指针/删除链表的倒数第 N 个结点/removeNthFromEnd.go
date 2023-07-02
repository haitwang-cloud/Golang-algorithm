package main

// https://leetcode.cn/problems/remove-nth-node-from-end-of-list/
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	dm := &ListNode{-1, head}
	fast, slow := dm, dm
	for i := 0; i < n; i++ {
		if fast.Next != nil {
			fast = fast.Next
		}
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dm.Next
}
