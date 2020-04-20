package classic

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}

	val, p := f(head, head)
	return val == p.Val && p.Next == nil
}

// recursive stack
func f(a, head *ListNode) (int, *ListNode) {
	if a.Next == nil {
		return a.Val, head
	}

	val, p := f(a.Next, head)
	if p != nil && p.Val == val {
		p = p.Next
	}
	return a.Val, p
}

func isPalindrome1(head *ListNode) bool {
	return f1(head, &head)
}

// another interesting submission (https://leetcode.com/submissions/detail/237334605/)
func f1(h1 *ListNode, h2 **ListNode) bool {
	if h1 == nil {
		return true
	}

	// also make recursive stack here
	if !f1(h1.Next, h2) {
		return false
	}

	if h1.Val != (*h2).Val {
		return false
	}

	*h2 = (*h2).Next
	return true
}

func isPalindrome2(head *ListNode) bool {
	if head == nil {
		return true
	}

	// get length
	length := 0
	for p := head; p != nil; length++ {
		p = p.Next
	}

	// go to middle position
	fp, sp := head, head
	for i := 0; i < length/2; i++ {
		fp = fp.Next
	}

	// reverse middle to end
	for p, np := fp, fp.Next; np != nil; np = p.Next {
		p.Next, np.Next, fp = np.Next, fp, np
	}

	for fp != nil && sp != nil {
		if fp.Val != sp.Val {
			return false
		}
		fp = fp.Next
		sp = sp.Next
	}

	return true
}
