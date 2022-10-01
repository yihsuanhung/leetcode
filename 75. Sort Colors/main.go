package main

import "fmt"

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
}

func sortColors(nums []int) {

	colorsMap := map[int]int{}
	ans := []int{}

	for _, n := range nums {

		if _, ok := colorsMap[n]; ok {
			// 如果有在map內

			colorsMap[n]++

		} else {
			// 如果沒在map內

			colorsMap[n] = 1
		}

	}

	for i := 0; i < colorsMap[0]; i++ {
		ans = append(ans, 0)
	}

	for i := 0; i < colorsMap[1]; i++ {
		ans = append(ans, 1)
	}
	for i := 0; i < colorsMap[2]; i++ {
		ans = append(ans, 2)
	}
	fmt.Println(ans)

}
