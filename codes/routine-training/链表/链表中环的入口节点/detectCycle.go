package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	var fast, slow *ListNode
    fast, slow = head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}
	if fast == nil || fast.Next == nil {
		return nil
	}
	// 重新指向头结点
	slow = head
	// 快慢指针同步前进，相交点就是环起点
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}
