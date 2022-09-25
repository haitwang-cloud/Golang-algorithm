package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//链接：https://leetcode.cn/problems/er-cha-sou-suo-shu-de-hou-xu-bian-li-xu-lie-lcof/solution/di-gui-fen-zhi-by-lryong-4s7v/

func verifyPostorder(postorder []int) bool {
	var dfs func(i, j int) bool
	dfs = func(i, j int) bool {
		if i >= j {
			return true
		}
		p := i
		for postorder[p] < postorder[j] {
			p++
		}
		m := p
		for postorder[p] > postorder[j] {
			p++
		}

		return p == j && dfs(i, m-1) && dfs(m, j-1)
	}
	return dfs(0, len(postorder)-1)
}
