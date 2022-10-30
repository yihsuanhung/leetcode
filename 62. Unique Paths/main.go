package main

import "fmt"

func main() {
	fmt.Println(uniquePaths(18, 18))
	fmt.Println(tabulationUniquePath(18, 18))
}

func uniquePaths(m int, n int) int {
	// if m <= 0 || n <= 0 {
	// 	return 0
	// }
	// if m == 1 && n == 1 {
	// 	return 1
	// }
	// return uniquePaths(m-1, n) + uniquePaths(m, n-1)
	// =============================

	memo := map[string]int{}
	return mUniquePaths(m, n, memo)
}

func mUniquePaths(m int, n int, memo map[string]int) int {

	key := string(m) + "," + string(n)
	if m, ok := memo[key]; ok {
		return m
	}

	if m <= 0 || n <= 0 {
		return 0
	}
	if m == 1 || n == 1 {
		return 1
	}

	paths := mUniquePaths(m-1, n, memo) + mUniquePaths(m, n-1, memo)

	memo[key] = paths

	return paths

}

func tabulationUniquePath(m, n int) int {

	// initialize a table
	table := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		row := make([]int, n+1)
		table[i] = row
	}
	table[1][1] = 1

	// iteration
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			current := table[i][j]
			if i+1 <= m {
				table[i+1][j] += current
			}
			if j+1 <= n {
				table[i][j+1] += current
			}
		}
	}

	return table[m][n]
}
