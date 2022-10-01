package main

import (
	"fmt"
	"strings"
)

func removeDuplicateLetters(s string) string {

	lastIdx := map[string]int{}
	available := map[string]bool{}
	stack := []string{}

	// 製造兩個map
	for i, l := range s {
		letter := string(l)
		lastIdx[letter] = i
		if _, ok := available[letter]; !ok {
			available[letter] = true
		}
	}

	for j, l := range s {

		letter := string(l)

		if available[letter] {

			for len(stack) > 0 && stack[len(stack)-1] > letter && lastIdx[stack[len(stack)-1]] > j {
				// pop & make top item available
				available[stack[len(stack)-1]] = true
				stack = stack[:len(stack)-1]
			}

			// push & make letter unavailable
			stack = append(stack, letter)
			available[letter] = false

		}

	}

	return strings.Join(stack, "")
}

func main() {
	s := "cbacdcbc"
	answer := removeDuplicateLetters(s)
	fmt.Println("answer", answer)
}

// if len(stack) == 0 {
// 	// TODO: push, true->false, continue
// 	stack = append(stack, letter)
// 	available[letter] = false
// 	continue
// }

// top := stack[len(stack)-1]

// if top <= letter {
// 	// TODO: push, true->false, continue
// 	stack = append(stack, letter)
// 	available[letter] = false
// 	continue
// }

// for top > letter {

// 	if lastIdx[top] > j {
// 		// * 後面還有出現
// 		// TODO: pop, fasle->true
// 		fmt.Println(stack)
// 		stack = stack[:len(stack)-2]
// 		available[letter] = true
// 	} else {
// 		// * 後面沒有出現
// 		// TODO: push, true->false, continue
// 		stack = append(stack, letter)
// 		available[letter] = false
// 		continue
// 	}

// }
