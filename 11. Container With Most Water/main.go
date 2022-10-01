package main

import "fmt"

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	area := maxArea(height)
	fmt.Println(area)
}

func maxArea(height []int) int {

	mArea := 0

	i := 0
	j := len(height) - 1

	for i != j {

		h := min(height[i], height[j])
		w := j - i
		area := h * w

		if area > mArea {
			mArea = area
		}

		if height[i] > height[j] {
			j--
		} else {
			i++
		}

	}

	return mArea
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// func maxArea(height []int) int {

// 	maxArea := 0

// 	for i := 0; i < len(height); i++ {
// 		for j := i + 1; j < len(height); j++ {
// 			w := j - 1
// 			h := min(height[i], height[j])
// 			area := w * h
// 			if area > maxArea {
// 				maxArea = area
// 			}
// 		}
// 	}
// 	return maxArea
// }
