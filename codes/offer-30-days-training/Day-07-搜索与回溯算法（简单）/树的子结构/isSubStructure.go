package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {

	/*
		当节点A和节点B都为空时，匹配成功，返回TRUE
		当节点A或者节点B有一个为空是，说明已经匹配失败,返回FALSE
	*/
	if A == nil && B == nil {
		return true
	}
	if A == nil || B == nil {
		return false
	}

	return validate(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

//validate 校验 B 是否与 A 的一个子树拥有相同的结构和节点值
func validate(A, B *TreeNode) bool {

	/*
		当节点B为空，说明B已经匹配完成,返回TRUE
		当节点A为空, 说明已经越过 A 叶子节点，匹配失败，返回FALSE
		当节点A和节点B的值不相同时，匹配失败，返回FALSE
	*/
	if B == nil {
		return true
	}
	if A == nil || A.Val != B.Val {
		return false
	}
	//a.Val == b.Val 递归校验 A B 左子树和右子树的结构和节点是否相同
	return validate(A.Left, B.Left) && validate(A.Right, B.Right)

}
