package main

import "fmt"

func main() {
	fmt.Println(climbStairs(2))
	fmt.Println(climbStairs(3))
}

// memo
func climbStairs(n int) int {
	memo := map[int]int{}
	return mClimbStairs(n, memo)
}

func mClimbStairs(n int, m map[int]int) int {
	// base case
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 0
	}
	if value, ok := m[n]; ok {
		return value
	}
	m[n] = mClimbStairs(n-1, m) + mClimbStairs(n-2, m)
	return m[n]
}

// bruce force
// func climbStairs(n int) int {
// 	// base cases
// 	if n == 0 {
// 		return 1
// 	}
// 	if n < 0 {
// 		return 0
// 	}
// 	return climbStairs(n-1) + climbStairs(n-2)
// }
