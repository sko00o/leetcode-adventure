package classic

import (
	"testing"
)

func Test_rotateRight(t *testing.T) {
	tasks := []struct {
		list   []int
		val    int
		expect []int
	}{
		{
			list:   []int{1, 2, 3, 4, 5},
			val:    2,
			expect: []int{4, 5, 1, 2, 3},
		},
		{
			list:   []int{0, 1, 2},
			val:    4,
			expect: []int{2, 0, 1},
		},
		{
			list:   []int{},
			val:    0,
			expect: []int{},
		},
	}

	for fIdx, f := range []func(*ListNode, int) *ListNode{rotateRight} {
		for i, task := range tasks {
			h1 := makeLinkedList(task.list...)
			h2 := f(h1, task.val)
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
