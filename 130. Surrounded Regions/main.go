package main

import "fmt"

func main() {
	// solve([][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}})
	// solve([][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'O', 'X', 'X'}})
	solve([][]byte{{'X', 'O', 'X', 'O', 'X', 'O'}, {'O', 'X', 'O', 'X', 'O', 'X'}, {'X', 'O', 'X', 'O', 'X', 'O'}, {'O', 'X', 'O', 'X', 'O', 'X'}})

}

func solve(board [][]byte) {
	fmt.Println("before", board)

	for i := range board {
		for j := range board[0] {
			if (i == 0 || i == len(board)-1 || j == 0 || j == len(board[0])-1) && board[i][j] == 'O' {
				dfs(board, i, j) // fill item that is on the boundary as 1
			}
		}
	}

	for i := range board {
		for j := range board[0] {
			if board[i][j] == '1' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}

	fmt.Println("after", board)

}

func dfs(board [][]byte, i, j int) {
	// out of boundary
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || board[i][j] != 'O' {
		return
	}

	board[i][j] = '1'
	dfs(board, i+1, j)
	dfs(board, i-1, j)
	dfs(board, i, j-1)
	dfs(board, i, j+1)

	return
}
