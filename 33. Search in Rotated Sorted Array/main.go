package main

import "fmt"

func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))    // 4
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2, 3}, 3)) // -1
	fmt.Println(search([]int{1}, 0))                      // -1
}

func search(nums []int, target int) int {

	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}

	l, r := 0, len(nums)-1

	for l+1 != r {
		m := (l + r) / 2
		if target <= nums[m] && nums[r] < nums[l] {
			l = m
		} else {
			r = m
		}
	}

	// if nums[l] == target {
	// 	return l
	// }

	return l
}
