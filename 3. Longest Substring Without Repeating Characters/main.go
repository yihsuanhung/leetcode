package main

import "fmt"

func lengthOfLongestSubstring(s string) int {

	l, r := 0, 0
	idxMap := map[string]int{}
	maxWindow := 0

	for r < len(s) {

		letter := string(s[r])

		if idx, ok := idxMap[letter]; ok {
			if idx+1 > l {
				l = idx + 1
			}
		}

		idxMap[letter] = r

		if (r - l + 1) > maxWindow {
			maxWindow = r - l + 1
		}
		r++
	}

	return maxWindow
}

func main() {
	s := "pwwkew"
	answer := lengthOfLongestSubstring(s)
	fmt.Println(answer)
}
