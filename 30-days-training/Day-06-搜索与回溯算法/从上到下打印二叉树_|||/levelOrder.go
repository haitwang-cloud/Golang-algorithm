package main

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	ans, queue := [][]int{}, list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		currLevel := queue.Len()
		nodeVal := make([]int, 0, currLevel)
		for i := 0; i < currLevel; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			nodeVal = append(nodeVal, node.Val)

			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		ans = append(ans, nodeVal)
	}
	for k := 0; k < len(ans); k++ {
		if k%2 == 1 {
			for i, j := 0, len(ans[k])-1; i < j; i, j = i+1, j-1 {
				ans[k][i], ans[k][j] = ans[k][j], ans[k][i]
			}
		}
	}
	return ans
}
