package problems

import (
	"testing"
)

func TestReverseBetween(t *testing.T) {
	tasks := []struct {
		list   []int
		m, n   int
		expect []int
	}{
		{
			list:   []int{1, 2, 3, 4, 5},
			m:      2,
			n:      4,
			expect: []int{1, 4, 3, 2, 5},
		},
		{
			list:   []int{2, 3, 5, 6, 4, 1, 8, 9, 0, 7},
			m:      2,
			n:      6,
			expect: []int{2, 1, 4, 6, 5, 3, 8, 9, 0, 7},
		},
		{
			list:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			m:      1,
			n:      6,
			expect: []int{6, 5, 4, 3, 2, 1, 7, 8, 9},
		},
		{
			list:   []int{1, 2, 3},
			m:      1,
			n:      1,
			expect: []int{1, 2, 3},
		},
		{
			list:   []int{1, 2, 3},
			m:      0,
			n:      1,
			expect: []int{1, 2, 3},
		},
		{
			list:   []int{1, 2, 3},
			m:      2,
			n:      1,
			expect: []int{1, 2, 3},
		},
		{
			list:   []int{},
			m:      1,
			n:      2,
			expect: []int{},
		},
	}

	for fIdx, f := range []func(*ListNode, int, int) *ListNode{reverseBetween} {
		for i, task := range tasks {
			h1 := makeLinkedList(task.list...)
			h2 := f(h1, task.m, task.n)
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

func makeLinkedList(list ...int) (head *ListNode) {
	var p *ListNode
	for _, v := range list {
		n := &ListNode{
			Val: v,
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

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
