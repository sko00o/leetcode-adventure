package cp

import (
	"testing"
)

func Test_oddEvenList(t *testing.T) {
	tasks := []struct {
		list   []int
		expect []int
	}{
		{
			list:   []int{1, 2, 3, 4, 5},
			expect: []int{1, 3, 5, 2, 4},
		},
		{
			list:   []int{2, 1, 3, 5, 6, 4, 7},
			expect: []int{2, 3, 6, 7, 1, 5, 4},
		},
		{
			list:   []int{},
			expect: []int{},
		},
	}

	for fIdx, f := range []func(*ListNode) *ListNode{oddEvenList} {
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
