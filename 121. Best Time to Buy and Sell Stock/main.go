package main

import "fmt"

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	answer := maxProfit(prices)

	fmt.Println(answer)
}

func maxProfit(prices []int) int {
	min := prices[0]
	profit := 0

	for _, price := range prices {
		if price < min {
			min = price
		} else {
			if price-min > profit {
				profit = price - min
			}
		}
	}
	return profit
}
