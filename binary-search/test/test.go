package test

import (
	"fmt"
	"testing"
)

type Cases []struct {
	nums   []int
	target int
	expect int
}

var NormalCases = Cases{
	{[]int{-9, -8, -7, -5, -3, -2, -1, 0, 1, 2, 4, 5, 6, 9}, -5, 3}, // 0
	{[]int{-9, -8, -7, -5, -3, -2, -1, 0, 1, 2, 4, 5, 6, 9}, 1, 8},  // 1
	{[]int{-1, 0, 3, 5, 9, 12}, 9, 4},                               // 2
	{[]int{-1, 0, 3, 5, 9, 12}, 2, -1},                              // 3
	{[]int{5}, -5, -1},                                              // 4
	{[]int{6}, 6, 0},                                                // 5
	{[]int{}, 9, -1},                                                // 6
}

var RotateCases = Cases{
	{[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},     // 7 // right side in rotate
	{[]int{4, 5, 6, 7, 0, 1, 2}, 5, 1},     // 8
	{[]int{4, 5, 6, 7, 0, 1, 2}, 3, -1},    // 9
	{[]int{5, 7, -3, -2, -1, 3, 4}, 3, 5},  // 10 // right side in rotate
	{[]int{5, 7, -3, -2, -1, 3, 4}, 7, 1},  // 11
	{[]int{5, 7, -3, -2, -1, 3, 4}, 0, -1}, // 12
}

func CommonTest(t *testing.T, fs ...func([]int, int) int) {
	for fIdx, f := range fs {
		for i, tc := range NormalCases {
			t.Run(fmt.Sprintf("func #%d, task#%d", fIdx, i), func(t *testing.T) {
				if got := f(tc.nums, tc.target); got != tc.expect {
					t.Errorf("got: %d, want: %d", got, tc.expect)
				}
			})
		}
	}
}

func RotateTest(t *testing.T, fs ...func([]int, int) int) {
	for fIdx, f := range fs {
		for i, tc := range append(NormalCases, RotateCases...) {
			t.Run(fmt.Sprintf("func #%d, task#%d", fIdx, i), func(t *testing.T) {
				if got := f(tc.nums, tc.target); got != tc.expect {
					t.Errorf("got: %d, want: %d", got, tc.expect)
				}
			})
		}
	}
}
