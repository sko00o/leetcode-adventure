package cp

// remove linked list elements
func removeElements(head *ListNode, val int) *ListNode {

	for p, pre := head, head; p != nil; p = p.Next {
		if p.Val == val {
			if p == head {
				head = p.Next
				pre = head
				continue
			} else {
				pre.Next = p.Next
				continue
			}
		}
		pre = p
	}

	return head
}
