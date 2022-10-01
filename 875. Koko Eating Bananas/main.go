package main

import (
	"fmt"
	"math"
)

func main() {
	piles := []int{312884470}
	// piles := []int{30, 11, 23, 4, 20}
	// h := 6
	h := 968709470
	fmt.Println(minEatingSpeed(piles, h))
}

func checkBoundary(m int, piles []int, h int) bool {
	sum := 0
	for _, p := range piles {

		sum += int(math.Ceil(float64(p) / float64(m)))
	}
	return sum > h
}

func findMax(arr []int) int {
	max := 0
	for _, i := range arr {
		if i > max {
			max = i
		}
	}
	return max
}

func minEatingSpeed(piles []int, h int) int {

	maxPile := findMax(piles)

	// left pointer should start from 0
	l, r := 0, maxPile+1

	for l+1 != r {
		m := (l + r) / 2

		if checkBoundary(m, piles, h) {
			l = m
		} else {
			r = m
		}
	}

	return r
}
