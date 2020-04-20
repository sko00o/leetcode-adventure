package cp

// Merge two sorted linked lists and return it as a new list
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, p, np *ListNode
	p1, p2 := l1, l2
	for p1 != nil && p2 != nil {

		if p1.Val < p2.Val {
			np = p1
			p1 = p1.Next
		} else {
			np = p2
			p2 = p2.Next
		}

		if p == nil {
			head = np
			p = head
		} else {
			p.Next = np
			p = p.Next
		}
	}

	// add the rest to tail
	if p1 != nil {
		np = p1
	}

	if p2 != nil {
		np = p2
	}

	if p == nil {
		head = np
	} else {
		p.Next = np
	}

	return head
}
