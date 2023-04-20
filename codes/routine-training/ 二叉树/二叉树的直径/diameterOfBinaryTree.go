package main

// * Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.cn/problems/diameter-of-binary-tree/

func diameterOfBinaryTree(root *TreeNode) int {
	maxDiameter := 0
	maxDiameter = maxDepth(root, &maxDiameter)
	return maxDiameter
}
func maxDepth(node *TreeNode, maxDiameter *int) int {
	if node == nil {
		return 0
	}
	leftMax := maxDepth(node.Left, maxDiameter)
	rightMax := maxDepth(node.Right, maxDiameter)
	curDiameter := leftMax + rightMax
	if curDiameter > *maxDiameter {
		*maxDiameter = curDiameter
	}
	return max(leftMax, rightMax) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
