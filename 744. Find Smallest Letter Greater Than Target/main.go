package main

import "fmt"

func nextGreatestLetter(letters []byte, target byte) byte {
	l, r, c := -1, len(letters), 1

	for (l + 1) != r {

		m := (l + r) / 2

		fmt.Println(c, "左邊界", l, "右邊界", r, "中線", m, "值", string(letters[m]))

		if letters[m] <= target {
			l = m
		} else {
			r = m
		}

		c++
	}

	fmt.Println("結束，回傳", r)

	if r < 0 {
		return letters[0]
	}

	find := letters[r%len(letters)]

	return find

}

func main() {
	letterString := "cfj"
	targetString := "d"
	letters := []byte(letterString)
	target := []byte(targetString)[0]
	ngl := nextGreatestLetter(letters, target)
	fmt.Println(string(ngl))
}
