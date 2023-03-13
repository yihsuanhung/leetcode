package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func kthSmallest(root *TreeNode, k int) int {
	store := []int{}
	inOrder(root, &store, k)
	return store[k-1]
}

func inOrder(node *TreeNode, store *[]int, k int) {
	if node == nil || len(*store) == k {
		return
	}
	inOrder(node.Left, store, k)
	*store = append(*store, node.Val)
	inOrder(node.Right, store, k)
}
