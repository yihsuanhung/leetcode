package main

import "fmt"

func main() {
	fmt.Println(insert([][]int{{1, 3}, {6, 9}}, []int{2, 5}))
	fmt.Println(insert([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}))
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

func insert(intervals [][]int, newInterval []int) [][]int {
	var res [][]int
	start, end := newInterval[0], newInterval[1]

	for _, itvl := range intervals {
		currStart, currEnd := itvl[0], itvl[1]
		if currEnd < start {
			// 沒重疊 cur 在左邊 推入 cur-start cur-end
			res = append(res, []int{currStart, currEnd})
		} else if currStart > end {
			// 沒重疊 cur 在右邊 推入 start end 並更新 start end
			res = append(res, []int{start, end})
			start, end = currStart, currEnd
		} else {
			start = min(start, currStart)
			end = max(end, currEnd)
		}
	}
	res = append(res, []int{start, end})

	return res
}
