package tpt

import (
	"testing"
)

func makeTwoIntersectLinkedLists(listA []int, listB []int, skipA, skipB int) (headA, headB, intersectP *ListNode) {
	var skipPA, skipPB *ListNode
	headA, skipPA = makeLinkedListWithSkip(nil, skipA, listA...)
	headB, skipPB = makeLinkedListWithSkip(skipPA, skipB, listB...)
	intersectP = skipPB
	return
}

func makeLinkedListWithSkip(sp *ListNode, skip int, list ...int) (head, skipP *ListNode) {
	var p *ListNode
	for idx, v := range list {
		n := &ListNode{
			Val: v,
		}

		if idx == skip {
			if sp != nil {
				n = sp
			}
			skipP = n
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

func Test_getIntersectionNode(t *testing.T) {
	tasks := []struct {
		listA, listB []int
		skipA, skipB int
		intersectVal int
	}{
		{
			listA: []int{4, 1, 8, 4, 5}, skipA: 2,
			listB: []int{5, 0, 1, 8, 4, 5}, skipB: 3,
			intersectVal: 8,
		},
		{
			listA: []int{0, 9, 1, 2, 4}, skipA: 3,
			listB: []int{3, 2, 4}, skipB: 1,
			intersectVal: 2,
		},
		{
			listA: []int{2, 6, 4}, skipA: 3,
			listB: []int{1, 5}, skipB: 2,
			intersectVal: 0,
		},
	}

	for i, task := range tasks {
		ha, hb, itp := makeTwoIntersectLinkedLists(task.listA, task.listB, task.skipA, task.skipB)
		if gItp := getIntersectionNode(ha, hb); gItp != itp {
			t.Errorf("task #%d failed, output: %v, expect: %v", i, gItp, itp)
		} else if gItp != nil && gItp.Val != task.intersectVal {
			t.Errorf("task #%d failed, wrong value, output: %v, expect: %v", i, gItp.Val, task.intersectVal)
		}
	}

}
