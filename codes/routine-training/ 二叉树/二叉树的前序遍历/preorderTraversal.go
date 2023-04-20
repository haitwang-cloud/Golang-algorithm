package main

// * Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	// 前序遍历：根左右
	res = append(res, root.Val)
	// 递归遍历左子树
	res = append(res, preorderTraversal(root.Left)...)
	// 递归遍历右子树
	res = append(res, preorderTraversal(root.Right)...)
	return res
}

func preorderTraverse(root *TreeNode) []int {
	//递归实现
	var res []int
	preorder(root, &res)
	return res
}

func preorder(node *TreeNode, res *[]int) {
	if node == nil { //递归停止条件
		return
	} //前序遍历：中左右
	*res = append(*res, node.Val)
	// 递归遍历左子树
	preorder(node.Left, res)
	// 递归遍历右子树
	preorder(node.Right, res)
}
