package problems

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m >= n || m < 1 || n < 1 {
		return head
	}
	// fix to index
	m--
	n--

	for idx, pre, p := 0, head, head; p != nil; idx, pre, p = idx+1, p, p.Next {
		if idx == m {
			nHead := reverseRange(p, n-m)
			if m == 0 {
				head = nHead
			} else {
				pre.Next = nHead
			}
			break
		}
	}

	return head
}

func reverseRange(head *ListNode, l int) *ListNode {
	if head != nil {
		nTail := head
		for nTail.Next != nil && l != 0 {
			tmp := nTail.Next
			nTail.Next = tmp.Next
			tmp.Next = head
			head = tmp
			if l > 0 {
				l--
			}
		}
	}
	return head
}
