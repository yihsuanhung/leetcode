package main

func main() {

}

func reverseString(s []byte) {
	i := 0
	j := len(s) - 1

	for i < j {
		left := s[i]
		s[i] = s[j]
		s[j] = left

		i++
		j--
	}
}
