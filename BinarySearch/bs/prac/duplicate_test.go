package prac

import (
	"fmt"
	"testing"
)

func Test_findDuplicate(t *testing.T) {
	tcs := []struct {
		nums []int
		want int
	}{
		{
			[]int{1, 3, 4, 2, 2},
			2,
		},
		{
			[]int{3, 1, 3, 4, 2},
			3,
		},
		{
			[]int{3, 3, 3, 4, 2, 1},
			3,
		},
	}

	for fIdx, f := range []func([]int) int{findDuplicate, findDuplicate1} {
		for idx, tc := range tcs {
			t.Run(fmt.Sprintf("func #%d, task #%d", fIdx, idx), func(t *testing.T) {
				if got := f(tc.nums); got != tc.want {
					t.Fatalf("got: %d, want: %d", got, tc.want)
				}
			})
		}
	}
}
