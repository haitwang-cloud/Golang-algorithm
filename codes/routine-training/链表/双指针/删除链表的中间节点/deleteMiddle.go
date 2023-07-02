package main

// https://leetcode.cn/problems/delete-the-middle-node-of-a-linked-list/
type ListNode struct {
	Val  int
	Next *ListNode
}

/*
在本题中，我们还需要删除链表的中间节点，因此除了慢指针 slow 外，
我们再使用一个指针 pre 时刻指向 slow 的前一个节点。
这样我们就可以在遍历结束后，通过 pre 将 slowslow 删除了。
*/
func deleteMiddle(head *ListNode) *ListNode {

	dm := &ListNode{-1, head}
	// fast先走一步，这样slow就是要删除的节点
	fast, slow := dm.Next, dm
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dm.Next

}
