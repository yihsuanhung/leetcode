package main

import (
	"container/heap"
	"fmt"
)

func main() {
	s := "aacvcaab"
	fmt.Println(reorganizeString(s))
}

type MaxHeap []Element

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].count > h[j].count }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(item interface{}) {
	*h = append(*h, item.(Element))
}
func (h *MaxHeap) Pop() interface{} {
	item := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return item
}

type Element struct {
	letter string
	count  int
}

func reorganizeString(s string) string {
	// make a cont map
	frequency := map[string]int{}
	for _, l := range s {
		frequency[string(l)]++
	}

	h := &MaxHeap{}

	// push all letters in the heap
	for letter, count := range frequency {
		element := Element{letter, count}
		heap.Push(h, element)
	}

	var result string
	var prevLetter string

	for h.Len() > 0 {
		top := heap.Pop(h).(Element)

		if top.letter == prevLetter {
			if h.Len() == 0 {
				return ""
			}
			hold := top
			top = heap.Pop(h).(Element)
			heap.Push(h, hold)
		}

		result += top.letter
		prevLetter = top.letter
		top.count--

		if top.count > 0 {
			heap.Push(h, top)
		}
	}

	return result

}
