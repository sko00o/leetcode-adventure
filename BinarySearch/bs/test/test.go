package test

import (
	"fmt"
	"testing"
)

var Cases = []struct {
	nums   []int
	target int
	expect int
}{
	{[]int{-1, 0, 3, 5, 9, 12}, 9, 4},
	{[]int{-1, 0, 3, 5, 9, 12}, 2, -1},
	{[]int{5}, -5, -1},
	{[]int{-9, -8, -7, -5, -3, -2, -1, 0, 1, 2, 4, 5, 6, 9}, -5, 3},
	{[]int{-9, -8, -7, -5, -3, -2, -1, 0, 1, 2, 4, 5, 6, 9}, 1, 8},
	{[]int{6}, 6, 0},
	{[]int{}, 9, -1},
}

func CommonTest(t *testing.T, f ...func([]int, int) int) {
	for fIdx, f := range f {
		for i, tc := range Cases {
			t.Run(fmt.Sprintf("func #%d, task#%d", fIdx, i), func(t *testing.T) {
				if got := f(tc.nums, tc.target); got != tc.expect {
					t.Errorf("got: %d, want: %d", got, tc.expect)
				}
			})
		}
	}
}
