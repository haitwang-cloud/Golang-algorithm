package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// leetcode link:https://leetcode-cn.com/problems/reverse-linked-list/solution/fan-zhuan-lian-biao-by-leetcode-solution-d1k2/
// 方法一 迭代
// 在遍历链表时，将当前节点的 next 指针改为指向前一个节点。由于节点没有引用其前一个节点，
// 因此必须事先存储其前一个节点。在更改引用之前，还需要存储后一个节点。最后返回新的头引用。
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var pre *ListNode
	curr := head
	for curr != nil {
		// 首先取next，然后将curr的next指向pre，然后将往后迭代将curr赋值给curr，最后
		next := curr.Next
		curr.Next = pre
		pre = curr
		curr = next
	}
	return pre

}

// 方法二 递归
func reverseListNew(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseListNew(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead

}
