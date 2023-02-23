package main

import (
	"fmt"
	"os"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre

}

func reverse(pre, cur *ListNode) *ListNode {
	if cur == nil {
		return pre
	}
	next := cur.Next
	cur.Next = pre
	return reverse(cur, next)
}

func reverseList2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	return reverse(nil, head)
}

func main() {
	originList := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	testResultNew := reverseList(originList)
	fmt.Fprintf(os.Stdout, "%v / %v / %v ", originList, testResultNew)
}
