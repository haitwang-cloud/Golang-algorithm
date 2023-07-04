package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum(head *ListNode) int {
	dm := &ListNode{-1, head}
	fast, slow := dm.Next, dm
	// 1. 找到中间节点
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	maxSum := 0

	// 1. 反转后半部分链表
	slow = reverseList(slow.Next)
	fast = dm.Next
	// 2. 计算最大孪生和
	for slow != nil {
		if slow.Val+fast.Val > maxSum {
			maxSum = slow.Val + fast.Val
		}
		slow = slow.Next
		fast = fast.Next
	}

	return maxSum
}

func reverseList(head *ListNode) *ListNode {
	dm := &ListNode{Next: head, Val: -1}
	cur := dm.Next
	for cur != nil && cur.Next != nil {
		tmp := cur.Next
		cur.Next = tmp.Next
		tmp.Next = dm.Next
		dm.Next = tmp
	}
	return dm.Next
}
