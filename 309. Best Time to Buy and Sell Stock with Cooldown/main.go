package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{1, 2, 3, 0, 2}))
}

// memo 的 key
type Decision struct {
	i   int
	buy bool
}

func maxProfit(prices []int) int {

	memo := make(map[Decision]int)

	// 先宣告一個變數，因為之後要遞迴呼叫
	var dfs func(i int, buy bool) int

	dfs = func(i int, buy bool) int {
		// base cases
		if i >= len(prices) {
			return 0
		}
		if m, ok := memo[Decision{i, buy}]; ok {
			return m
		}

		// dfs 走訪，如果選擇買或賣的話
		if buy {
			// 買的話，獲利要剪掉現在的價格，而且children不能再買
			buying := dfs(i+1, !buy) - prices[i]
			// 冷卻的話，獲利不變，children跟隨當前的buy狀態
			cooldown := dfs(i+1, buy)
			// 找出所有children最大值，回傳給parent
			memo[Decision{i, buy}] = max(buying, cooldown)
		} else {
			// 賣的話，獲利要加上現在的價格，而且children不能再賣
			selling := dfs(i+2, !buy) + prices[i]
			cooldown := dfs(i+1, buy)
			memo[Decision{i, buy}] = max(selling, cooldown)
		}

		return memo[Decision{i, buy}]
	}

	return dfs(0, true)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
