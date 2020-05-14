package practice

import (
	"fmt"
	"testing"
)

func Test_numSquares(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{
				n: 12,
			},
			want: 3,
		},
		{
			name: "example 2",
			args: args{
				n: 13,
			},
			want: 2,
		},
		{
			name: "test 1",
			args: args{
				n: 3,
			},
			want: 3,
		},
		{
			name: "test 2",
			args: args{
				n: 1,
			},
			want: 1,
		},
		{
			name: "test 3",
			args: args{
				n: 0,
			},
			want: -1,
		},
		{
			name: "test 4",
			args: args{
				n: 25,
			},
			want: 1,
		},
	}

	for idx, f := range []func(int) int{
		numSquares,
	} {
		t.Run(fmt.Sprintf("func#%d", idx), func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					if got := f(tt.args.n); got != tt.want {
						t.Errorf("numSquares() = %v, want %v", got, tt.want)
					}
				})
			}
		})
	}
}
