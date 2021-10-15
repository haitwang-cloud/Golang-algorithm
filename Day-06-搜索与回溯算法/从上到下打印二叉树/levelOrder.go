package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) []int {
	res, queue := []int{}, []*TreeNode{}
	if root == nil {
		return res
	}
	queue = append(queue, root)
	for i := 0; i < len(queue); i++ {
		node := queue[i]
		res = append(res, node.Val)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return res
}
