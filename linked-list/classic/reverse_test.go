package cp

import (
	"testing"
)

func Test_reverseList(t *testing.T) {
	tasks := []struct {
		list   []int
		expect []int
	}{
		{
			list:   []int{1, 2, 3, 4, 5},
			expect: []int{5, 4, 3, 2, 1},
		},
		{
			list:   []int{2, 3, 5, 6, 4, 1, 8, 9, 0, 7},
			expect: []int{7, 0, 9, 8, 1, 4, 6, 5, 3, 2},
		},
		{
			list:   []int{2333, 233, 23, 2},
			expect: []int{2, 23, 233, 2333},
		},
		{
			list:   []int{},
			expect: []int{},
		},
	}

	for fIdx, f := range []func(*ListNode) *ListNode{
		reverseList,
		reverseList1,
		reverseList2,
		reverseList3,
	} {
		for i, task := range tasks {
			h1 := makeLinkedList(task.list...)
			h2 := f(h1)
			var got []int
			for p := h2; p != nil; p = p.Next {
				got = append(got, p.Val)
			}
			if !equal(got, task.expect) {
				t.Errorf("func #%d, task #%d failed, output: %v, expect: %v", fIdx, i, got, task.expect)
			}
		}
	}
}
