package cp

import (
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	tasks := []struct {
		listA, listB []int
		expect       []int
	}{
		{
			// 342 + 465 = 807
			listA:  []int{2, 4, 3},
			listB:  []int{5, 6, 4},
			expect: []int{7, 0, 8},
		},
		{
			// 99 + 99 =  198
			listA:  []int{9, 9},
			listB:  []int{9, 9},
			expect: []int{8, 9, 1},
		},
		{
			// 1 + 123 = 124
			listA:  []int{1},
			listB:  []int{3, 2, 1},
			expect: []int{4, 2, 1},
		},
		{
			// 123 + 9 = 123
			listA:  []int{3, 2, 1},
			listB:  []int{9},
			expect: []int{2, 3, 1},
		},
		{
			// 1 + 99 = 100
			listA:  []int{1},
			listB:  []int{9, 9},
			expect: []int{0, 0, 1},
		},
	}

	for fIdx, f := range []func(l1, l2 *ListNode) *ListNode{addTwoNumbers, addTwoNumbers1} {
		for i, task := range tasks {
			h1 := makeLinkedList(task.listA...)
			h2 := makeLinkedList(task.listB...)
			var got []int
			for p := f(h1, h2); p != nil; p = p.Next {
				got = append(got, p.Val)
			}
			if !equal(got, task.expect) {
				t.Errorf("func #%d, task #%d failed, output: %v, expect: %v", fIdx, i, got, task.expect)
			}
		}
	}
}

func Test_addWithBase(t *testing.T) {
	tasks := []struct {
		listA, listB []int
		base         int
		expect       []int
	}{
		{
			// 1 + 99 = 100
			listA:  []int{1},
			listB:  []int{9, 9},
			base:   10,
			expect: []int{0, 0, 1},
		},
		{
			// 1 + 99 = 9A
			listA:  []int{1},
			listB:  []int{9, 9},
			base:   16,
			expect: []int{10, 9},
		},
		{
			// 4EF + 2 = 4F1
			listA:  []int{15, 14, 4},
			listB:  []int{2},
			base:   16,
			expect: []int{1, 15, 4},
		},
	}

	for fIdx, f := range []func(l1, l2 *ListNode, base int) *ListNode{addWithBase} {
		for i, task := range tasks {
			h1 := makeLinkedList(task.listA...)
			h2 := makeLinkedList(task.listB...)
			var got []int
			for p := f(h1, h2, task.base); p != nil; p = p.Next {
				got = append(got, p.Val)
			}
			if !equal(got, task.expect) {
				t.Errorf("func #%d, task #%d failed, output: %v, expect: %v", fIdx, i, got, task.expect)
			}
		}
	}
}
