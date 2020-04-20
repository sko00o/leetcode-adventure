package rec

import (
	"testing"
)

func Test_swapPairs(t *testing.T) {
	tasks := []struct {
		list   []int
		expect []int
	}{
		{
			list:   []int{1, 2, 3, 4},
			expect: []int{2, 1, 4, 3},
		},
		{
			list:   []int{1, 2, 3},
			expect: []int{2, 1, 3},
		},
		{
			list:   []int{1},
			expect: []int{1},
		},
		{
			list:   []int{},
			expect: []int{},
		},
	}

	for fIdx, f := range []func(*ListNode) *ListNode{swapPairs} {
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
