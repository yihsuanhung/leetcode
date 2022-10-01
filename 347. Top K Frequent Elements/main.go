package main

import (
	"container/heap"
	"fmt"
)

type MaxHeap [][]int

func (h MaxHeap) Len() int {
	return len(h)
}
func (h MaxHeap) Less(i, j int) bool {
	return h[i][1] > h[j][1]
}
func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *MaxHeap) Push(item interface{}) {
	*h = append(*h, item.([]int))
}
func (h *MaxHeap) Pop() interface{} {
	item := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return item
}

func topKFrequent(nums []int, k int) []int {

	freqMap := map[int]int{}

	for _, num := range nums {
		freqMap[num]++
	}

	// freqMap looks like
	// {
	// 	1: 3
	// }

	h := new(MaxHeap)
	heap.Init(h)

	for num, freq := range freqMap {
		node := []int{num, freq}
		heap.Push(h, node)
	}

	answer := []int{}
	for i := 0; i < k; i++ {
		node := heap.Pop(h)
		answer = append(answer, node.([]int)[0])
	}

	return answer
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	result := topKFrequent(nums, k)
	fmt.Println(result)
}
