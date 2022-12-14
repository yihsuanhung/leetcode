package main

import "fmt"

func main() {
	l1 := []byte{'1', '1', '1', '1', '0'}
	l2 := []byte{'1', '1', '0', '1', '0'}
	l3 := []byte{'1', '1', '0', '0', '0'}
	l4 := []byte{'0', '0', '0', '0', '0'}
	fmt.Println(numIslands([][]byte{l1, l2, l3, l4}))                                                                                             // 1
	fmt.Println(numIslands([][]byte{{'1', '1', '0', '0', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '1', '0', '0'}, {'0', '0', '0', '1', '1'}})) // 3
}

func numIslands(grid [][]byte) int {

	num := 0

	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == '1' {
				num++
				bfs(grid, i, j)
			}
		}
	}

	return num
}

func dfs(grid [][]byte, i, j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] != '1' {
		return
	}
	grid[i][j] = 'v' // visited
	dfs(grid, i-1, j)
	dfs(grid, i+1, j)
	dfs(grid, i, j-1)
	dfs(grid, i, j+1)
}

func bfs(grid [][]byte, i, j int) {
	queue := [][]int{}
	queue = append(queue, []int{i, j}) // store coordinates
	offsets := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for len(queue) > 0 {
		pop := queue[0]
		queue = queue[1:] // pop
		for _, offset := range offsets {
			x, y := pop[0]+offset[0], pop[1]+offset[1]
			if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) && grid[x][y] == '1' {
				grid[x][y] = 'v' // visited
				queue = append(queue, []int{x, y})
			}
		}
	}
}
