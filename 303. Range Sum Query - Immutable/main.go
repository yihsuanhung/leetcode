package main

import "fmt"

func main() {
	obj := Constructor([]int{-2, 0, 3, -5, 2, -1})
	fmt.Println(obj.SumRange(0, 2)) // 1
	fmt.Println(obj.SumRange(2, 5)) // -1
	fmt.Println(obj.SumRange(0, 5)) // -3
}

type NumArray struct {
	nums []int
	memo []int
}

func Constructor(nums []int) NumArray {
	memo := make([]int, len(nums))
	sum := 0

	for i, num := range nums {
		sum += num
		memo[i] = sum
	}

	return NumArray{nums, memo}
}

func (this *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return this.memo[right]
	}
	return this.memo[right] - this.memo[left-1]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
