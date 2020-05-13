package practice

import (
	"fmt"
	"testing"
)

func Test_numIslands(t *testing.T) {
	type args struct {
		grid [][]byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{
				grid: [][]byte{
					[]byte("11110"),
					[]byte("11010"),
					[]byte("11000"),
					[]byte("00000"),
				},
			},
			want: 1,
		},
		{
			name: "example 2",
			args: args{
				grid: [][]byte{
					[]byte("11000"),
					[]byte("11000"),
					[]byte("00100"),
					[]byte("00011"),
				},
			},
			want: 3,
		},
		{
			name: "test 1",
			args: args{
				grid: [][]byte{
					[]byte("001010"),
					[]byte("011111"),
					[]byte("000100"),
					[]byte("011010"),
					[]byte("001001"),
				},
			},
			want: 4,
		},
	}

	for idx, f := range []func(grid [][]byte) int{
		numIslands,
		numIslands1,
	} {
		t.Run(fmt.Sprint("func#d", idx), func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					if got := f(tt.args.grid); got != tt.want {
						t.Errorf("numIslands() = %v, want %v", got, tt.want)
					}
				})
			}
		})
	}
}
