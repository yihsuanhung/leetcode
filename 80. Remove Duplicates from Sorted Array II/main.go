package main

import "fmt"

func main() {
	nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	fmt.Println(removeDuplicates(nums))
}
func removeDuplicates(nums []int) int {
	i, j := 2, 2

	if len(nums) < 2 {
		return len(nums)
	}

	for j < len(nums) {
		if nums[j] != nums[i-2] {

			nums[i] = nums[j]
			i++

		}
		j++
	}

	return i
}
