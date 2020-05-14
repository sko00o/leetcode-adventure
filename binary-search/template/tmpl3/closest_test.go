package tmpl3

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findClosestElements(t *testing.T) {
	tcs := []struct {
		nums   []int
		k, x   int
		expect []int
	}{
		{[]int{1, 2, 3, 4, 5}, 4, 3, []int{1, 2, 3, 4}},
		{[]int{1, 2, 3, 4, 5}, 4, -1, []int{1, 2, 3, 4}},
		{[]int{1, 2, 3, 4, 5}, 2, 3, []int{2, 3}},
		{[]int{1, 2, 4, 5, 6}, 2, 3, []int{2, 4}},
		{[]int{1, 2, 3}, 3, 100, []int{1, 2, 3}},
		{[]int{1, 2, 3}, 2, 100, []int{2, 3}},
		{[]int{1, 2, 3}, 2, -100, []int{1, 2}},
		{[]int{1, 2, 3}, 3, -100, []int{1, 2, 3}},
		{[]int{1, 2, 3, 9}, 3, 4, []int{1, 2, 3}},
		{[]int{1, 1, 2, 3, 4, 5}, 6, 2, []int{1, 1, 2, 3, 4, 5}},
		{[]int{0, 0, 0, 1, 3, 5, 6, 7, 8, 8}, 2, 2, []int{1, 3}},
	}

	for fIdx, f := range []func([]int, int, int) []int{
		findClosestElements,
		findClosestElements1,
	} {
		for i, tc := range tcs {
			t.Run(fmt.Sprintf("task #%d, func #%d", i, fIdx), func(t *testing.T) {
				assert := require.New(t)
				got := f(tc.nums, tc.k, tc.x)
				assert.Equal(tc.expect, got)
			})
		}
	}
}
