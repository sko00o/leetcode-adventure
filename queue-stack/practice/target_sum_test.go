package practice

import (
	"fmt"
	"testing"
)

func Test_findTargetSumWays(t *testing.T) {
	type args struct {
		nums []int
		S    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{
				nums: []int{1, 1, 1, 1, 1},
				S:    3,
			},
			want: 5,
		},
		{
			name: "test 1",
			args: args{
				nums: []int{0, 0, 0, 0, 0, 0, 0, 0, 1},
				S:    1,
			},
			want: 256,
		},
		{
			name: "test 2",
			args: args{
				nums: []int{},
				S:    0,
			},
			want: 1,
		},
	}

	for idx, f := range []func([]int, int) int{
		findTargetSumWays,
		findTargetSumWays1,
		findTargetSumWays2,
		findTargetSumWays31,
		findTargetSumWays32,
		findTargetSumWays33,
		findTargetSumWays34,
	} {
		t.Run(fmt.Sprintf("func #%d", idx), func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					if got := f(tt.args.nums, tt.args.S); got != tt.want {
						t.Errorf("findTargetSumWays() = %v, want %v", got, tt.want)
					}
				})
			}
		})
	}
}
