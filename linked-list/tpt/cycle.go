package tpt

/*
Notes:
○ 判断是否有环
	§ 快慢指针，快指针一次走两步，满指针一次走一步，如果快指针遇到了慢指针，就有环。
	  快指针如果能到达链表尾部，就无环。
○ 寻找入环节点
	§ 快慢指针，快指针一次两步，幔指针一次一步。
	  链表总长为 e， 环长为 r，易证明，在相遇处的幔指针到入环节点的距离等于
	  头节点到入环节点的距离。（令从头节点到相遇处距离为 l，头节点到入环节点
	  距离为 x ，快指针走过的距离就是 2l, 快慢指针所走距离的差，就是 n 倍的环长，
	  也会是 l 。画图割补一下很好理解。）
*/

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

// time complexity: O(n)
func detectCycle1(head *ListNode) *ListNode {
	sp, fp := head, head
	for fp != nil {
		sp, fp = sp.Next, fp.Next
		if fp != nil {
			fp = fp.Next
		} else {
			return nil
		}

		if fp == sp {
			// we can prove that slow point to
			// the entry point of cycle distances
			// equal to head to entry point of cycle
			fp = head
			for fp != nil && fp != sp {
				sp, fp = sp.Next, fp.Next
			}
			return fp
		}
	}
	return nil
}
