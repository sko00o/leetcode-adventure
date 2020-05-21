package problems

import (
	"fmt"
	"testing"
)

func Test_validMountainArray(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "example 1",
			args: args{
				A: []int{2, 1},
			},
			want: false,
		},
		{
			name: "example 2",
			args: args{
				A: []int{3, 5, 5},
			},
			want: false,
		},
		{
			name: "example 3",
			args: args{
				A: []int{0, 3, 2, 1},
			},
			want: true,
		},
		{
			name: "test 1",
			args: args{
				A: []int{3, 2, 1},
			},
			want: false,
		},
		{
			name: "test 2",
			args: args{
				A: []int{1, 2, 3},
			},
			want: false,
		},
	}

	for idx, f := range []func([]int) bool{
		validMountainArray,
		validMountainArray1,
	} {
		t.Run(fmt.Sprintf("func#%d", idx), func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					if got := f(tt.args.A); got != tt.want {
						t.Errorf("validMountainArray() = %v, want %v", got, tt.want)
					}
				})
			}
		})
	}
}
