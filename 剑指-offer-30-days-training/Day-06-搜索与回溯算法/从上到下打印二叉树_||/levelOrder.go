package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}
	overallLevel := []*TreeNode{root}
	for i := 0; len(overallLevel) > 0; i++ {
		ret = append(ret, []int{})
		currentLevel := []*TreeNode{}
		for j := 0; j < len(overallLevel); j++ {
			node := overallLevel[j]
			ret[i] = append(ret[i], node.Val)
			if node.Left != nil {
				currentLevel = append(currentLevel, node.Left)
			}
			if node.Right != nil {
				currentLevel = append(currentLevel, node.Right)
			}
		}
		overallLevel = currentLevel
	}
	return ret
}
