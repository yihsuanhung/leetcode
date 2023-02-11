package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	return dfs(root, 0)
}

func dfs(node *TreeNode, depth int) int {
	if node == nil {
		return depth
	}
	return max(dfs(node.Left, depth+1), dfs(node.Right, depth+1))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
