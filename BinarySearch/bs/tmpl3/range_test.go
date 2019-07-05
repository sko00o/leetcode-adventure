package tmpl3

import (
	"fmt"
	"testing"
)

func TestSearchRange(t *testing.T) {
	tcs := []struct {
		nums   []int
		target int
		expect []int
	}{
		{[]int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},
		{[]int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1}},
		{[]int{}, 6, []int{-1, -1}},
		{[]int{1}, 1, []int{0, 0}},
	}
	for i, tc := range tcs {
		t.Run(fmt.Sprintf("task #%d", i), func(t *testing.T) {
			got := searchRange(tc.nums, tc.target)
			if len(got) != 2 {
				t.Fatal("wrong output")
			}
			if got[0] != tc.expect[0] || got[1] != tc.expect[1] {
				t.Errorf("got: %v, expect: %v", got, tc.expect)
			}
		})
	}
}
