package cp

// reversed iteratively
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

// reversed recursively
// time O(n)
// space O(n)
func reverseList2(h *ListNode) *ListNode {
	l, _ := recv(h)
	return l
}

func recv(h *ListNode) (l *ListNode, r *ListNode) {
	// tail
	if h == nil || h.Next == nil {
		return h, h
	}

	l, r = recv(h.Next)
	r.Next, h.Next, r = h, r.Next, h
	return
}

func reverseList3(h *ListNode) *ListNode {
	return recv1(h)
}

func recv1(h *ListNode) *ListNode {
	if h == nil || h.Next == nil {
		return h
	}

	// new tail will be h.Next
	newHead := recv1(h.Next)
	h.Next.Next, h.Next = h, nil
	return newHead
}
