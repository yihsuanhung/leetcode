package main

import "fmt"

func main() {
	fmt.Println(numTrees(1))
	fmt.Println(numTrees(2))
	fmt.Println(numTrees(3))
	fmt.Println(numTrees(4))
}

func numTrees(n int) int {
	trees := make([]int, n+1)
	trees[0], trees[1] = 1, 1

	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			trees[i] += trees[j-1] * trees[i-j]
		}
	}

	return trees[n]
}
