package cp

import (
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	tasks := []struct {
		listA, listB []int
		expect       []int
	}{
		{
			listA:  []int{1, 2, 4},
			listB:  []int{1, 3, 4},
			expect: []int{1, 1, 2, 3, 4, 4},
		},
		{
			listA:  []int{2, 3, 5, 8},
			listB:  []int{1, 6, 7, 9},
			expect: []int{1, 2, 3, 5, 6, 7, 8, 9},
		},
		{
			listA:  []int{2, 6},
			listB:  []int{1, 6, 7, 9},
			expect: []int{1, 2, 6, 6, 7, 9},
		},
		{
			listA:  []int{},
			listB:  []int{},
			expect: []int{},
		},
		{
			listA:  []int{0},
			listB:  []int{},
			expect: []int{0},
		},
	}

	for _, f := range []func(l1, l2 *ListNode) *ListNode{mergeTwoLists} {
		for i, task := range tasks {
			h1 := makeLinkedList(task.listA...)
			h2 := makeLinkedList(task.listB...)
			var got []int
			for p := f(h1, h2); p != nil; p = p.Next {
				got = append(got, p.Val)
			}
			if !equal(got, task.expect) {
				t.Errorf("func %T, task #%d failed, output: %v, expect: %v", f, i, got, task.expect)
			}
		}
	}
}
