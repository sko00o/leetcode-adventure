package problems

import (
	"fmt"
	"testing"
)

func Test_heightChecker(t *testing.T) {
	type args struct {
		heights []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{
				heights: []int{1, 1, 4, 2, 1, 3},
			},
			want: 3,
		},
		{
			name: "example 2",
			args: args{
				heights: []int{5, 1, 2, 3, 4},
			},
			want: 5,
		},
		{
			name: "example 3",
			args: args{
				heights: []int{1, 2, 3, 4, 5},
			},
			want: 0,
		},
	}

	for idx, f := range []func([]int) int{
		heightChecker,
		heightChecker1,
	} {
		t.Run(fmt.Sprintf("func#%d", idx), func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					if got := f(tt.args.heights); got != tt.want {
						t.Errorf("heightChecker() = %v, want %v", got, tt.want)
					}
				})
			}
		})
	}
}
