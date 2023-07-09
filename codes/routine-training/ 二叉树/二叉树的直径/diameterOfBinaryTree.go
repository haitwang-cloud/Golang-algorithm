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

	var maxDepth func(node *TreeNode) int
	maxDepth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftMaxDepth := maxDepth(node.Left)
		rightMaxDepth := maxDepth(node.Right)
		// 后序遍历位置顺便计算最大直径
		maxDiameter = max(maxDiameter, leftMaxDepth+rightMaxDepth)
		// 在每次递归调用 maxDepth 时，你需要计算左子树和右子树的最大深度，并将它们加一
		return max(leftMaxDepth, rightMaxDepth) + 1

	}
	maxDepth(root)
	return maxDiameter
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
