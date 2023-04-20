package main

// https://leetcode.cn/problems/fan-zhuan-lian-biao-lcof/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {

	cur := head
	var dummyHead *ListNode
	for cur != nil {
		next := cur.Next
		cur.Next = dummyHead
		dummyHead = cur
		cur = next
	}
	return dummyHead

}
