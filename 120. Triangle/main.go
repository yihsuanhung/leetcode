package main

import (
	"fmt"
)

func main() {
	fmt.Println(minimumTotal([][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}))
}

type Element struct {
	i int
	j int
}

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 1 {
		return triangle[0][0]
	}
	memo := make(map[Element]int)
	return findMin(triangle, 0, 0, memo)
}

func findMin(triangle [][]int, i, j int, memo map[Element]int) int {
	if i == len(triangle) {
		return 0
	}
	if m, ok := memo[Element{i, j}]; ok {
		return m
	}

	localMin := min(findMin(triangle, i+1, j, memo)+triangle[i][j], findMin(triangle, i+1, j+1, memo)+triangle[i][j])
	memo[Element{i, j}] = localMin

	return localMin
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
