package tpt

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if n < 0 || head == nil {
		return head
	}
	fp, sp := head, head

	// move n step forward
	for i := 0; i < n; i++ {
		if fp == nil {
			return head
		}
		fp = fp.Next
	}

	// keep moving, slow pointer
	// will at Nth from end of list
	pre := sp
	for fp != nil {
		fp = fp.Next
		pre = sp
		sp = sp.Next
	}

	// delete sp
	if sp == head {
		head = sp.Next
	} else {
		pre.Next = sp.Next
	}

	return head
}
