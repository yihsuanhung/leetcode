package main

import (
	"container/heap"
	"fmt"
)

func main() {
	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k := 4
	result := findKthLargest(nums, k)
	fmt.Println(result)
}

// Solution: Max heap -- O(n+klogn)

type MaxHeap []int

func (h MaxHeap) Len() int               { return len(h) }
func (h MaxHeap) Less(i, j int) bool     { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)          { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(item interface{}) { *h = append(*h, item.(int)) }
func (h *MaxHeap) Pop() interface{} {
	item := (*h)[len(*h)-1]
	(*h) = (*h)[:len(*h)-1]
	return item
}

// func findKthLargest(nums []int, k int) int {

// 	h := MaxHeap(nums)
// 	heap.Init(&h)

// 	res := 0

// 	for i := 0; i < k; i++ {
// 		res = heap.Pop(&h).(int)
// 	}

// 	return res

// }

// Solution: Min heap -- O(nlogk)
type MinHeap []int

func (h MinHeap) Len() int               { return len(h) }
func (h MinHeap) Less(i, j int) bool     { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)          { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(item interface{}) { *h = append(*h, item.(int)) }
func (h *MinHeap) Pop() interface{} {
	item := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return item
}

func findKthLargest(nums []int, k int) int {

	h := MinHeap([]int{})
	// heapify
	heap.Init(&h)

	for _, n := range nums {
		if len(h) < k {
			heap.Push(&h, n)
		} else if n >= h[0] {
			heap.Pop(&h)
			heap.Push(&h, n)
		}
	}

	reuslt := heap.Pop(&h).(int)

	return reuslt

}
