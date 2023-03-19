package main

import "fmt"

func main() {
	fmt.Println(insert([][]int{{1, 3}, {6, 9}}, []int{2, 5}))
	fmt.Println(insert([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}))
}

func insert(intervals [][]int, newInterval []int) [][]int {
	// step 1 insert new interval into intervals
	inserted := false
	mergedIntervals := [][]int{}

	for _, interval := range intervals {
		if interval[0] > newInterval[0] && !inserted {
			mergedIntervals = append(mergedIntervals, newInterval)
			inserted = true
		}
		mergedIntervals = append(mergedIntervals, interval)
	}

	if !inserted {
		mergedIntervals = append(mergedIntervals, newInterval)
	}

	if len(mergedIntervals) == 1 {
		return mergedIntervals
	}

	// step 2 merge intervals
	result := [][]int{}
	result = append(result, mergedIntervals[0])
	for i := 1; i < len(mergedIntervals); i++ {
		if isOverlap(mergedIntervals[i], result[len(result)-1]) {
			result[len(result)-1] = merge(result[len(result)-1], mergedIntervals[i])
		} else {
			result = append(result, mergedIntervals[i])
		}
	}

	return result
}

func merge(a, b []int) []int {
	return []int{min(a[0], b[0]), max(a[1], b[1])}
}

func isOverlap(a, b []int) bool {
	return max(a[0], b[0])-min(a[1], b[1]) <= 0
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// func insert(intervals [][]int, newInterval []int) [][]int {
// 	var res [][]int
// 	start, end := newInterval[0], newInterval[1]

// 	for _, itvl := range intervals {
// 		currStart, currEnd := itvl[0], itvl[1]
// 		if currEnd < start {
// 			// 沒重疊 cur 在左邊 推入 cur-start cur-end
// 			res = append(res, []int{currStart, currEnd})
// 		} else if currStart > end {
// 			// 沒重疊 cur 在右邊 推入 start end 並更新 start end
// 			res = append(res, []int{start, end})
// 			start, end = currStart, currEnd
// 		} else {
// 			start = min(start, currStart)
// 			end = max(end, currEnd)
// 		}
// 	}
// 	res = append(res, []int{start, end})

// 	return res
// }
