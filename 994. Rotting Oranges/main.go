package main

import "fmt"

func main() {
	// grid := [][]int{{0, 2}} // should be 0
	// grid := [][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}} // should be -1
	// grid := [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}} // should be 4
	// grid := [][]int{{2, 1, 1}, {1, 1, 1}, {0, 1, 2}} // should be 2
	grid := [][]int{{2, 2}, {1, 1}, {0, 0}, {2, 0}} // should be 1
	answer := orangesRotting(grid)
	fmt.Println(answer)
}

func orangesRotting(grid [][]int) int {

	queue := [][]int{}
	offsets := [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	time := -1

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 2 {
				queue = append(queue, []int{i, j})
			}
		}
	}

	fmt.Println(queue)

	for len(queue) > 0 {
		size := len(queue)
		for k := 0; k < size; k++ {
			node := queue[0]
			fmt.Println("node", node)
			queue = queue[1:]
			for _, offset := range offsets {
				x := node[0] + offset[0]
				y := node[1] + offset[1]
				if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) && grid[x][y] == 1 {
					grid[x][y] = 2
					queue = append(queue, []int{x, y})
				}
			}
		}
		fmt.Println("======")
		time++
	}

	for l := 0; l < len(grid); l++ {
		for m := 0; m < len(grid[0]); m++ {
			if grid[l][m] == 1 {
				return -1
			}
		}
	}

	if time == -1 {
		return 0
	}

	return time

}
