package classic

import (
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	tasks := []struct {
		list   []int
		expect bool
	}{
		{
			list:   []int{1, 2, 3, 4, 5},
			expect: false,
		},
		{
			list:   []int{2, 2, 3, 2, 5},
			expect: false,
		},
		{
			list:   []int{2, 1, 3, 3, 1, 2},
			expect: true,
		},
		{
			list:   []int{6, 5, 5, 6},
			expect: true,
		},
		{
			list:   []int{7},
			expect: true,
		},
		{
			list:   []int{},
			expect: true,
		},
	}

	for fIdx, f := range []func(*ListNode) bool{isPalindrome, isPalindrome1, isPalindrome2} {
		for i, task := range tasks {
			h1 := NewLinkedList(task.list...)
			if got := f(h1); got != task.expect {
				t.Errorf("func #%d, task #%d failed, output: %v, expect: %v", fIdx, i, got, task.expect)
			}
		}
	}
}
