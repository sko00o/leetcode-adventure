package prac

import (
	"testing"
)

func TestFindMin(t *testing.T) {
	tasks := []struct {
		list   []int
		expect int
	}{
		{
			list:   []int{1, 2, 3, 4, 5},
			expect: 1,
		},
		{
			list:   []int{2, 3, 4, 5, 6, 7, 8, 9, 0, 1},
			expect: 0,
		},
		{
			list:   []int{1, 3, 5},
			expect: 1,
		},

		// test for
		// `Find Minimum in Rotated Sorted Array II`
		{
			list:   []int{2, 2, 2, 0, 1},
			expect: 0,
		},
		{
			list:   []int{3, 3, 3, 1, 3},
			expect: 1,
		},
		{
			list:   []int{3, 1, 3, 3, 3},
			expect: 1,
		},
	}

	for fIdx, f := range []func([]int) int{findMin, findMin1} {
		for i, task := range tasks {
			got := f(task.list)
			if got != task.expect {
				t.Errorf("func #%d, task #%d failed, output: %v, expect: %v", fIdx, i, got, task.expect)
			}
		}
	}
}
