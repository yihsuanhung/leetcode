package main

import "fmt"

func main() {
	nums := []int{3, 4, 5, 1, 2}
	fmt.Println(findMin(nums))
}

func findMin(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}

	l, r := 0, len(nums)-1
	for l+1 != r {
		m := (l + r) / 2
		if nums[m] > nums[l] && nums[m] > nums[r] {
			l = m
		} else {
			r = m
		}
		fmt.Println("l", l, "r", r)
	}

	if nums[r] > nums[l] {
		return nums[l]
	}

	return nums[r]
}
