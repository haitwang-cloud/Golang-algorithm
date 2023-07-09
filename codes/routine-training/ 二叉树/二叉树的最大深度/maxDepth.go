package main

// * Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftMax := maxDepth(root.Left)
	rightMax := maxDepth(root.Right)
	res := max(leftMax, rightMax) + 1
	return res
}

func max(a, b int) int {

	if a > b {
		return a
	}
	return b

}
