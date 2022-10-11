package main

import (
	"container/heap"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	list1 := &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}
	list2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}
	list3 := &ListNode{Val: 2, Next: &ListNode{Val: 6}}
	lists := []*ListNode{list1, list2, list3}

	// list1 := &ListNode{}
	// lists := []*ListNode{list1}

	fmt.Println(lists[0])

	result := mergeKLists(lists)

	r := result
	for r != nil {
		fmt.Println(r.Val)
		r = r.Next
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

func mergeKLists(lists []*ListNode) *ListNode {

	h := MinHeap([]int{})
	heap.Init(&h)

	for _, n := range lists {
		for n != nil {
			heap.Push(&h, n.Val)
			n = n.Next
		}
	}

	result := new(ListNode)
	current := result

	for h.Len() > 0 {

		node := &ListNode{Val: heap.Pop(&h).(int)}
		current.Next = node
		current = current.Next
	}

	return result.Next
}
