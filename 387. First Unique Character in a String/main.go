package main

import "fmt"

func main() {

	s := "dddccdbba"

	answer := firstUniqChar(s)

	fmt.Println(answer)

}

func firstUniqChar(s string) int {

	count := map[string]int{}

	for i := range s {

		current := string(s[i])

		if _, ok := count[current]; ok {
			count[current]++
		} else {
			count[current] = 1
		}

	}

	// for key, value := range count {
	// 	if value == 1 {
	// 		return key
	// 	}
	// }

	fmt.Println(count)

	// for i := len(s) - 1; i >= 0; i-- {

	// 	current := string(s[i])

	// 	if _, ok := count[current]; ok {

	// 		count[current]++

	// 		unique = -1

	// 	} else {

	// 		count[current] = 1

	// 		unique = i
	// 	}

	// }

	// if c, ok := count[string(s[unique])]; ok {
	// 	if c > 1 {
	// 		return -1
	// 	}
	// }

	return 1

}
