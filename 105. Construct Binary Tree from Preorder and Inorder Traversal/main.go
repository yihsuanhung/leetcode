package main

func main() {}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	// base csae
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	midIndex := findIndex(inorder, preorder[0])
	root.Left = buildTree(preorder[1:midIndex+1], inorder[:midIndex])
	root.Right = buildTree(preorder[midIndex+1:], inorder[midIndex+1:])

	return root
}

func findIndex(array []int, target int) int {
	for i, element := range array {
		if element == target {
			return i
		}
	}
	return -1
}
