package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(numSquares(13))
}

func numSquares(n int) int {

	sqrtFloor := int(math.Floor(math.Sqrt(float64(n))))

	perfectSquares := make([]int, sqrtFloor)

	for i := range perfectSquares {
		perfectSquares[i] = int(math.Pow(float64((i + 1)), 2))
	}

	memo := map[int][]int{}

	combination := findSumn(n, perfectSquares, memo)

	fmt.Println(combination)

	return len(combination)
}

func findSumn(n int, perfectSquares []int, memo map[int][]int) []int {

	// handle base cases
	if n == 0 {
		return []int{}
	}
	if n < 0 {
		return nil
	}
	if m, ok := memo[n]; ok {
		return m
	}

	var minCombination []int

	// recursion
	for _, ps := range perfectSquares {
		newTarget := n - ps
		combination := findSumn(newTarget, perfectSquares, memo)
		if combination != nil {
			newCombination := append(combination, ps)
			if minCombination == nil || len(newCombination) < len(minCombination) {
				minCombination = newCombination
				memo[n] = newCombination
			}
		}

	}

	return minCombination
}
