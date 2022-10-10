package main

import "fmt"

func main() {
	asteroids := []int{5, 10, -5}
	fmt.Println(asteroidCollision((asteroids)))

}

func asteroidCollision(asteroids []int) []int {

	stack := []int{}

	for _, a := range asteroids {
		survivor := a

		for len(stack) != 0 && (survivor < 0 && stack[len(stack)-1] > 0) && survivor != 0 {
			defender := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			fmt.Println("collecting...", defender, survivor)
			survivor = collision(defender, survivor)
		}
		if survivor != 0 {
			stack = append(stack, survivor)

		}

	}

	return stack
}

func collision(a, b int) int {

	fmt.Println("collision!", a, b)

	absA := a
	absB := b

	if a < 0 {
		absA = -a
	}

	if b < 0 {
		absB = -b
	}

	if absA > absB {
		return a
	}

	if absB > absA {
		return b
	}

	return 0
}
