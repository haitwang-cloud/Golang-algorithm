package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left, right := mirrorTree(root.Left), mirrorTree(root.Right)
	root.Left, root.Right = right, left
	return root

}
