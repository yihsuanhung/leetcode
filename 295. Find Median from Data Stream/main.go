package main

import (
	"container/heap"
	"fmt"
)

func main() {
	obj := Constructor()
	obj.AddNum(1)
	obj.AddNum(2)
	m1 := obj.FindMedian()
	obj.AddNum(3)
	m2 := obj.FindMedian()

	fmt.Println(m1, m2)
}

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(item interface{}) {
	*h = append(*h, item.(int))
}
func (h *MinHeap) Pop() interface{} {
	item := (*h)[len(*h)-1]
	(*h) = (*h)[:len(*h)-1]
	return item
}

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(item interface{}) {
	*h = append(*h, item.(int))
}
func (h *MaxHeap) Pop() interface{} {
	item := (*h)[len(*h)-1]
	(*h) = (*h)[:len(*h)-1]
	return item
}

type MedianFinder struct {
	larger  *MinHeap
	smaller *MaxHeap
}

func Constructor() MedianFinder {
	larger := &MinHeap{}
	smaller := &MaxHeap{}
	mediumFinter := MedianFinder{larger: larger, smaller: smaller}
	return mediumFinter
}

func (this *MedianFinder) AddNum(num int) {
	heap.Push(this.smaller, num)

	if (*this.smaller).Len() > 0 && (*this.larger).Len() > 0 && (*this.smaller)[0] > (*this.larger)[0] {
		top := heap.Pop(this.smaller).(int)
		heap.Push(this.larger, top)
	}

	if (*this.smaller).Len()-(*this.larger).Len() > 1 {
		top := heap.Pop(this.smaller).(int)
		heap.Push(this.larger, top)
	}

	if (*this.larger).Len()-(*this.smaller).Len() > 1 {
		top := heap.Pop(this.larger).(int)
		heap.Push(this.smaller, top)
	}

}

func (this *MedianFinder) FindMedian() float64 {

	if (*this.smaller).Len() > (*this.larger).Len() {
		return float64((*this.smaller)[0])
	}
	if (*this.larger).Len() > (*this.smaller).Len() {
		return float64((*this.larger)[0])
	}
	return float64((*this.larger)[0]+(*this.smaller)[0]) / 2
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
