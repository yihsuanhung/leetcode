package main

import (
	"container/heap"
	"fmt"
)

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(item interface{}) { *h = append(*h, item.(int)) }
func (h *MaxHeap) Pop() interface{} {
	item := (*h)[len(*h)-1]
	(*h) = (*h)[:len(*h)-1]
	return item
}

func findKthLargest(nums []int, k int) int {

	h := MaxHeap(nums)
	heap.Init(&h)

	res := 0

	for i := 0; i < k; i++ {
		res = heap.Pop(&h).(int)
	}

	return res

}

func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 2
	result := findKthLargest(nums, k)
	fmt.Println(result)
}
