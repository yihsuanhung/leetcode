package main

import "fmt"

func main() {
	l1 := []byte{'1', '1', '1', '1', '0'}
	l2 := []byte{'1', '1', '0', '1', '0'}
	l3 := []byte{'1', '1', '0', '0', '0'}
	l4 := []byte{'0', '0', '0', '0', '0'}
	grid := [][]byte{l1, l2, l3, l4}
	answer := numIslands(grid)
	fmt.Println(answer)
}

func numIslands(grid [][]byte) int {

	islands := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				islands++
				dfs(grid, i, j)
				// bfs(grid, i, j)
			}
		}
	}

	return islands
}

func dfs(grid [][]byte, i int, j int) {
	// 檢查是否出界 + 檢查是否visited + 檢查是否非陸地
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] != '1' {
		return
	}

	grid[i][j] = 'v'

	dfs(grid, i-1, j)
	dfs(grid, i+1, j)
	dfs(grid, i, j-1)
	dfs(grid, i, j+1)
}

func bfs(grid [][]byte, i int, j int) {
	queue := [][]int{}
	queue = append(queue, []int{i, j})
	grid[i][j] = 'v'

	offsets := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		for _, offset := range offsets {
			x := current[0] + offset[0]
			y := current[1] + offset[1]
			if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) && grid[x][y] == '1' {
				grid[x][y] = 'v'
				queue = append(queue, []int{x, y})
			}
		}
	}
}
