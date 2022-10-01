package main

import "fmt"

func main() {
	arr := []int{3, 5, 3, 2, 0}
	fmt.Println(peakIndexInMountainArray(arr))
}

func peakIndexInMountainArray(arr []int) int {

	if len(arr) == 3 {
		return 1
	}

	l, r := 0, len(arr)-1

	for l+1 != r {

		m := (l + r) / 2

		if arr[m] > arr[m+1] {
			r = m
		} else {
			l = m
		}

	}

	return r
}
