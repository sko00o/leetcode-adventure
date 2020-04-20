package classic

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addWithBase(l1, l2, 10)
}

// will not modify origin linked list, and more flexible
func addWithBase(p1 *ListNode, p2 *ListNode, base int) *ListNode {
	var head, p *ListNode
	var flow int

	for ; p1 != nil && p2 != nil; p1, p2 = p1.Next, p2.Next {
		p, flow = add(&head, p, p1.Val+p2.Val+flow, base)
	}

	var rest *ListNode
	if p1 != nil {
		rest = p1
	} else if p2 != nil {
		rest = p2
	}
	for ; rest != nil; rest = rest.Next {
		p, flow = add(&head, p, rest.Val+flow, base)
	}

	for flow != 0 {
		p, flow = add(&head, p, flow, base)
	}

	return head
}

func add(head **ListNode, p *ListNode, val int, base int) (*ListNode, int) {
	np := &ListNode{Val: val}
	// head init
	if p == nil {
		*head = np
		p = *head
	} else {
		p.Next = np
		p = p.Next
	}

	nVal := p.Val
	p.Val = nVal % base
	return p, nVal / base
}

// low memory cost, but change origin linked list
func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, p *ListNode
	var flow int
	base := 10
	for l1 != nil && l2 != nil {
		if p == nil {
			head = l1
			p = head
		} else {
			p.Next = l1
			p = p.Next
		}

		p.Val = p.Val + l2.Val + flow
		flow = p.Val / base
		p.Val %= base

		l1 = l1.Next
		l2 = l2.Next
	}

	if l2 != nil {
		l1 = l2
	}
	for ; l1 != nil; l1 = l1.Next {
		p.Next = l1
		p = p.Next

		p.Val = p.Val + flow
		flow = p.Val / base
		p.Val %= base
		if flow == 0 {
			break
		}
	}

	for flow != 0 {
		p.Next = &ListNode{Val: flow}
		p = p.Next
		flow = p.Val / base
		p.Val %= base
	}

	return head
}
