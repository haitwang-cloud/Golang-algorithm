package main

type ListNode struct {
	Val  int
	Next *ListNode
}

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
