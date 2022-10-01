package main

import "fmt"

func main() {

	nums := []int{-1, -2}

	asw := maxSubArray(nums)

	fmt.Print(asw)

}

func maxSubArray(nums []int) int {

	max := nums[0]
	for i := 1; i < len(nums); i++ {

		if nums[i]+nums[i-1] > nums[i] {

			nums[i] += nums[i-1]
		}
		if nums[i] > max {

			max = nums[i]
		}
	}
	return max
}

// func maxSubArray(nums []int) int {

// 	if len(nums) == 1 {
// 		return nums[0]
// 	}

// 	acc := nums[0]
// 	max := nums[0]

// 	for i := 1; i < len(nums); i++ {

// 		curr := nums[i]

// 		if curr > curr+acc {

// 			acc = curr
// 			max = acc

// 		} else {

// 			acc = curr + acc

// 			if acc > max {

// 				max = acc
// 			}

// 		}

// 	}

// 	return max

// }
