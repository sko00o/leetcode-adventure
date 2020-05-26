package problems

import (
	"fmt"
	"testing"
)

func Test_findKthLargest(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "exmaple 1",
			args: args{
				nums: []int{3, 2, 1, 5, 6, 4},
				k:    2,
			},
			want: 5,
		},
		{
			name: "exmaple 2",
			args: args{
				nums: []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
				k:    4,
			},
			want: 4,
		},
		{
			name: "test 1",
			args: args{
				nums: []int{-1, 2, 0},
				k:    1,
			},
			want: 2,
		},
	}

	for idx, f := range []func([]int, int) int{
		findKthLargest,
		findKthLargest1,
		findKthLargest2,
		findKthLargest3,
	} {
		t.Run(fmt.Sprintf("func#%d", idx), func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					var nums = make([]int, len(tt.args.nums))
					copy(nums, tt.args.nums)
					if got := f(nums, tt.args.k); got != tt.want {
						t.Errorf("findKthLargest() = %v, want %v", got, tt.want)
					}
				})
			}
		})
	}
}
