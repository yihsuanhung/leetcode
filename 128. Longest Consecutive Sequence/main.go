package main

import "fmt"

func main() {
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
	fmt.Println(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
}

func longestConsecutive(nums []int) int {
	max := 0
	numsMap := map[int]bool{}

	for _, num := range nums {
		numsMap[num] = true
	}

	for _, num := range nums {
		localMax := 0

		if _, ok := numsMap[num-1]; !ok {
			localMax = 1

			for _, ok := numsMap[num+1]; ok; _, ok = numsMap[num+1] {
				localMax++
				num++
			}

		}

		if localMax > max {
			max = localMax
		}
	}

	return max

}
