package main

import "fmt"

func main() {
	s := "))" //"{()]}"

	answer := isValid(s)
	fmt.Println(answer)
}

func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}

	bracketsMap := map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
	}
	stack := []string{}

	for _, s := range s {
		currentBracket := string(s)

		if closeBracket, ok := bracketsMap[currentBracket]; ok {
			stack = append(stack, closeBracket)
		} else {

			if len(stack) < 1 {
				return false
			}

			if currentBracket == stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}

	}

	if len(stack) != 0 {
		return false
	}

	return true
}

// var small = 0
// var mid = 0
// var big = 0
// for _, s := range s {

// 	switch string(s) {

// 	case "(":
// 		small += 1
// 	case ")":
// 		small -= 1
// 		if small < 0 {
// 			return false
// 		}
// 		if mid+big != 0 {
// 			return false
// 		}

// 	case "[":
// 		mid += 1
// 	case "]":
// 		mid -= 1
// 		if mid < 0 {
// 			return false
// 		}
// 		if small+big != 0 {
// 			return false
// 		}
// 	case "{":
// 		big += 1
// 	case "}":
// 		big -= 1
// 		if big < 0 {
// 			return false
// 		}
// 		if small+mid != 0 {
// 			return false
// 		}
// 	default:
// 		return true
// 	}
// 	for small+mid+big < 0 {
// 		return false
// 	}
// }
