package main

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
链接：https://leetcode.cn/problems/er-cha-shu-zhong-he-wei-mou-yi-zhi-de-lu-jing-lcof/solution/er-cha-shu-zhong-he-wei-mou-yi-zhi-de-lu-68dg/
*/

func pathSum(root *TreeNode, target int) (ans [][]int) {
	path := []int{}
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, left int) {
		if node == nil {
			return
		}
		left -= node.Val
		path = append(path, node.Val)
		defer func() { path = path[:len(path)-1] }()
		if node.Left == nil && node.Right == nil && left == 0 {
			ans = append(ans, append([]int(nil), path...))
			return
		}
		dfs(node.Left, left)
		dfs(node.Right, left)
	}
	dfs(root, target)
	return
}
