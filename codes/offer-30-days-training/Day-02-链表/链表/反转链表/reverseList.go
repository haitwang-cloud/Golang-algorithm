package main

import (
	"fmt"
	"os"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// leetcode link:https://leetcode-cn.com/problems/reverse-linked-list/solution/fan-zhuan-lian-biao-by-leetcode-solution-d1k2/
// 方法一 迭代
/*

首先定义一个cur指针，指向头结点，再定义一个pre指针，初始化为null。
然后就要开始反转了，首先要把 cur->next 节点用tmp指针保存一下，也就是保存一下这个节点。
为什么要保存一下这个节点呢，因为接下来要改变 cur->next 的指向了，将cur->next 指向pre ，此时已经反转了第一个节点了。
接下来，就是循环走如下代码逻辑了，继续移动pre和cur指针。
最后，cur 指针已经指向了null，循环结束，链表也反转完毕了。 此时我们return pre指针就可以了，pre指针就指向了新的头结点。

*/
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

// 方法二 递归
//https://leetcode.cn/problems/fan-zhuan-lian-biao-lcof/solution/jian-zhi-offer-24-fan-zhuan-lian-biao-by-dint/
func reverseListNew(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseListNew(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead

}

func main() {
	originList := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	//testResult := reverseList(originList)
	testResultNew := reverseListNew(originList)
	fmt.Fprintf(os.Stdout, "%v / %v / %v ", originList, testResultNew)
}
