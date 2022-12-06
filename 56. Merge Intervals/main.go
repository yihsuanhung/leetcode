package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
	fmt.Println(merge([][]int{{1, 4}, {4, 5}}))
}

func merge(intervals [][]int) [][]int {

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		if isOverlap(intervals[i], result[len(result)-1]) {
			pop := result[len(result)-1]
			merged := combine(pop, intervals[i])
			result = result[:len(result)-1]
			result = append(result, merged)
		} else {
			result = append(result, intervals[i])
		}
	}
	return result
}

func isOverlap(a, b []int) bool {
	// if a is in tte left, --a-- --b--
	if a[0] <= b[0] {
		return a[1] >= b[0]
	}
	// if a is in the right, --b-- --a--
	return b[1] >= a[0]
}

func combine(a, b []int) []int {
	return []int{min(a[0], b[0]), max(a[1], b[1])}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
