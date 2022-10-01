package main

func main() {

	nums := []int{1, 2, 3, 4}
	getConcatenation(nums)
}

func getConcatenation(nums []int) []int {

	a := append(nums, nums...)

	return a

}
