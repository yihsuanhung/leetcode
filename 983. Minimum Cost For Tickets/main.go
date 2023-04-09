package main

import "fmt"

func main() {
	fmt.Println(mincostTickets([]int{1, 4, 6, 7, 8, 20}, []int{2, 7, 15}))                     // 11
	fmt.Println(mincostTickets([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31}, []int{2, 7, 15})) // 17

}
func mincostTickets(days []int, costs []int) int {

	dayMap := make(map[int]bool)

	memo := make(map[int]int)

	for _, d := range days {
		dayMap[d] = true
	}

	return dp(1, dayMap, costs, memo)
}

func dp(day int, dayMap map[int]bool, costs []int, memo map[int]int) int {
	// base cases
	if day > 356 {
		return 0
	}
	// memo
	if m, ok := memo[day]; ok {
		return m
	}

	cost := 0
	// check if day is one of the traveling days
	if _, ok := dayMap[day]; ok {
		cost = min(min(costs[0]+dp(day+1, dayMap, costs, memo), costs[1]+dp(day+7, dayMap, costs, memo)), costs[2]+dp(day+30, dayMap, costs, memo))

	} else {
		cost = dp(day+1, dayMap, costs, memo)
	}

	return cost
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
