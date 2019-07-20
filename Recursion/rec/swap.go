package rec

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head != nil {
		second := head.Next
		if second != nil {
			second.Next = swapPairs(second.Next)
			head.Next = second.Next
			second.Next = head
			head = second
		}
	}
	return head
}
