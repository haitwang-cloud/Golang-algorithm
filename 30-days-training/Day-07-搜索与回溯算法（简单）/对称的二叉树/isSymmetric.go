package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return validate(root.Left, root.Right)
}

func validate(L, R *TreeNode) bool {

	/*
		当节点A和节点B都为空时，匹配成功，返回TRUE
		当节点A为空，或节点B为空，或节点A和节点B的值不相同，匹配失败，返回FALSE
	*/
	if L == nil && R == nil {
		return true
	}
	if L == nil || R == nil || L.Val != R.Val {
		return false
	}
	return validate(L.Left, R.Right) && validate(L.Right, R.Left)
}
