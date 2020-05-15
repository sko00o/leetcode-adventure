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
		{
			name: "test 2",
			args: args{
				grid: [][]byte{
					[]byte("110000"),
					[]byte("011110"),
					[]byte("010010"),
					[]byte("010010"),
					[]byte("011110"),
					[]byte("000000"),
				},
			},
			want: 1,
		},
		{
			name: "empty",
			want: 0,
		},
	}

	for idx, f := range []func(grid [][]byte) int{
		numIslands,
		numIslands1,
		numIslands2,
		numIslands3,
	} {
		t.Run(fmt.Sprint("func#d", idx), func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					if got := f(copyGrid(tt.args.grid)); got != tt.want {
						t.Errorf("numIslands() = %v, want %v", got, tt.want)
					}
				})
			}
		})
	}
}

func copyGrid(grid [][]byte) (out [][]byte) {
	out = make([][]byte, len(grid))
	for i := range grid {
		out[i] = make([]byte, len(grid[i]))
		for j := range grid[i] {
			out[i][j] = grid[i][j]
		}
	}
	return
}
