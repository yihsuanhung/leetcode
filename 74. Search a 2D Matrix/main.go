package main

import "fmt"

func main() {
	layer1 := []int{1, 3, 5, 7}
	layer2 := []int{10, 11, 16, 20}
	layer3 := []int{23, 30, 34, 60}
	matrix := [][]int{layer1, layer2, layer3}
	ans := searchMatrix(matrix, 13)
	fmt.Println(ans)
	fmt.Println("==========bs==========")
	layer4 := []int{1}
	matrix2 := [][]int{layer4}
	r := searchMatrixBinary(matrix2, 2)
	fmt.Println(r)
}

func searchMatrix(matrix [][]int, target int) bool {

	for i := 0; i < len(matrix); i++ {

		for j := 0; j < len(matrix[0]); j++ {

			if matrix[i][j] == target {
				return true
			}

		}
	}

	return false
}

func searchMatrixBinary(matrix [][]int, target int) bool {

	l, r := -1, len(matrix)

	for (l + 1) != r {
		m := (l + r) / 2

		if matrix[m][len(matrix[m])-1] <= target {
			l = m
		} else {
			r = m
		}
	}

	// if r == len(matrix) {
	// 	r = len(matrix) - 1
	// }

	fmt.Println(r)
	// return false

	return BinarySearch(matrix[r], target)

}

func BinarySearch(numbers []int, target int) bool {
	l, r := -1, len(numbers)
	for (l + 1) != r {
		m := (l + r) / 2
		if numbers[m] <= target {
			l = m
		} else {
			r = m
		}
	}

	// if r == len(numbers) {
	// 	r = len(numbers) - 1
	// }

	if numbers[r] == target {
		return true
	}

	return false
}
