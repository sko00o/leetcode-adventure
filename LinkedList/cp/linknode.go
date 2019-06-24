package cp

// definition for linked list node
type ListNode struct {
	// singly-linked list
	Val  int
	Next *ListNode

	// extra for doubly-linked list
	Prev *ListNode

	// extra for multilevel doubly-linked list
	Child *ListNode
}

func makeLinkedList(list ...int) (head *ListNode) {
	var p *ListNode
	for _, v := range list {
		n := &ListNode{
			Val: v,
		}

		if p != nil {
			p.Next = n
			p = p.Next
		} else {
			// first node
			p = n
			head = n
		}
	}
	return
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
