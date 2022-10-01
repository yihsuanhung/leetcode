package main

import "fmt"

func main() {
	grid := [][]int{{1, 1, 1, 1, 1, 1, 1, 0}, {1, 0, 0, 0, 0, 1, 1, 0}, {1, 0, 1, 0, 1, 1, 1, 0}, {1, 0, 0, 0, 0, 1, 0, 1}, {1, 1, 1, 1, 1, 1, 1, 0}}
	cl := closedIsland(grid)
	fmt.Print(cl)
}

func closedIsland(grid [][]int) int {

	cl := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				// do dfs
				cl += dfs(grid, i, j)
			}
		}
	}

	return cl

}

func dfs(grid [][]int, i, j int) int {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
		return 0
	}
	if grid[i][j] > 0 {
		return 1
	}
	grid[i][j] = 2
	return dfs(grid, i+1, j) * dfs(grid, i-1, j) * dfs(grid, i, j+1) * dfs(grid, i, j-1)
}
