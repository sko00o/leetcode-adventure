package problems

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
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
				nums: []int{1, 1, 2},
			},
			want:     2,
			wantNums: []int{1, 2},
		},
		{
			name: "example 2",
			args: args{
				nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			},
			want:     5,
			wantNums: []int{0, 1, 2, 3, 4},
		},
		{
			name: "test 1",
			args: args{
				nums: []int{},
			},
			want:     0,
			wantNums: []int{},
		},
	}

	for idx, f := range []func([]int) int{
		removeDuplicates,
		removeDuplicates1,
	} {
		t.Run(fmt.Sprintf("func#%d", idx), func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					assert := require.New(t)
					var nums = make([]int, len(tt.args.nums))
					copy(nums, tt.args.nums)
					got := f(nums)
					assert.Equal(tt.want, got)
					nums = nums[:got]
					assert.Equal(tt.wantNums, nums)
				})
			}
		})
	}
}
