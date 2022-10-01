package main

import "fmt"

func main() {

	answer := generate(6)
	fmt.Println(answer)

}

func generate(numRows int) [][]int {

	var t [][]int

	t = append(t, []int{1})
	if numRows == 1 {
		return t
	}

	t = append(t, []int{1, 1})
	if numRows == 2 {
		return t
	}

	for i := 2; i < numRows; i++ {

		var layer []int

		layer = append(layer, 1)

		previousLayer := t[i-1]

		for j := 0; j < i-1; j++ {
			layer = append(layer, previousLayer[j]+previousLayer[j+1])
		}

		layer = append(layer, 1)

		t = append(t, layer)

	}

	return t

}
