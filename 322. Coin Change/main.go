package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
	fmt.Println(tabulation([]int{1, 2, 5}, 11))
	fmt.Println(coinChange([]int{2}, 3))
	fmt.Println(coinChange([]int{1}, 0))
}

func coinChange(coins []int, amount int) int {

	memo := make(map[int]int)

	return getCoinChange(coins, amount, memo)
}

func getCoinChange(coins []int, amount int, memo map[int]int) int {
	if m, ok := memo[amount]; ok {
		return m
	}
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}

	minCount := -1

	for _, coin := range coins {
		remaining := amount - coin
		count := getCoinChange(coins, remaining, memo)

		if count != -1 {
			if minCount == -1 || count+1 < minCount {
				minCount = count + 1
			}
		}
	}

	memo[amount] = minCount
	return minCount
}

func tabulation(coins []int, amount int) int {
	dp := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		minCoin := math.MaxInt

		for _, coin := range coins {
			remaining := i - coin
			if remaining >= 0 && dp[remaining] != -1 {
				// dp[remaining] != -1 表示現在的 amount 可以用 coin 換出來
				// dp[remaining]+1 表示可以用 remaining 的方式換出來以外，又多一種，就是現在這種
				minCoin = min(minCoin, dp[remaining]+1)
			}
		}

		if minCoin == math.MaxInt {
			dp[i] = -1
		} else {
			dp[i] = minCoin
		}
	}

	return dp[amount]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
