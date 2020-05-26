package tpt

/*
Notes:
○ 摘除倒数第 n 个节点
	§ 使用双指针，快指针先走 n 步，然后快慢指针一起推进，直到快指针到达链表尾部。
	  此时满指针指向倒数第 n 个节点的前一个节点，直接摘除即可。
	  只需要一次遍历，时间复杂度 O(n)，空间 O(1) 。
*/

// remove nth node from end of list
// space complexity: O(1)
// time complexity: O(n)
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if n < 0 || head == nil {
		return head
	}
	fp, sp := head, head

	// move n step forward
	for i := 0; i < n; i++ {
		if fp == nil {
			return head
		}
		fp = fp.Next
	}

	// keep moving, slow pointer
	// will at Nth from end of list
	pre := sp
	for fp != nil {
		fp = fp.Next
		pre = sp
		sp = sp.Next
	}

	// delete sp
	if sp == head {
		head = sp.Next
	} else {
		pre.Next = sp.Next
	}

	return head
}
