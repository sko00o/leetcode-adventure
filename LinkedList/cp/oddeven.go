package cp

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	odd, even := head, head.Next

	for odd != nil && even != nil && even.Next != nil {
		evenH := odd.Next
		odd.Next = even.Next
		odd = odd.Next // odd move forward

		even.Next = odd.Next
		even = even.Next // even move forward
		odd.Next = evenH
	}

	return head
}
