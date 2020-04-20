package classic

func copyRandomList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var nHead, np *ListNode
	for p := head; p != nil; p = p.Next {
		nNode := &ListNode{
			Val:    p.Val,
			Next:   p.Next,
			Random: p.Random,
		}
		// point the old node random to new node
		p.Random = nNode

		// recover pointer before return
		defer func(p, r *ListNode) {
			p.Random = r
		}(p, nNode.Random)

		if np == nil {
			nHead, np = nNode, nNode
		} else {
			np.Next = nNode
			np = np.Next
		}
	}

	for p := nHead; p != nil; p = p.Next {
		if p.Random != nil {
			p.Random = p.Random.Random
		}
	}

	return nHead
}
