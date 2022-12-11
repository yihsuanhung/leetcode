package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// result := removeNthFromEnd(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}, 3)
	// result := removeNthFromEnd(&ListNode{Val: 1}, 1)
	result := removeNthFromEnd(&ListNode{Val: 1, Next: &ListNode{Val: 2}}, 1)
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
	fast := head

	for n > 0 {
		fast = fast.Next
		n--
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

	return dummy.Next
}

// func removeNthFromEnd(head *ListNode, n int) *ListNode {

// 	dummy := new(ListNode)
// 	dummy.Next = head
// 	slow := dummy
// 	primer := dummy.Next.Next

// 	for n > 1 {
// 		primer = primer.Next
// 		n--
// 	}

// 	for primer != nil {
// 		slow = slow.Next
// 		primer = primer.Next
// 	}

// 	slow.Next = slow.Next.Next

// 	return dummy.Next
// }
