package tpt

// detect intersection of two linked lists
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	// go to tail
	pa, pb := headA, headB
	for pa.Next != nil && pb.Next != nil {
		pa = pa.Next
		pb = pb.Next
	}

	// make cycle
	if pa.Next == nil {
		pa.Next = headB
		defer func() { pa.Next = nil }()
		return detectCycle(headA)
	}

	pb.Next = headA
	defer func() { pb.Next = nil }()
	return detectCycle(headB)
}
