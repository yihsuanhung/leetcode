class Solution:
    def oddEvenList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        # len = 0
        if head is None:
            return None
        # len = 1
        if head.next is None:
            return head
        # len >= 2
        odd = head
        even = head.next
        even_head = head.next
        while odd and even:
            if even.next is not None:
                odd.next = even.next
                odd = odd.next
            else:
                odd.next = even_head
                break
            if odd.next is not None:
                even.next = odd.next
                even = even.next
            else:
                even.next = None
                odd.next = even_head
                break
        return head
