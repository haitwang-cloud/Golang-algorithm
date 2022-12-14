package main

import (
	"fmt"
	"os"
)

/*

当链表 headA 和 headB 都不为空时，创建两个指针pA 和 pB，初始时分别指向两个链表的头节点 headA 和 headB，然后将两个指针依次遍历两个链表的每个节点。具体做法如下：
每步操作需要同时更新指针 pA 和pB。
如果指针 pA 不为空，则将指针pA 移到下一个节点；
如果指针pB 不为空，则将指针 pB 移到下一个节点。
如果指针pA 为空，则将指针pA 移到链表headB 的头节点；
如果指针pB为空，则将指针pB 移到链表headA 的头节点。
当指针pA 和 pB 指向同一个节点或者都为空时，返回它们指向的节点或者null

链接：https://leetcode.cn/problems/liang-ge-lian-biao-de-di-yi-ge-gong-gong-jie-dian-lcof/solution/liang-ge-lian-biao-de-di-yi-ge-gong-gong-pzbs/
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pA, pB := headA, headB
	for pA != pB {
		if pA == nil {
			pA = headB
		} else {
			pA = pA.Next
		}
		if pB == nil {
			pB = headA
		} else {
			pB = pB.Next
		}
	}
	return pA
}
func main() {
	originList1 := &ListNode{1, &ListNode{2, &ListNode{4, &ListNode{6, nil}}}}
	originList2 := &ListNode{3, &ListNode{2, &ListNode{4, &ListNode{6, nil}}}}
	testResult := getIntersectionNode(originList1, originList2)
	fmt.Fprintf(os.Stdout, "%v / %v ", originList1, testResult)
}
