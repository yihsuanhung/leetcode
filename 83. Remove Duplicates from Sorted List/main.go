package main

func mian() {}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {

	if head == nil {
		return head

	}

	current := head

	for current.Next != nil {

		if current.Next.Val == current.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}

	return head

}

// func reverseList(head *ListNode) *ListNode {

// 	var previous *ListNode

// 	current := head

// 	for current != nil {
// 		next := current.Next
// 		current.Next = previous
// 		previous = current
// 		current = next
// 	}

// 	return previous

// }
