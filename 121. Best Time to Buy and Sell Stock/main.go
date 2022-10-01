package main

import "fmt"

func main() {
	prices := []int{7, 6, 4, 3, 1}
	answer := maxProfit(prices)

	fmt.Println(answer)
}

func maxProfit(prices []int) int {
	var min int = prices[0]
	var profit int

	for i := 1; i < len(prices); i++ {

		if prices[i] > min {
			// 有獲利

			if prices[i]-min > profit {
				profit = prices[i] - min
			}

		} else {
			//沒有獲利

			min = prices[i]

		}

	}

	return profit
}
