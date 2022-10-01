package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {

	if head == nil {
		return head
	}

	dummy := new(ListNode)
	p := dummy
	current := head

	for current != nil {

		if current.Val != val {
			p.Next = current
		} else {
			p.Next = current.Next
		}

		current = current.Next

	}

	return dummy.Next

}

func main() {}
