package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}

func groupAnagrams(strs []string) [][]string {

	dict := make(map[string][]string)

	for _, str := range strs {
		// sort string by making a new byte slice copy
		sorted := make([]byte, len([]byte(str)))
		copy(sorted, []byte(str))
		sort.Slice(sorted, func(i, j int) bool {
			return sorted[i] < sorted[j]
		})

		// check is sorted string in map
		if group, ok := dict[string(sorted)]; ok {
			dict[string(sorted)] = append(group, str)
		} else {
			dict[string(sorted)] = []string{str}
		}
	}

	result := [][]string{}
	for _, value := range dict {
		result = append(result, value)
	}

	return result

}
