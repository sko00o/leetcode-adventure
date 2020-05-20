package problems

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_removeElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name     string
		args     args
		want     int
		wantNums []int
	}{
		{
			name: "example 1",
			args: args{
				nums: []int{3, 2, 2, 3},
				val:  3,
			},
			want:     2,
			wantNums: []int{2, 2},
		},
		{
			name: "example 2",
			args: args{
				nums: []int{0, 1, 2, 2, 3, 0, 4, 2},
				val:  2,
			},
			want:     5,
			wantNums: []int{0, 1, 3, 0, 4},
		},
	}

	for idx, f := range []func([]int, int) int{
		removeElement,
		removeElement1,
	} {
		t.Run(fmt.Sprintf("func#%d", idx), func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					assert := require.New(t)
					var nums = make([]int, len(tt.args.nums))
					copy(nums, tt.args.nums)
					got := f(nums, tt.args.val)
					nums = nums[:got]
					assert.Equal(tt.want, got)
					// order can be changed.
					sort.Sort(sort.IntSlice(tt.wantNums))
					sort.Sort(sort.IntSlice(nums))
					assert.Equal(tt.wantNums, nums)
				})
			}
		})
	}
}
