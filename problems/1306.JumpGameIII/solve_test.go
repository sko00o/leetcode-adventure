package problems

import (
	"fmt"
	"testing"
)

func Test_canReach(t *testing.T) {
	type args struct {
		arr   []int
		start int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "example 1",
			args: args{
				arr:   []int{4, 2, 3, 0, 3, 1, 2},
				start: 5,
			},
			want: true,
		},
		{
			name: "example 2",
			args: args{
				arr:   []int{4, 2, 3, 0, 3, 1, 2},
				start: 0,
			},
			want: true,
		},
		{
			name: "example 3",
			args: args{
				arr:   []int{3, 0, 2, 1, 2},
				start: 2,
			},
			want: false,
		},
	}

	for idx, f := range []func([]int, int) bool{
		canReach,
		canReach1,
		canReach2,
	} {
		t.Run(fmt.Sprintf("func#%d", idx), func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					if got := f(tt.args.arr, tt.args.start); got != tt.want {
						t.Errorf("canReach() = %v, want %v", got, tt.want)
					}
				})
			}
		})
	}
}
