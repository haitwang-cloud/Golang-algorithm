package main

import (
	"fmt"
	"os"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	pre := &ListNode{-1, nil}
	cur := pre
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 == nil && l2 != nil {
		cur.Next = l2
	} else if l1 != nil && l2 == nil {
		cur.Next = l1
	}

	return pre.Next

}
func mergeTwoListsNew(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	} else if l1.Val < l2.Val {
		l1.Next = mergeTwoListsNew(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoListsNew(l2.Next, l1)
		return l2
	}

}
func main() {
	originList1 := &ListNode{1, &ListNode{2, &ListNode{4, &ListNode{6, nil}}}}
	originList2 := &ListNode{3, &ListNode{5, &ListNode{7, &ListNode{8, nil}}}}
	testResult := mergeTwoLists(originList1, originList2)
	testResultNew := mergeTwoListsNew(originList1, originList2)
	fmt.Fprintf(os.Stdout, "%v / %v / %v ", originList1, testResult, testResultNew)
}
