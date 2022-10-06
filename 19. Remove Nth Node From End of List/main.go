package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	// head := &ListNode{Val: 1}
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2}}
	n := 1
	result := removeNthFromEnd(head, n)
	r := result
	for r != nil {
		fmt.Print(r.Val)
		r = r.Next
	}
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	dummy := new(ListNode)
	dummy.Next = head
	slow := dummy
	primer := dummy.Next.Next

	for n > 1 {
		primer = primer.Next
		n--
	}

	for primer != nil {
		slow = slow.Next
		primer = primer.Next
	}

	slow.Next = slow.Next.Next

	return dummy.Next
}
