package tpt

/*
	1. If there is no cycle, the fast pointer will stop at the end of the linked list.
	2. If there is a cycle, the fast pointer will eventually meet with the slow pointer.
*/

// determine cycle
// space complexity: O(1)
// time complexity: O(n)
func hasCycle(head *ListNode) bool {
	fp := head
	sp := head

	for fp != nil {
		sp = sp.Next
		fp = fp.Next
		if fp != nil {
			fp = fp.Next
		} else {
			return false
		}

		if sp == fp {
			return true
		}
	}

	return false
}

// detect cycle
// space complexity: O(1)
// time complexity: O(n^2)
func detectCycle(head *ListNode) *ListNode {
	fp := head
	sp := head

	for fp != nil {
		sp = sp.Next
		fp = fp.Next
		if fp != nil {
			fp = fp.Next
		} else {
			return nil
		}

		if sp == fp {
			// slow but it works
			for p := head; p != sp; p = p.Next {
				for tp := sp.Next; tp != sp; tp = tp.Next {
					if p == tp {
						return p
					}
				}
			}
			return sp
		}
	}

	return nil
}
