package bs

import (
	"testing"
)

func TestSearch(t *testing.T) {
	tasks := []struct {
		nums   []int
		target int
		expect int
	}{
		{[]int{-1, 0, 3, 5, 9, 12}, 9, 4},
		{[]int{-1, 0, 3, 5, 9, 12}, 2, -1},
		{[]int{5}, -5, -1},
		{[]int{-9,-8,-7,-5,-3,-2,-1,0,1,2,4,5,6,9}, -5, 3},
		{[]int{-9,-8,-7,-5,-3,-2,-1,0,1,2,4,5,6,9}, 1, 8},
		{[]int{6}, 6, 0},
	}

	for i, tc := range tasks {
		if got := search(tc.nums, tc.target); got != tc.expect {
			t.Errorf("task #%d, got: %d, want: %d", i, got, tc.expect)
		}
	}
}
