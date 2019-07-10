package problems

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_twoSum(t *testing.T) {
	tcs := []struct {
		nums   []int
		target int
		expect []int
	}{
		{
			[]int{2, 7, 11, 15},
			9,
			[]int{0, 1},
		},
		{
			[]int{-1, -2, -3, -4, -5},
			-8,
			[]int{2, 4},
		},
		{
			[]int{3, 2, 1, 4},
			6,
			[]int{1, 3},
		},
		{
			[]int{3, 3},
			6,
			[]int{0, 1},
		},
	}

	for fIdx, f := range []func([]int, int) []int{twoSum, twoSum1, twoSum2} {
		for idx, tc := range tcs {
			t.Run(fmt.Sprintf("func #%d, task #%d", fIdx, idx), func(t *testing.T) {
				got := f(tc.nums, tc.target)
				if d := cmp.Diff(got, tc.expect); d != "" {
					t.Fatal(d)
				}
			})
		}
	}
}
