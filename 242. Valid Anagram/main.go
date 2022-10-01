package main

import "fmt"

func main() {
	s := "ab"
	t := "a"
	answer := isAnagram(s, t)
	fmt.Println(answer)
}

func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	anagramDict := map[string]int{}

	for _, l := range s {

		letter := string(l)

		if _, ok := anagramDict[letter]; ok {
			anagramDict[letter]++
		} else {
			anagramDict[letter] = 1
		}

	}

	for _, l := range t {

		letter := string(l)

		if _, ok := anagramDict[letter]; ok && anagramDict[letter] > 0 {
			anagramDict[letter]--
		} else {
			return false
		}

	}

	return true
}
