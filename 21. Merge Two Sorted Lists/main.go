package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	}
	list2.Next = mergeTwoLists(list1, list1.Next)
	return list2

}

func mergeTwoListsLoop(list1 *ListNode, list2 *ListNode) *ListNode {

	sortedList := new(ListNode)
	placeholder := sortedList

	// 	迴圈，直到list1或list2其中一個到尾端
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			placeholder.Next = list1
			list1 = list1.Next
		} else {
			placeholder.Next = list2
			list2 = list2.Next
		}
		placeholder = placeholder.Next
	}

	// 迴圈完之後，把剩下的list接到合併後的list尾端
	if list1 != nil {
		placeholder.Next = list1
	}

	if list2 != nil {
		placeholder.Next = list2
	}

	return sortedList.Next
}

func main() {}
