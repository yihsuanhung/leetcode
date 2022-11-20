package main

import "fmt"

func main() {
	fmt.Println(combinationSum4([]int{1, 2, 3}, 4))
}

func combinationSum4(nums []int, target int) int {
	memo := make(map[int]int)
	return getCombination(nums, target, memo)
}

func getCombination(nums []int, target int, memo map[int]int) int {
	// base cases
	if m, ok := memo[target]; ok {
		return m
	}
	if target == 0 {
		return 1
	}
	if target < 0 {
		return 0
	}

	var combination int

	for _, num := range nums {
		remaining := target - num
		combination += getCombination(nums, remaining, memo)
	}

	memo[target] = combination

	return memo[target]

}
