package cp

// time O(n)
// space O(1)
func reverseList(head *ListNode) *ListNode {
	if head != nil {
		p := head
		for p.Next != nil {
			newHead := p.Next
			p.Next = newHead.Next
			newHead.Next = head
			head = newHead
		}
	}
	return head
}

func reverseList1(h *ListNode) *ListNode {
	if h != nil {
		for p, np := h, h.Next; np != nil; np = p.Next {
			p.Next, np.Next, h = np.Next, h, np
		}
	}
	return h
}

// recursive version
// time O(n)
// space O(n)
func reverseList2(h *ListNode) *ListNode {
	if h != nil && h.Next != nil {
		p := reverseList2(h.Next)
		h.Next.Next = h
		h.Next = nil
		h = p
	}
	return h
}
