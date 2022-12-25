package main

type Node struct {
	Val       int
	Neighbors []*Node
}

func main() {}

func cloneGraph(node *Node) *Node {
	ref := map[*Node]*Node{}

	var dfs func(node *Node) *Node

	dfs = func(node *Node) *Node {
		if node == nil {
			return nil
		}
		if n, ok := ref[node]; ok {
			return n
		}
		newNode := &Node{Val: node.Val, Neighbors: make([]*Node, 0, len(node.Neighbors))}

		ref[node] = newNode

		for _, neighbor := range node.Neighbors {
			newNode.Neighbors = append(newNode.Neighbors, dfs(neighbor))

		}
		return newNode
	}

	return dfs(node)
}

// func dfs(node *Node, ref map[*Node]*Node) *Node {
// 	if n, ok := ref[node]; ok {
// 		return n
// 	}
// 	newNode := new(Node)
// 	ref[node] = newNode

// 	for _, n := range node.Neighbors {
// 		newNode.Neighbors = append(newNode.Neighbors, dfs(n, ref))
// 	}

// 	return newNode
// }
