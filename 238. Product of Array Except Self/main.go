package main

import "fmt"

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
	fmt.Println(productExceptSelf([]int{-1, 1, 0, -3, 3}))
}

func productExceptSelf(nums []int) []int {

	// make a array with product from front to back
	front := make([]int, len(nums))
	product := 1
	for i, num := range nums {
		product = product * num
		front[i] = product
	}

	// make a array with product from back to front
	back := make([]int, len(nums))
	product = 1
	for j := len(nums) - 1; j >= 0; j-- {
		fmt.Println(j, nums[j])
		product = product * nums[j]
		back[j] = product
	}

	result := make([]int, len(nums))

	for i := range nums {
		productFront := 1
		if i-1 >= 0 {
			productFront = front[i-1]
		}
		productBack := 1
		if i+1 < len(nums) {
			productBack = back[i+1]
		}

		product := productFront * productBack
		result[i] = product
	}

	return result
}
