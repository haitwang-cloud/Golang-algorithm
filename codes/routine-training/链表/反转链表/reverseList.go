package main

// https://leetcode.cn/problems/fan-zhuan-lian-biao-lcof/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {

	dm := &ListNode{Next: head, Val: -1}
	cur := dm.Next
	for cur != nil && cur.Next != nil {
		// 1. 保存当前节点的下一个节点
		tmp := cur.Next
		// 2. 将当前节点的下一个节点指向下下个节点
		cur.Next = tmp.Next
		// 3. 将下一个节点指向当前节点
		tmp.Next = dm.Next
		// 4. 将哨兵节点指向下一个节点
		dm.Next = tmp
	}
	return dm.Next
}
