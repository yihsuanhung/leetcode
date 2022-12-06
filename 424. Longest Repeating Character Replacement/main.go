package main

import "fmt"

func main() {
	fmt.Println(characterReplacement("ABAB", 2))
	fmt.Println(characterReplacement("AABABBA", 1))
}

func characterReplacement(s string, k int) int {
	l := 0
	count := map[string]int{}
	result := 0

	for r := 0; r < len(s); r++ {
		// add current string into map
		count[string(s[r])]++

		// find the most frequent charactor's count
		mostFeqChart := 0
		for _, value := range count {
			if value > mostFeqChart {
				mostFeqChart = value
			}
		}

		// check is window valid, if not, slide l
		for r-l+1-mostFeqChart > k {
			count[string(s[l])]--
			l++
		}

		// update result if current result is bigger than max result
		if r-l+1 > result {
			result = r - l + 1
		}
	}

	return result
}
