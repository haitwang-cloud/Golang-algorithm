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
	// 内置函数 append() 也是一个可变参数的函数。这意味着可以在一次调用中传递多个值。
	// 如果使用 … 运算符，可以将一个切片的所有元素追加到另一个切片里：
	returnResult := []int{}
	for head != nil {
		returnResult = append([]int{head.Val}, returnResult...)
		head = head.Next
	}
	return returnResult
}
