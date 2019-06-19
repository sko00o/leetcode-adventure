package tpt

import (
	"testing"
)

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

func Test_removeNthFromEnd(t *testing.T) {
	tasks := []struct {
		list   []int
		n      int
		expect []int
	}{
		{list: []int{1, 2, 3, 4, 5}, n: 2, expect: []int{1, 2, 3, 5}},
		{list: []int{1}, n: 1, expect: []int{}},
		{list: []int{1, 2}, n: 2, expect: []int{2}},
	}

	for i, task := range tasks {
		h1 := makeLinkedList(task.list...)
		h2 := removeNthFromEnd(h1, task.n)
		var got []int
		for p := h2; p != nil; p = p.Next {
			got = append(got, p.Val)
		}
		if !equal(got, task.expect) {
			t.Errorf("task #%d failed, output: %v, expect: %v", i, got, task.expect)
		}
	}
}
