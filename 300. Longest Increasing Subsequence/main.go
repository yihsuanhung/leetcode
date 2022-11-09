package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
}

func lengthOfLIS(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	result := math.MinInt
	p := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		currentMax := math.MinInt
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				currentMax = max(currentMax, p[j])
			}
		}
		if currentMax == math.MinInt {
			p[i] = 1
		} else {
			p[i] = currentMax + 1
		}
		result = max(result, p[i])
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
