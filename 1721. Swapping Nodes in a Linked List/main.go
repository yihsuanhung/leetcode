package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// result := swapNodes(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}, 2)
	result := swapNodes(&ListNode{Val: 7, Next: &ListNode{Val: 9, Next: &ListNode{Val: 6, Next: &ListNode{Val: 6, Next: &ListNode{Val: 7, Next: &ListNode{Val: 8, Next: &ListNode{Val: 3, Next: &ListNode{Val: 0, Next: &ListNode{Val: 9, Next: &ListNode{Val: 5}}}}}}}}}}, 5)
	r := result
	for r != nil {
		fmt.Print(r.Val)
		r = r.Next
	}

}

func swapNodes(head *ListNode, k int) *ListNode {

	dummy := new(ListNode)
	dummy.Next = head

	left := dummy // starts with dummy, go k steps, and "left" will stop on the kth node
	for i := k; i > 0; i-- {
		left = left.Next
	}

	primer := left
	right := dummy // starts with dummy, go any steps until primer is nil, and "right" will stop on the kth node form end of list
	for primer != nil {
		right = right.Next
		primer = primer.Next
	}

	left.Val, right.Val = right.Val, left.Val

	return dummy.Next
}
