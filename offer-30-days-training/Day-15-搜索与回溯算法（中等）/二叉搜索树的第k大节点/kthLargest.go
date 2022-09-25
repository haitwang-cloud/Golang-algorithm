package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//https://leetcode.cn/problems/er-cha-sou-suo-shu-de-di-kda-jie-dian-lcof/solution/mian-shi-ti-54-er-cha-sou-suo-shu-de-di-k-da-jie-d/

func kthLargest(root *TreeNode, k int) int {
	var dfs func(*TreeNode)
	var ans int
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		dfs(node.Right)
		if k -= 1; k == 0 {
			ans = node.Val
			return
		}
		dfs(node.Left)
	}

	dfs(root)
	return ans

}
