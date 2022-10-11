package main

import (
	"fmt"
	"strings"
)

func main() {
	// path := "/a/./b/../../c/"
	path := "/../"
	fmt.Println("result", simplifyPath(path))
}

func simplifyPath(path string) string {

	paths := strings.Split(path, "/")
	stack := []string{}

	for _, p := range paths {
		if len(stack) > 0 && p == ".." {
			// pop
			stack = stack[:len(stack)-1]
		} else if p != "." && p != "" && p != ".." {
			stack = append(stack, p)
		}
	}
	result := "/"

	if len(stack) == 0 {
		return result
	}

	result += strings.Join(stack, "/")

	return result
}
