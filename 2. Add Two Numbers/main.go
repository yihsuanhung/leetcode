package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	lSum := new(ListNode)
	result := lSum
	cS := result
	c1 := l1
	c2 := l2
	carry := false

	for c1 != nil || c2 != nil {

		if c1 == nil {
			c1 = new(ListNode)
			c1.Val = 0
		}

		if c2 == nil {
			c2 = new(ListNode)
			c2.Val = 0
		}

		n1 := c1.Val
		n2 := c2.Val

		sum := n1 + n2
		if carry {
			sum += 1
		}
		if sum >= 10 {
			carry = true
		} else {
			carry = false
		}

		lUnit := new(ListNode)
		lUnit.Val = sum % 10
		cS.Next = lUnit

		c1 = c1.Next
		c2 = c2.Next
		cS = cS.Next

	}

	if carry {
		lLeading := new(ListNode)
		lLeading.Val = 1
		cS.Next = lLeading
	}

	return result.Next
}

func main() {
	// l1 := new(ListNode)
	// l1.Val = 2
	// l1.Next = new(ListNode)
	// l1.Next.Val = 4
	// l1.Next.Next = new(ListNode)
	// l1.Next.Next.Val = 3
	// l2 := new(ListNode)
	// l2.Val = 5
	// l2.Next = new(ListNode)
	// l2.Next.Val = 6
	// l2.Next.Next = new(ListNode)
	// l2.Next.Next.Val = 4

	l1 := new(ListNode)
	l1.Val = 0
	l2 := new(ListNode)
	l2.Val = 0

	// l1 := new(ListNode)
	// l1.Val = 9
	// l1.Next = new(ListNode)
	// l1.Next.Val = 9
	// l1.Next.Next = new(ListNode)
	// l1.Next.Next.Val = 9
	// l1.Next.Next.Next = new(ListNode)
	// l1.Next.Next.Next.Val = 9
	// l1.Next.Next.Next.Next = new(ListNode)
	// l1.Next.Next.Next.Next.Val = 9
	// l1.Next.Next.Next.Next.Next = new(ListNode)
	// l1.Next.Next.Next.Next.Next.Val = 9
	// l1.Next.Next.Next.Next.Next.Next = new(ListNode)
	// l1.Next.Next.Next.Next.Next.Next.Val = 9
	// l2 := new(ListNode)
	// l2.Val = 9
	// l2.Next = new(ListNode)
	// l2.Next.Val = 9
	// l2.Next.Next = new(ListNode)
	// l2.Next.Next.Val = 9
	// l2.Next.Next.Next = new(ListNode)
	// l2.Next.Next.Next.Val = 9

	result := addTwoNumbers(l1, l2)
	r := result
	for r != nil {
		fmt.Print(r.Val)
		r = r.Next
	}
}
