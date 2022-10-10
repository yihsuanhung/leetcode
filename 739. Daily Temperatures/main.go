package main

import "fmt"

func main() {
	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperatures(temperatures))

}

func dailyTemperatures(T []int) []int {
	stack := []int{}
	res := make([]int, len(T))
	for i := 0; i < len(T); i++ {
		for len(stack) != 0 && T[stack[len(stack)-1]] < T[i] {
			res[stack[len(stack)-1]] = i - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return res
}
