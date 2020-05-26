package classic

/*
Notes:
○ 将链表中不同索引的节点分为两组，偶数号为一组，奇数号为一组。偶数组接在奇数组后面。
	§ 双指针法
*/

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	odd, even := head, head.Next

	for odd != nil && even != nil && even.Next != nil {
		evenH := odd.Next
		odd.Next = even.Next
		odd = odd.Next // odd move forward

		even.Next = odd.Next
		even = even.Next // even move forward
		odd.Next = evenH
	}

	return head
}
