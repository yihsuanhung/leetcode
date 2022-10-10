package main

import "fmt"

func main() {
	nums1, nums2 := []int{1, 3, 5, 2, 4}, []int{6, 5, 4, 3, 2, 1, 7}
	fmt.Println("RESULT:", nextGreaterElement(nums1, nums2))
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {

	// init the map
	idxMap := map[int]int{}
	for i, num := range nums1 {
		idxMap[num] = i
	}

	// init the result array fill with -1
	result := make([]int, len(nums1))
	for i := range result {
		result[i] = -1
	}

	// init a stack
	stack := []int{}

	for _, num := range nums2 {

		for len(stack) > 0 && num > stack[len(stack)-1] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if topIdx, ok := idxMap[top]; ok {
				result[topIdx] = num
			}
		}

		stack = append(stack, num)
	}

	return result
}
