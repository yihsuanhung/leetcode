package main

import "fmt"

func main() {

	fmt.Println(nextGreaterElements([]int{1, 2, 3, 4, 3}))
}

func nextGreaterElements(nums []int) []int {

	nums2 := append(nums, nums...)

	result := make([]int, len(nums))

	for i := range result {
		result[i] = -1
	}

	stack := []int{}

	for i, num := range nums2 {

		for len(stack) > 0 && num > nums2[stack[len(stack)-1]] {

			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[top%len(nums)] = num

		}

		stack = append(stack, i)

	}

	return result

}
