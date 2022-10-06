package main

func main() {}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {

	// 使用 dummy 是為了好處理，不用考慮一些edge case，例如斷頭等等
	dummy := new(ListNode)
	dummy.Next = head

	prevLeft := dummy
	current := head

	// 把 curr 走到 left 上， 同時 prevLeft 走到 left 的前一個
	for i := 1; i < left; i++ {
		current = current.Next
		prevLeft = prevLeft.Next
	}

	// 把 prevLeft 停留在開始 reverse 之前，使用一個新的 prev 當作 pointer
	// 迴圈走 r - l + 1 遍，因為等一下會動到curr.Next，所以先用next複製起來
	// 每次把 curr 指向 prev，最後 curr 與 prev 各前走一格
	prev := new(ListNode)
	for j := 0; j < right-left+1; j++ {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	// 更新指針，prevLeft 的下下個變成curr，下一個變成prev
	prevLeft.Next.Next = current
	prevLeft.Next = prev

	return dummy.Next
}
