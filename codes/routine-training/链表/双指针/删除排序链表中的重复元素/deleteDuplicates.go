package main

// https://leetcode.cn/problems/remove-duplicates-from-sorted-list/
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	fast, slow := head, head
	for fast != nil {
		if fast.Val != slow.Val {
			// slow指针与快指针相等的值赋值为快指针的值
			slow.Next = fast
			// 将指针后移，操作下一个节点
			slow = slow.Next
		}
		fast = fast.Next
	}
	// 断开与后面重复元素的连接
	slow.Next = nil
	return head
}
