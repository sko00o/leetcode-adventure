package cp

import (
	"testing"
)

func Test_removeElements(t *testing.T) {
	tasks := []struct {
		list   []int
		val    int
		expect []int
	}{
		{
			list:   []int{1, 2, 3, 4, 5},
			val:    3,
			expect: []int{1, 2, 4, 5},
		},
		{
			list:   []int{2, 3, 5, 6, 4, 1, 8, 9, 0, 7},
			val:    6,
			expect: []int{2, 3, 5, 4, 1, 8, 9, 0, 7},
		},
		{
			list:   []int{22, 2, 2, 22},
			val:    2,
			expect: []int{22, 22},
		},
		{
			list:   []int{33, 1, 2, 33, 3},
			val:    33,
			expect: []int{1, 2, 3},
		},
	}

	for _, f := range []func(*ListNode, int) *ListNode{removeElements} {
		for i, task := range tasks {
			h1 := makeLinkedList(task.list...)
			h2 := f(h1, task.val)
			var got []int
			for p := h2; p != nil; p = p.Next {
				got = append(got, p.Val)
			}
			if !equal(got, task.expect) {
				t.Errorf("func %T, task #%d failed, output: %v, expect: %v", f, i, got, task.expect)
			}
		}
	}
}
