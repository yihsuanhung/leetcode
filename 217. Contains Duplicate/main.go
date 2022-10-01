package main

import "fmt"

func containsDuplicate(nums []int) bool {

	dupMap := map[int]bool{}

	for _, num := range nums {

		if _, ok := dupMap[num]; ok {
			return true
		} else {
			dupMap[num] = true
		}

	}

	return false

}

func main() {

	nums := []int{1, 2, 3, 3, 6}

	answer := containsDuplicate(nums)

	fmt.Print(answer)

}
