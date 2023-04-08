package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	dummyHead := &ListNode{}
	dummyHead.Next = head
	fast, slow := dummyHead, dummyHead
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
	return dummyHead.Next
}
