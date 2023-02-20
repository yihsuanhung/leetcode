package main

import "fmt"

func main() {
	fmt.Println(floodFill([][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}, 1, 1, 2))
}

func floodFill(img [][]int, sr, sc, newColor int) [][]int {
	dfs(img, sr, sc, img[sr][sc], newColor)
	return img
}

func dfs(img [][]int, r, c, oldColor, newColor int) {
	if r < 0 || r >= len(img) || c < 0 || c >= len(img[0]) || img[r][c] != oldColor {
		return
	}
	img[r][c] = newColor
	dfs(img, r-1, c, oldColor, newColor)
	dfs(img, r+1, c, oldColor, newColor)
	dfs(img, r, c-1, oldColor, newColor)
	dfs(img, r, c+1, oldColor, newColor)
}
