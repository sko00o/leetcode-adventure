package problems

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_merge(t *testing.T) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "example 1",
			args: args{
				nums1: []int{1, 2, 3, 0, 0, 0},
				m:     3,
				nums2: []int{2, 5, 6},
				n:     3,
			},
			want: []int{1, 2, 2, 3, 5, 6},
		},
		{
			name: "test 1",
			args: args{
				nums1: []int{0},
				m:     0,
				nums2: []int{2},
				n:     1,
			},
			want: []int{2},
		},
	}

	for idx, f := range []func([]int, int, []int, int){
		merge,
		merge1,
	} {
		t.Run(fmt.Sprintf("func#%d", idx), func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					assert := require.New(t)
					var nums = make([]int, len(tt.args.nums1))
					copy(nums, tt.args.nums1)
					f(nums, tt.args.m, tt.args.nums2, tt.args.n)
					assert.Equal(nums, tt.want)
				})
			}
		})
	}
}
