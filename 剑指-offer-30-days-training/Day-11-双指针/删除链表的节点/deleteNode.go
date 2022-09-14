package main

import (
	"fmt"
	"os"
)

//链接：https://leetcode.cn/problems/shan-chu-lian-biao-de-jie-dian-lcof/solution/mian-shi-ti-18-shan-chu-lian-biao-de-jie-dian-sh-2/

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {
	if head.Val == val {
		return head.Next
	}
	pre, cur := head, head.Next
	for cur != nil && cur.Val != val {
		pre, cur = cur, cur.Next
	}
	if cur != nil {
		pre.Next = cur.Next
	}
	return head

}

func main() {
	originList := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	testResultNew := deleteNode(originList, 2)
	fmt.Fprintf(os.Stdout, "%v / %v / %v ", originList, testResultNew)
}
