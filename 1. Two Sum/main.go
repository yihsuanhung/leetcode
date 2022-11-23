package main

import "fmt"

func main() {
	// fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
}

func twoSum(nums []int, target int) []int {

	// make a dict to store difference bwtween num and target as key, index as value
	// diff:index
	dict := make(map[int]int)

	for i, num := range nums {

		diff := target - num

		if diffIdx, ok := dict[num]; ok {
			return []int{i, diffIdx}
		} else {
			dict[diff] = i
		}

	}

	return []int{}
}
