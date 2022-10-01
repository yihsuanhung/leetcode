package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			tracker := head
			for tracker != slow {
				slow = slow.Next
				tracker = tracker.Next
			}
			return tracker
		}
	}

	return nil
}

func main() {}
