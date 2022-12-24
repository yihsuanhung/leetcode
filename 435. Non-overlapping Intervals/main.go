package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(eraseOverlapIntervals([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}))
	fmt.Println(eraseOverlapIntervals([][]int{{1, 2}, {1, 2}, {1, 2}}))
	fmt.Println(eraseOverlapIntervals([][]int{{1, 2}, {2, 3}}))
}

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	removeCount := 0
	prevEnd := intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		current := intervals[i]
		if current[0] >= prevEnd {
			prevEnd = current[1]
		} else {
			prevEnd = min(current[1], prevEnd)
			removeCount++
		}
	}

	return removeCount
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
