package main

// https://leetcode.cn/problems/odd-even-linked-list/?envType=study-plan-v2&envId=leetcode-75
type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	dm1, dm2 := &ListNode{-1, nil}, &ListNode{-1, nil}
	p1, p2 := dm1, dm2
	cur := head
	curIndex := 1
	for cur != nil {
		if curIndex%2 == 1 {
			p1.Next = cur
			p1 = p1.Next
		} else {
			p2.Next = cur
			p2 = p2.Next
		}
		cur = cur.Next
		curIndex++
	}
	p1.Next = dm2.Next
	p2.Next = nil
	return dm1.Next
}
