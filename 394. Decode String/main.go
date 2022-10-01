package main

import (
	"fmt"
	"strconv"
)

func IsLetter(s string) bool {
	for _, r := range s {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
			return false
		}
	}
	return true
}

func IsNumber(s string) bool {
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}

func decode(code string) string {

	fmt.Println(code)

	answer := ""
	k := 0

	for _, i := range code {
		current := string(i)

		if IsNumber(current) {
			kInt, _ := strconv.Atoi(current)
			k = kInt
			fmt.Println("NESTED", decode(code[2:]))
			answer += decode(code[2:])
		} else if IsLetter(current) {
			answer += current
		} else if current == "]" {
			for j := 0; j < k; j++ {
				answer += answer
			}
		}

	}
	return answer
}

func main() {

	answer := decode("2[a]")
	fmt.Println(answer)

}
