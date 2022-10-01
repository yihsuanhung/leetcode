package main

func main() {}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCompleteTree(root *TreeNode) bool {

	queue := []*TreeNode{}
	queue = append(queue, root)

	empty := false

	for len(queue) > 0 {
		currentNode := queue[0]
		queue = queue[1:]

		if currentNode.Left != nil {
			if empty {
				return false
			}
			queue = append(queue, currentNode.Left)
		} else {
			empty = true
		}

		if currentNode.Right != nil {
			if empty {
				return false
			}
			queue = append(queue, currentNode.Right)
		} else {
			empty = true
		}

	}

	return true
}
