package main

import "math"

func main() {}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return valid(root, math.MinInt, math.MaxInt)
}

func valid(root *TreeNode, l, r int) bool {
	if root == nil {
		return true
	}
	if root.Val <= l || root.Val >= r {
		return false
	}
	return valid(root.Left, l, root.Val) && valid(root.Right, root.Val, r)
}
