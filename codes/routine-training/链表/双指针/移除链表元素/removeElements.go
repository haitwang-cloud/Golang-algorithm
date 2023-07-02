package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.cn/problems/remove-linked-list-elements/
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	dm := &ListNode{Next: head, Val: -1}
	cur := dm
	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}

	}
	return dm.Next
}
