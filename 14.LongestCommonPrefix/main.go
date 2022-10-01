package main

import (
	"fmt"
	"sort"
	"strings"
)

func longestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}

	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i]) < len(strs[j])
	})

	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
		// while loop
		for !strings.HasPrefix(strs[i], prefix) {

			prefix = prefix[0 : len(prefix)-1]

			if len(prefix) == 0 {
				return ""
			}

		}

	}

	return prefix
}

func main() {
	strs := []string{"jhtml", "jcss", "js", "java"}
	prefix := longestCommonPrefix(strs)
	fmt.Print("prefix is ", prefix)
}
