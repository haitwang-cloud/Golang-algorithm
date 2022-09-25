package main

import (
	"fmt"
	"os"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

/*

初始化： 前指针 former 、后指针 latter ，双指针都指向头节点 head 。
构建双指针距离： 前指针 former 先向前走k 步（结束后，双指针 former 和 latter 间相距k 步）。
双指针共同移动： 循环中，双指针 former 和 latter 每轮都向前走一步，直至 former 走过链表 尾节点 时跳出（跳出后， latter 与尾节点距离为k−1，即 latter 指向倒数第
k 个节点）。
返回值： 返回 latter 即可。

链接：https://leetcode.cn/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/solution/mian-shi-ti-22-lian-biao-zhong-dao-shu-di-kge-j-11/

*/
func getKthFromEnd(head *ListNode, k int) *ListNode {
	former, latter := head, head
	for index := 0; index < k; index++ {
		former = former.Next
	}
	for former != nil {
		former, latter = former.Next, latter.Next

	}
	return latter

}

func main() {
	originList := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	testResultNew := getKthFromEnd(originList, 2)
	fmt.Fprintf(os.Stdout, "%v / %v / %v ", originList, testResultNew)
}
