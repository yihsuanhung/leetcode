package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 2, 3, 4, 5}
	nums2 := []int{2, 3, 2, 4, 5, 6}
	ans := intersect(nums1, nums2)

	fmt.Println(ans)
}

func intersect(nums1 []int, nums2 []int) []int {

	numMap := map[int]int{}
	answer := []int{}

	for _, num := range nums1 {

		if _, ok := numMap[num]; ok {
			numMap[num]++
		} else {
			numMap[num] = 1
		}
	}

	for _, num := range nums2 {

		if count, ok := numMap[num]; ok {

			if count > 0 {
				numMap[num]--
				answer = append(answer, num)
			}

		}

	}

	return answer

}
