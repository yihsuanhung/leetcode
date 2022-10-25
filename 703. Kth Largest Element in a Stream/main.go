package main

import (
	"container/heap"
	"fmt"
)

func main() {
	k := 3
	nums := []int{4, 5, 8, 2}
	obj := Constructor(k, nums)

	ps := []int{3, 5, 10, 9, 4}

	for _, p := range ps {
		fmt.Println("push", p, "return", obj.Add(p))
	}

}

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(item interface{}) {
	*h = append(*h, item.(int))
}
func (h *MinHeap) Pop() interface{} {
	top := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return top
}

type KthLargest struct {
	heap *MinHeap
	k    int
}

func Constructor(k int, nums []int) KthLargest {
	h := &MinHeap{}
	for _, num := range nums {
		if h.Len() < k {
			heap.Push(h, num)
		} else if num > (*h)[0] {
			heap.Pop(h)
			heap.Push(h, num)
		}
	}

	return KthLargest{h, k}
}

func (this *KthLargest) Add(val int) int {

	heap.Push(this.heap, val)

	for i := 0; i < (*this.heap).Len()-this.k; i++ {
		heap.Pop(this.heap)
	}

	return (*this.heap)[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
