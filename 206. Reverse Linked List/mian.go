package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	rev := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return rev
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
