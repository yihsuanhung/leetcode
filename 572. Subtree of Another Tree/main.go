package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil && subRoot == nil {
		return true
	}
	if root == nil || subRoot == nil {
		return false
	}
	if root.Val == subRoot.Val {
		return isSubtree(root.Left, subRoot.Left) && isSubtree(root.Right, subRoot.Right)
	}
	return isSubtree(root.Left, subRoot.Left) || isSubtree(root.Right, subRoot.Right)
}
