package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode-cn.com/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/solution/python3-c-by-z1m/
func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}
	return append(reversePrint(head.Next), head.Val)

}
func reversePrintNew(head *ListNode) []int {
	returnResult := []int{}
	for head != nil {
		returnResult = append([]int{head.Val}, returnResult...)
		head = head.Next
	}
	return returnResult
}
