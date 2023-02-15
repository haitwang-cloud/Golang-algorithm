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
	if head == nil {
		return nil
	}
	dummyHead := &ListNode{}
	dummyHead.Next = head
	cur := dummyHead
	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return dummyHead.Next

}

func main() {
	originList := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	testResultNew := deleteNode(originList, 2)
	fmt.Fprintf(os.Stdout, "%v / %v / %v ", originList, testResultNew)
}
