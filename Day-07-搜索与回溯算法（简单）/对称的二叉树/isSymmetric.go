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
	return recur(root.Left, root.Right)
}

func recur(L, R *TreeNode) bool {
	if L == nil && R == nil {
		return true
	}
	if L == nil || R == nil || L.Val != R.Val {
		return false
	}
	return recur(L.Left, R.Right) && recur(L.Right, R.Left)
}

// boolean recur(TreeNode L, TreeNode R) {
// 	if(L == null && R == null) return true;
// 	if(L == null || R == null || L.val != R.val) return false;
// 	return recur(L.left, R.right) && recur(L.right, R.left);
// }
