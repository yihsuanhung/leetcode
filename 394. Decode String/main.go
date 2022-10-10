package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(decodeString("2[abc]3[cd]ef"))
}

// 2022/10/9 stack solution
func decodeString(s string) string {
	stack := []string{}

	for _, l := range s {

		current := string(l)

		if current == "]" {
			tmp := ""
			for len(stack) > 0 && stack[len(stack)-1] != "[" {
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1] // pop
				tmp = top + tmp

			}

			stack = stack[:len(stack)-1] // pop [

			k := ""
			for len(stack) > 0 && IsNumber(stack[len(stack)-1]) {
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1] // pop
				k = top + k
			}

			intK, _ := strconv.Atoi(k)

			newString := ""
			for i := 0; i < intK; i++ {
				newString += tmp
			}

			stack = append(stack, newString)

		} else {
			stack = append(stack, current)
		}

	}
	result := ""

	for i := 0; i < len(stack); i++ {
		result += stack[i]
	}

	return result

}

// 2022/8/10 recursive solution
// func IsLetter(s string) bool {
// 	for _, r := range s {
// 		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
// 			return false
// 		}
// 	}
// 	return true
// }

func IsNumber(s string) bool {
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}

// func decode(code string) string {

// 	fmt.Println(code)

// 	answer := ""
// 	k := 0

// 	for _, i := range code {
// 		current := string(i)

// 		if IsNumber(current) {
// 			kInt, _ := strconv.Atoi(current)
// 			k = kInt
// 			fmt.Println("NESTED", decode(code[2:]))
// 			answer += decode(code[2:])
// 		} else if IsLetter(current) {
// 			answer += current
// 		} else if current == "]" {
// 			for j := 0; j < k; j++ {
// 				answer += answer
// 			}
// 		}

// 	}
// 	return answer
// }
