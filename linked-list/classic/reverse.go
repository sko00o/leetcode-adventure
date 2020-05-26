package classic

/*
Notes:
○ 反转单链表
	§ 一轮遍历，pn := p->next, p->next = p->next->next, pn->next = start, start = pn，
	  注意维护头节点的指针，之后返回。
	§ 递归实现，方便理解，但由于递归栈，空间复杂度 O(1)，见 reverseList3 。
*/

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
