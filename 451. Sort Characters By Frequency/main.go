package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(frequencySort("tree"))
	fmt.Println(frequencySort("cccaaa"))
	fmt.Println(frequencySort("Aabb"))
}

type Pair struct {
	char string
	freq int
}
type MaxHeap []Pair

func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Less(i, j int) bool {
	// h => [ {char:t, freq:2}, {char:r, freq:1} ]
	return h[i].freq > h[j].freq
}
func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *MaxHeap) Push(item interface{}) {
	*h = append(*h, item.(Pair))
}
func (h *MaxHeap) Pop() interface{} {
	item := (*h)[len(*h)-1]
	(*h) = (*h)[:len(*h)-1]
	return item
}

func frequencySort(s string) string {

	// setp 1
	freqMap := map[string]int{}
	// "sting" -> "s" type != string, type == []byte
	for _, c := range s {
		char := string(c)
		freqMap[char]++
	}

	// setp 2
	h := new(MaxHeap)
	heap.Init(h)

	for key, value := range freqMap {
		heap.Push(h, Pair{char: key, freq: value})
	}

	// setp 3
	result := ""
	for h.Len() > 0 {
		top := heap.Pop(h).(Pair)
		char, freq := top.char, top.freq
		for freq > 0 {
			result += char
			freq--
		}

	}

	return result
}
