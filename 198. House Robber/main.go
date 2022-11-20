package main

import "fmt"

func main() {
	fmt.Println(rob([]int{2, 7, 9, 3, 1}))
}

func rob(nums []int) int {
	prev1, prev2 := 0, 0

	// [ prev1, prev2, n, n+1, ...]
	for _, num := range nums {
		prev1, prev2 = prev2, max(num+prev1, prev2)
	}

	return prev2
}

// type Decision struct {
// 	idx int
// 	rob bool
// }

// func rob(nums []int) int {

// 	memo := make(map[Decision]int)

// 	var dfs func(i int, robPrev bool, memo map[Decision]int) int

// 	dfs = func(i int, robPrev bool, memo map[Decision]int) int {
// 		if i >= len(nums) {
// 			return 0
// 		}
// 		if m, ok := memo[Decision{i, robPrev}]; ok {
// 			return m
// 		}

// 		total := 0

// 		if robPrev {
// 			total += dfs(i+2, !robPrev, memo) + nums[i]
// 		} else {
// 			rob := dfs(i+2, !robPrev, memo) + nums[i]
// 			skip := dfs(i+1, !robPrev, memo)
// 			total += max(rob, skip)
// 		}

// 		memo[Decision{i, robPrev}] = total
// 		return total
// 	}

// 	return dfs(0, false, memo)

// }

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
