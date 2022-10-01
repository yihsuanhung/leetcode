package main

import "fmt"

func main() {
	s := "abbaca"
	fmt.Println(removeDuplicates(s))
}

// stack
// func removeDuplicates(s string) string {
// 	stack := []rune{}

// 	for _, b := range s {
// 		if len(stack) > 0 && stack[len(stack)-1] == b {
// 			stack = stack[:len(stack)-1]
// 			continue
// 		}
// 		stack = append(stack, b)

// 	}

// 	fmt.Println(stack)

// 	return string(stack)

// }

// two pointer

func removeDuplicates(s string) string {

	i, j := -1, 0

	bs := []byte(s)

	for j < len(bs) {
		if i == -1 || bs[i] != bs[j] {
			i++
			bs[i] = bs[j]
		} else {
			i--
		}

		j++

	}

	return string(bs[:i+1])

}
