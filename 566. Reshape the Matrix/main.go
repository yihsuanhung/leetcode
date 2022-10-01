package main

import "fmt"

func main() {

	layer1 := []int{1, 2}
	layer2 := []int{3, 4}
	mat := [][]int{layer1, layer2}

	ans := matrixReshape(mat, 2, 4)

	fmt.Println("answer", ans)

}

func matrixReshape(mat [][]int, r int, c int) [][]int {

	if r*c != len(mat)*len(mat[0]) {
		return mat
	}

	// 把mat壓平
	flat := []int{}

	for _, layer := range mat {

		flat = append(flat, layer...)

	}

	// 組裝成新的mat

	result := [][]int{}

	for r > 0 {

		layer := flat[:c]
		flat = flat[c:]

		result = append(result, layer)
		r--

	}

	return result

}
