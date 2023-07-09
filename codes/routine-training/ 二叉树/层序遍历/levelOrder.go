package main

// https://leetcode.cn/problems/cong-shang-dao-xia-da-yin-er-cha-shu-lcof/
// * Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		// 从队列中取出第一个节点 node，并将其值 node.Val 添加到结果切片 res 中。
		node := queue[0]
		queue = queue[1:]
		res = append(res, node.Val)
		// 将 node 的左子节点和右子节点（如果存在）加入队列中。
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return res
}
