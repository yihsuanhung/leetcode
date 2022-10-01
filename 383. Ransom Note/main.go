package main

import "fmt"

func main() {
	ransomNote := "aa"
	magazine := "aab"
	answer := canConstruct(ransomNote, magazine)
	fmt.Println(answer)
}

func canConstruct(ransomNote string, magazine string) bool {

	letterDict := map[string]int{}

	for _, s := range magazine {

		letter := string(s)

		if _, ok := letterDict[letter]; ok {
			letterDict[letter]++
		} else {
			letterDict[letter] = 1
		}

	}

	for _, s := range ransomNote {

		letter := string(s)

		if _, ok := letterDict[letter]; ok && letterDict[letter] > 0 {
			letterDict[letter]--
		} else {
			return false
		}

	}

	return true
}
