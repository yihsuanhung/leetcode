package main

import "fmt"

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
}

func moveZeroes(nums []int) {

	i, j := 0, 0

	for j < len(nums) {

		if nums[i] == 0 && nums[j] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
		}

		if nums[i] != 0 {
			i++
		}
		j++
	}
}
