package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb")) // 3
	fmt.Println(lengthOfLongestSubstring("bbbbb"))    // 1
	fmt.Println(lengthOfLongestSubstring("pwwkew"))   // 3
	fmt.Println(lengthOfLongestSubstring("abba"))     // 2
}

func lengthOfLongestSubstring(s string) int {
	l := 0
	charMap := map[string]int{} // {letter:index}
	result := 0

	for r := 0; r < len(s); r++ {
		currentChar := string(s[r])

		if idx, ok := charMap[currentChar]; ok {
			// update l if l is behind idx
			if idx+1 > l {
				l = idx + 1
			}
		}

		charMap[currentChar] = r

		if r-l+1 > result {
			result = r - l + 1
		}

	}
	return result
}

// func lengthOfLongestSubstring(s string) int {

// 	l, r := 0, 0
// 	idxMap := map[string]int{}
// 	maxWindow := 0

// 	for r < len(s) {

// 		letter := string(s[r])

// 		if idx, ok := idxMap[letter]; ok {
// 			if idx+1 > l {
// 				l = idx + 1
// 			}
// 		}

// 		idxMap[letter] = r

// 		if (r - l + 1) > maxWindow {
// 			maxWindow = r - l + 1
// 		}
// 		r++
// 	}

// 	return maxWindow
// }

// func main() {
// 	s := "pwwkew"
// 	answer := lengthOfLongestSubstring(s)
// 	fmt.Println(answer)
// }
