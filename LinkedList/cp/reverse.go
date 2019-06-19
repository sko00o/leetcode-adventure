package cp

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

func reverseList2(h *ListNode) *ListNode {
	if h != nil {
		for p, np := h, h.Next; np != nil; np = p.Next {
			p.Next, np.Next, h = np.Next, h, np
		}
	}
	return h
}
