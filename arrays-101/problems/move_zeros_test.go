package problems

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_moveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "example 1",
			args: args{
				nums: []int{0, 1, 0, 3, 12},
			},
			want: []int{1, 3, 12, 0, 0},
		},
		{
			name: "test 1",
			args: args{
				nums: []int{0, 0, 1, 0, 3, 0, 4},
			},
			want: []int{1, 3, 4, 0, 0, 0, 0},
		},
		{
			name: "test 2",
			args: args{
				nums: []int{1, 2, 0, 0, 0, 4},
			},
			want: []int{1, 2, 4, 0, 0, 0},
		},
	}

	for idx, f := range []func([]int){
		moveZeroes,
		moveZeroes1,
		moveZeroes2,
	} {
		t.Run(fmt.Sprintf("func#%d", idx), func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					assert := require.New(t)
					var nums = make([]int, len(tt.args.nums))
					copy(nums, tt.args.nums)
					f(nums)
					assert.Equal(tt.want, nums)
				})
			}
		})
	}
}
