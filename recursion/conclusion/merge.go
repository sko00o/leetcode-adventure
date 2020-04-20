package conclusion

import (
	cp "github.com/sko00o/leetcode-adventure/linked-list/classic"
)

func mergeTwoLists(l1 *cp.ListNode, l2 *cp.ListNode) *cp.ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}

	l2.Next = mergeTwoLists(l2.Next, l1)
	return l2
}
