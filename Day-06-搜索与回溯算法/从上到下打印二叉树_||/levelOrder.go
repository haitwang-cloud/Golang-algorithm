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

	ans := [][]int{}
	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		curLevelNodes := queue.Len()
		nodeVals := make([]int, 0, curLevelNodes)
		for i := 0; i < curLevelNodes; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			nodeVals = append(nodeVals, node.Val)

			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		ans = append(ans, nodeVals)
	}
	return ans

}
