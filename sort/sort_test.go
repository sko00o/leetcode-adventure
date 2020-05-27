package sort

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSort(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{
			nums: []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3},
			want: []int{1, 3, 4, 5, 6, 6, 6, 8, 9, 14, 25, 49},
		},
		{
			nums: []int{5},
			want: []int{5},
		},
		{
			nums: []int{3, 6},
			want: []int{3, 6},
		},
		{
			nums: []int{2, 5, 4, 1, 3},
			want: []int{1, 2, 3, 4, 5},
		},
	}

	for idx, f := range []func([]int){
		bubbleSort,
		selectSort,
		selectSort1,
		insertSort,
		shellSort,
		mergeSort,
		mergeSort1,
		mergeSort2,
		quickSort,
		quickSort1,
		quickSort2,
		quickSort3,
		quickSort4,
	} {
		t.Run(fmt.Sprintf("func#%d", idx), func(t *testing.T) {
			for i, tt := range tests {
				t.Run(fmt.Sprintf("case#%d", i), func(t *testing.T) {
					assert := require.New(t)
					var nums = make([]int, len(tt.nums))
					copy(nums, tt.nums)
					f(nums)
					assert.Equal(tt.want, nums)
				})
			}
		})
	}
}
