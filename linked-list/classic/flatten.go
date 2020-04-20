package classic

func flatten(head *ListNode) *ListNode {
	out, _ := dfs(head, true)
	return out
}

func dfs(head *ListNode, topLevel bool) (out *ListNode, np *ListNode) {
	if head == nil {
		return
	}

	var p *ListNode
	for np, p = head, head; p != nil; np, p = p, p.Next {
		if topLevel {
			out = head
		}

		if p.Child != nil {
			next := p.Next
			cHead := p.Child
			// clear the child
			p.Child = nil
			_, cTail := dfs(cHead, false)
			p.Next = cHead
			cHead.Prev = p
			if cTail != nil {
				cTail.Next = next
				if next != nil {
					next.Prev = cTail
				}
				// jump to the tail of child list
				p = cTail
			}
		}
	}

	return
}
