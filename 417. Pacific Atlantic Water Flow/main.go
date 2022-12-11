package main

import "fmt"

func main() {
	fmt.Println(pacificAtlantic([][]int{{1, 2, 2, 3, 5}, {3, 2, 3, 4, 4}, {2, 4, 5, 3, 1}, {6, 7, 1, 4, 5}, {5, 1, 1, 2, 4}}))
	// fmt.Println(pacificAtlantic([][]int{{1, 1}, {1, 1}, {1, 1}}))
}

func pacificAtlantic(heights [][]int) [][]int {
	result := [][]int{}

	dest := make([][]int, len(heights))
	for i := range dest {
		dest[i] = make([]int, len(heights[0]))
	}

	for i := range heights {
		for j := range heights[0] {
			if i == 0 || j == 0 {
				// if pacific => mark as 1
				dfs(heights, i, j, 1, dest, &result)
			}
		}
	}

	for i := range heights {
		for j := range heights[0] {
			if i == len(heights)-1 || j == len(heights[0])-1 {
				// if atlantic
				dfs(heights, i, j, 2, dest, &result)
			}
		}
	}

	return result
}

func dfs(heights [][]int, i, j, target int, dest [][]int, result *[][]int) {

	// if out of boundary
	if i < 0 || i >= len(heights) || j < 0 || j >= len(heights[0]) || dest[i][j] == target {
		return
	}

	// if can flow to both ocean
	if dest[i][j] == 1 && target == 2 {
		*result = append(*result, []int{i, j})
	}

	dest[i][j] = target // target 1 => pacific, 2 => atlantic

	// dfs, reverse flow
	// up
	if i > 0 && heights[i][j] <= heights[i-1][j] {
		dfs(heights, i-1, j, target, dest, result)
	}
	// down
	if i < len(heights)-1 && heights[i][j] <= heights[i+1][j] {
		dfs(heights, i+1, j, target, dest, result)
	}
	// left
	if j > 0 && heights[i][j] <= heights[i][j-1] {
		dfs(heights, i, j-1, target, dest, result)
	}
	// right
	if j < len(heights[0])-1 && heights[i][j] <= heights[i][j+1] {
		dfs(heights, i, j+1, target, dest, result)
	}

}

// =======================

// func pacificAtlantic(heights [][]int) [][]int {
// 	result := [][]int{}

// 	dest := make([][]int, len(heights))
// 	for i := range dest {
// 		dest[i] = make([]int, len(heights[0]))
// 	}
// 	visited := make([][]int, len(heights))
// 	for i := range visited {
// 		visited[i] = make([]int, len(heights[0]))
// 	}

// 	for i := range heights {
// 		for j := range heights[0] {
// 			if i == 0 || j == 0 {
// 				// if pacific => mark as 1
// 				dfs(heights, i, j, dest, visited)
// 			}
// 		}
// 	}

// 	for i := range visited {
// 		for j := range visited[0] {
// 			visited[i][j] = 0
// 		}
// 	}

// 	for i := range heights {
// 		for j := range heights[0] {
// 			if i == len(heights)-1 || j == len(heights[0])-1 {
// 				// if atlantic
// 				dfs(heights, i, j, dest, visited)
// 			}
// 		}
// 	}

// 	for i := range dest {
// 		for j := range dest[0] {
// 			if dest[i][j] == 2 {
// 				result = append(result, []int{i, j})
// 			}
// 		}
// 	}

// 	return result
// }

// func dfs(heights [][]int, i, j int, dest, visited [][]int) {

// 	// if out of boundary
// 	if i < 0 || i >= len(heights) || j < 0 || j >= len(heights[0]) || visited[i][j] == 1 {
// 		return
// 	}

// 	dest[i][j]++
// 	visited[i][j] = 1

// 	// dfs, reverse flow
// 	// up
// 	if i > 0 && heights[i][j] <= heights[i-1][j] {
// 		dfs(heights, i-1, j, dest, visited)
// 	}
// 	// down
// 	if i < len(heights)-1 && heights[i][j] <= heights[i+1][j] {
// 		dfs(heights, i+1, j, dest, visited)
// 	}
// 	// left
// 	if j > 0 && heights[i][j] <= heights[i][j-1] {
// 		dfs(heights, i, j-1, dest, visited)
// 	}
// 	// right
// 	if j < len(heights[0])-1 && heights[i][j] <= heights[i][j+1] {
// 		dfs(heights, i, j+1, dest, visited)
// 	}

// }

// ==============
