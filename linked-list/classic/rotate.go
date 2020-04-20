package classic

// Given a linked list, rotate the list to the right by k places, where k is non-negative.
func rotateRight(head *ListNode, k int) *ListNode {
	var len int
	var p, pre *ListNode
	for p, pre = head, head; p != nil; pre, p, len = p, p.Next, len+1 {
	}

	if len == 0 {
		return head
	}
	k = len - (k % len)
	for pre.Next, p = head, head; k != 0; {
		pre, p, k = p, p.Next, k-1
	}

	head = p
	pre.Next = nil
	return head
}
